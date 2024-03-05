package cisco

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

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
