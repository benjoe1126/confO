package network

type InterfaceNetwork interface {
	AddIpv4(addr IPv4Addr)
	AddIpv6(addr IPv6Addr)
	SetState() []string
	Configure() (string, error)
}
