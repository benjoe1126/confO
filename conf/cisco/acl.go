package cisco

import (
	"confdecl/network"
)

const (
	SIMPLE   = 0
	EXTENDED = 1
)

type ACL interface {
	Configure() (string, error)
	Allow(network.IPv4Addr)
	Deny()
}
