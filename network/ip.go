package network

import (
	"confdecl/utils"
	"fmt"
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
	_addrNumeric []uint64
	prefix       int8 `yaml:"prefix"`
}

func (ip IPv6Addr) Print() string {
	return fmt.Sprintf("%s/%d", ip.Addr, ip.prefix)
}

func NewIPv6Addr(addr string, prefixlen int8) (IPv6Addr, error) {
	address, _ := utils.CalcIpv6ToNumeric(addr)
	return IPv6Addr{
		Addr:         addr,
		_addrNumeric: address,
		prefix:       prefixlen,
	}, nil
}
