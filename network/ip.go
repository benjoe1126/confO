package network

import (
	"confdecl/utils"
	"fmt"
	"strconv"
	"strings"
)

type IP interface {
	assign() string
}

type IPv4Addr struct {
	Addr         string `yaml:"addr"`
	_addrNumeric uint32
	Netmask      string `yaml:"mask"`
	Prefix       int8   `yaml:"prefix"`
}

func (ip IPv4Addr) PrintWPrefix() string {
	return fmt.Sprintf("%s/%d", ip.Addr, ip.Prefix)
}
func (ip IPv4Addr) PrintWNetmask() string {
	return fmt.Sprintf("%s %s", ip.Addr, ip.Netmask)
}

func NewIpv4(addr string, netmask string) (IPv4Addr, error) {

	prefix, err := utils.CalcPrefix(netmask)
	if err != nil {
		return IPv4Addr{}, err
	}
	numeric, err := utils.CalcIpToNumeric(addr)
	if err != nil {
		return IPv4Addr{}, err
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

func (ip IPv6Addr) Print() string {
	return fmt.Sprintf("%s/%d", ip.Addr, ip.prefix)
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
		num2, err := strconv.ParseUint(addrSplt[i+4], 10, 16)
		if err != nil {
			return IPv6Addr{}, err
		}
		num1 += num
		num1 = num1 << 16
		num2 += num2
		num2 = num2 << 16

	}
	return IPv6Addr{
		Addr:         addr,
		_addrNumeric: [2]uint64{num1, num2},
		prefix:       prefixlen,
	}, nil
}
