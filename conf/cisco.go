package conf

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

const (
	DEFAULT    = 0
	PRIVILEGED = 1
	CONF_T     = 2
	CONF_INT   = 3
	CONF_OSPF  = 4
	CONF_EIGRP = 5
	CONF_LINE  = 6
	CONF_BGP   = 7
)

var (
	state            = DEFAULT
	currentInterface = ""
)

func enable() (string, bool) {
	if state == DEFAULT {
		state = PRIVILEGED
		return "enable", true
	}
	return "", false
}

func confT() (string, bool) {
	if state == PRIVILEGED {
		state = CONF_T
		return "configure terminal", true
	}
	return "", false
}

func inter(str string) (string, bool) {
	if state != PRIVILEGED && state != DEFAULT {
		state = CONF_INT
		currentInterface = str
		return fmt.Sprintf("interface %s", str), true
	}
	return "", false
}

type IPv4Addr struct {
	addr         string
	_addrNumeric uint32
	netmask      string
	prefix       int
}

func (ip IPv4Addr) Print() {
	fmt.Printf("%s/%d\n", ip.addr, ip.prefix)
}
func (ip IPv4Addr) assign() (string, bool) {
	if state == CONF_INT {
		return fmt.Sprintf("ip address %s %s", ip.addr, ip.netmask), true
	}
	return "", false
}
func NewIpv4(addr string, netmask string) (IPv4Addr, error) {
	prefArr := strings.Split(netmask, ".")
	prefix := 0
	var numeric uint32 = 0
	for i, val := range prefArr {
		temp, err := strconv.ParseUint(val, 10, 8)
		if err != nil {
			return IPv4Addr{}, err
		}
		numeric += uint32(temp) << uint32(24-i*8)
		prefix += bits.OnesCount(uint(temp))
		if temp != 255 {
			break
		}
	}
	return IPv4Addr{
		addr:         addr,
		netmask:      netmask,
		prefix:       prefix,
		_addrNumeric: numeric,
	}, nil
}

type IPv6Addr struct {
	addr         string
	_addrNumeric [2]int64
	prefix       int8
}

func (i *IPv6Addr) assign() (string, bool) {
	if state == CONF_INT {
		return fmt.Sprintf("ipv6 address %s/%d", i.addr, i.prefix), true
	}
	return "", false
}

type Interface struct {
	name          string
	ipv4Addresses []IPv4Addr
	ipv6Addresses []IPv6Addr
}

func (i *Interface) AddIpv4(addr IPv4Addr) {
	i.ipv4Addresses = append(i.ipv4Addresses, addr)
}
func (i *Interface) AddIpv6(addr IPv6Addr) {
	i.ipv6Addresses = append(i.ipv6Addresses, addr)
}

func (i *Interface) setState() []string {
	ret := make([]string, 0)
	for currentInterface != i.name || state != CONF_INT {
		switch state {
		case DEFAULT:
			str, _ := enable()
			ret = append(ret, str)
			break
		case PRIVILEGED:
			str, _ := confT()
			ret = append(ret, str)
			break
		case CONF_T:
			str, _ := inter(i.name)
			ret = append(ret, str)
			break
		case CONF_INT:
			if currentInterface != i.name {
				str, _ := inter(i.name)
				ret = append(ret, str)
			}
			break
		default:
			str, _ := inter(i.name)
			ret = append(ret, str)
			break
		}
	}
	return ret
}

func (i *Interface) Configure() (string, error) {
	ret := i.setState()
	for _, addr := range i.ipv4Addresses {
		str, _ := addr.assign()
		ret = append(ret, str)
	}
	for _, addr := range i.ipv6Addresses {
		str, _ := addr.assign()
		ret = append(ret, str)
	}
	return strings.Join(ret, "\n"), nil
}

func NewInterface(name string) Interface {
	return Interface{
		name:          name,
		ipv4Addresses: make([]IPv4Addr, 0),
		ipv6Addresses: make([]IPv6Addr, 0),
	}
}

type Interfaces struct {
	interfaces []Interface
}
