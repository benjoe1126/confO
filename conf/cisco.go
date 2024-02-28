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
	Addr         string `yaml:"addr"`
	_addrNumeric uint32
	Netmask      string `yaml:"mask"`
	Prefix       int    `yaml:"prefix"`
}

func (ip IPv4Addr) Print() {
	fmt.Printf("%s/%d\n", ip.Addr, ip.Prefix)
}
func (ip IPv4Addr) assign() (string, bool) {
	if state == CONF_INT {
		return fmt.Sprintf("ip address %s %s", ip.Addr, ip.Netmask), true
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
		Addr:         addr,
		Netmask:      netmask,
		Prefix:       prefix,
		_addrNumeric: numeric,
	}, nil
}

type IPv6Addr struct {
	Addr         string `yaml:"ipv6addr"`
	_addrNumeric [2]uint64
	prefix       int8 `yaml:"prefix"`
}

func (i *IPv6Addr) assign() (string, bool) {
	if state == CONF_INT {
		return fmt.Sprintf("ipv6 address %s/%d", i.Addr, i.prefix), true
	}
	return "", false
}

func NewIPv6Addr(addr string, prefixlen int8) (IPv6Addr, error) {
	addrSplt := strings.Split(addr, ":")
	var num1 uint64 = 0
	var num2 uint64 = 0
	for i := 0; 4 > i; i++ {
		num, err := strconv.ParseUint(addrSplt[i], 10, 16)
		if err != nil {
			return IPv6Addr{}, err
		}
		num_2, err := strconv.ParseUint(addrSplt[i+4], 10, 16)
		if err != nil {
			return IPv6Addr{}, err
		}
		num1 += num
		num1 = num1 << 16
		num2 += num_2
		num2 = num2 << 16

	}
	return IPv6Addr{
		Addr:         addr,
		_addrNumeric: [2]uint64{num1, num2},
		prefix:       prefixlen,
	}, nil
}

type Interface struct {
	Name          string     `yaml:"id"`
	Ipv4Addresses []IPv4Addr `yaml:"ipv4,omitempty"`
	Ipv6Addresses []IPv6Addr `yaml:"ipv6,omitempty"`
}

func (i *Interface) AddIpv4(addr IPv4Addr) {
	i.Ipv4Addresses = append(i.Ipv4Addresses, addr)
}
func (i *Interface) AddIpv6(addr IPv6Addr) {
	i.Ipv6Addresses = append(i.Ipv6Addresses, addr)
}

func (i *Interface) setState() []string {
	ret := make([]string, 0)
	for currentInterface != i.Name || state != CONF_INT {
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
			str, _ := inter(i.Name)
			ret = append(ret, str)
			break
		case CONF_INT:
			if currentInterface != i.Name {
				str, _ := inter(i.Name)
				ret = append(ret, str)
			}
			break
		default:
			str, _ := inter(i.Name)
			ret = append(ret, str)
			break
		}
	}
	return ret
}

func (i *Interface) Configure() (string, error) {
	ret := i.setState()
	for _, addr := range i.Ipv4Addresses {
		str, _ := addr.assign()
		ret = append(ret, str)
	}
	for _, addr := range i.Ipv6Addresses {
		str, _ := addr.assign()
		ret = append(ret, str)
	}
	return strings.Join(ret, "\n"), nil
}

func NewInterface(name string) Interface {
	return Interface{
		Name:          name,
		Ipv4Addresses: make([]IPv4Addr, 0),
		Ipv6Addresses: make([]IPv6Addr, 0),
	}
}

type ACL struct {
}
