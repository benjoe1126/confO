package cisco

import (
	"confdecl/network"
	"fmt"
	"strings"
)

type InterfaceCisco struct {
	Name          string             `yaml:"id"`
	Ipv4Addresses []network.IPv4Addr `yaml:"ipv4,omitempty"`
	Ipv6Addresses []network.IPv6Addr `yaml:"ipv6,omitempty"`
}

func (i *InterfaceCisco) AddIpv4(addr network.IPv4Addr) {
	i.Ipv4Addresses = append(i.Ipv4Addresses, addr)
}
func (i *InterfaceCisco) AddIpv6(addr network.IPv6Addr) {
	i.Ipv6Addresses = append(i.Ipv6Addresses, addr)
}

func (i *InterfaceCisco) SetState() []string {
	ret := make([]string, 0)
	for currentInterface != i.Name || State != CONF_INT {
		switch State {
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

func (i *InterfaceCisco) Configure() (string, error) {
	ret := i.SetState()
	for _, addr := range i.Ipv4Addresses {
		ret = append(ret, fmt.Sprintf("ip address %s", addr.PrintWNetmask()))
	}
	for _, addr := range i.Ipv6Addresses {
		ret = append(ret, fmt.Sprintf("ip address %s", addr.Print()))
	}
	return strings.Join(ret, "\n"), nil
}

func NewInterface(name string) InterfaceCisco {
	return InterfaceCisco{
		Name:          name,
		Ipv4Addresses: make([]network.IPv4Addr, 0),
		Ipv6Addresses: make([]network.IPv6Addr, 0),
	}
}

type SubInterface struct {
	Inter  InterfaceCisco `yaml:"interface"`
	Vlanid int16          `yaml:"vlanid"`
}

func (s *SubInterface) AddIpv4(addr network.IPv4Addr) {
	s.Inter.Ipv4Addresses = append(s.Inter.Ipv4Addresses, addr)
}
func (s *SubInterface) AddIpv6(addr network.IPv6Addr) {
	s.Inter.Ipv6Addresses = append(s.Inter.Ipv6Addresses, addr)
}
func (s *SubInterface) SetState() []string {
	return s.Inter.SetState()
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
		Inter: InterfaceCisco{
			Name:          name,
			Ipv4Addresses: make([]network.IPv4Addr, 0),
			Ipv6Addresses: make([]network.IPv6Addr, 0),
		},
		Vlanid: vlanid,
	}, nil
}
