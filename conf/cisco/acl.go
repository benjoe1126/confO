package cisco

import (
	"confdecl/network"
	"confdecl/utils"
	"strconv"
	"strings"
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

func strToIpv4(str string) network.IPv4Addr {
	ssplit := strings.Split(str, "/")
	prefNum, _ := strconv.ParseInt(ssplit[1], 10, 8)
	wildcardMask := utils.ChangeNetmaskToWildcard(utils.PrefixToDottedDecimal(int(prefNum)))
	ret, _ := network.NewIpv4(ssplit[0], wildcardMask)
	return ret
}

func newRule(src string, dst string, protocol string, port int, policy string) rule {

	source := strToIpv4(src)
	if dst == "" || protocol == "" || port == -1 {
		return rule{Source: source, Policy: policy}
	}
	dest := strToIpv4(dst)
	return rule{
		Source:      source,
		Destination: dest,
		Protocol:    protocol,
		Port:        port,
		Policy:      policy,
	}

}

type ACL struct {
	Number int    `yaml:"number,omitempty"`
	Name   string `yaml:"name,omitempty"`
	Rules  []rule `yaml:"rules"`
}

// TODO write function, it should use NewRule, but it's not visible from outside BIG BRAIN
func NewAcl(num int, name string) ACL {
	return ACL{
		Number: num,
		Name:   name,
		Rules:  make([]rule, 0),
	}
}
func (acl *ACL) AddRule(src string, dst string, protocol string, port int, policy string) {
	acl.Rules = append(acl.Rules, newRule(src, dst, protocol, port, policy))
}
func (acl *ACL) Configure() (string, error) {
	return "", nil
}
