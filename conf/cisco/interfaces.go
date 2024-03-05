package cisco

import (
	"fmt"
	"strings"
)

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

type SubInterface struct {
	Inter  Interface `yaml:"interface"`
	Vlanid int16     `yaml:"vlanid"`
}

func (s *SubInterface) AddIpv4(addr IPv4Addr) {
	s.Inter.Ipv4Addresses = append(s.Inter.Ipv4Addresses, addr)
}
func (s *SubInterface) AddIpv6(addr IPv6Addr) {
	s.Inter.Ipv6Addresses = append(s.Inter.Ipv6Addresses, addr)
}
func (s *SubInterface) setState() []string {
	return s.Inter.setState()
}
func (s *SubInterface) Configure() (string, error) {
	ret := make([]string, 0)
	temp, err := s.Inter.Configure()
	if err != nil {
		return "", err
	}
	ret = append(ret, temp)
	ret = append(ret, fmt.Sprintf("encapsulation dot1q %d", s.Vlanid))
	return strings.Join(ret, "\n"), nil

}

func NewSubinterface(name string, vlanid int16) (SubInterface, error) {
	return SubInterface{
		Inter: Interface{
			Name:          name,
			Ipv4Addresses: make([]IPv4Addr, 0),
			Ipv6Addresses: make([]IPv6Addr, 0),
		},
		Vlanid: vlanid,
	}, nil
}
