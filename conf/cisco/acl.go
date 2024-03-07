package cisco

import (
	"confdecl/network"
)

const (
	deny   = "deny"
	permit = "permit"
)

type rule struct {
	Source      network.IPv4Addr `yaml:"source"`
	Destination network.IPv4Addr `yaml:"destination,omitempty"`
	Protocol    string           `yaml:"protocol,omitempty"`
	Port        int              `yaml:"port"`
	Policy      string           `yaml:"policy"`
}

func NewRule() rule {
	return rule{}
}

type ACL struct {
	Number int    `yaml:"number,omitempty"`
	Name   string `yaml:"name,omitempty"`
	Rules  []rule `yaml:"rules"`
}

// TODO write function, it should use NewRule, but it's not visible from outside BIG BRAIN
func NewAcl() ACL {
	return ACL{}
}
