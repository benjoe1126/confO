package cisco

import (
	"confdecl/network"
	"confdecl/utils"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	Source      network.IPv4Addr `yaml:"source"`
	Destination network.IPv4Addr `yaml:"destination,omitempty"`
	Protocol    string           `yaml:"protocol,omitempty"`
	Port        int              `yaml:"port"`
	SeqNum      int              `yaml:"seqNum,omitempty"`
}

func newRule(src string, dst string, protocol string, port int, seq int) Rule {

	source := strToIpv4(src)
	if dst == "" || protocol == "" || port == -1 {
		return Rule{Source: source}
	}
	dest := strToIpv4(dst)
	return Rule{
		Source:      source,
		Destination: dest,
		Protocol:    protocol,
		Port:        port,
		SeqNum:      seq,
	}

}

type ACL struct {
	Number int    `yaml:"number,omitempty"`
	Name   string `yaml:"name,omitempty"`
	Permit []Rule `yaml:"permit"`
	Deny   []Rule `yaml:"deny"`
}

// TODO write function, it should use NewRule, but it's not visible from outside BIG BRAIN
func NewAcl(num int, name string) ACL {
	return ACL{
		Number: num,
		Name:   name,
		Permit: make([]Rule, 0),
		Deny:   make([]Rule, 0),
	}
}

func foundAndMax(r []Rule, seq int) (bool, int) {
	in := false
	maxSeq := 0
	for _, rul := range r {
		fmt.Printf("rul.seq: %d, seq: %d\n", rul.SeqNum, seq)
		if rul.SeqNum == seq {
			in = true
		}
		maxSeq = utils.Max(rul.SeqNum, maxSeq)
	}
	return in, maxSeq
}

// PermitRule adds rules to permit list, if the seq number of the rule is already in use, it will append it to the list giving it the seq number of the highest existing rule + 5
func (acl *ACL) PermitRule(r Rule) {
	wasInPermit, max1 := foundAndMax(acl.Permit, r.SeqNum)
	wasInDeny, max2 := foundAndMax(acl.Deny, r.SeqNum)
	seqToUse := r.SeqNum
	if r.SeqNum == 0 {
		seqToUse = 5
	}
	if wasInPermit || wasInDeny {
		seqToUse = (utils.Max(max1, max2)) + 5
	}
	r.SeqNum = seqToUse
	acl.Permit = append(acl.Permit, r)
}
func (acl *ACL) DenyRule(r Rule) {
	wasInPermit, max1 := foundAndMax(acl.Permit, r.SeqNum)
	wasInDeny, max2 := foundAndMax(acl.Deny, r.SeqNum)
	seqToUse := r.SeqNum
	if r.SeqNum == 0 {
		seqToUse = 5
	}
	if wasInPermit || wasInDeny {
		seqToUse = (utils.Max(max1, max2)) + 5
	}
	r.SeqNum = seqToUse
	acl.Deny = append(acl.Deny, r)
}
func (acl *ACL) Configure() (string, error) {
	return "", nil
}

func (acl *ACL) Print() {
	fmt.Println(acl.Name + " Permits:")
	prettyPrint(acl.Permit)
	fmt.Println("Denies:")
	fmt.Println(prettyPrint(acl.Deny))
}

// TODO ha már tartományt adnak ne legyen host portion
func strToIpv4(str string) network.IPv4Addr {
	ssplit := strings.Split(str, "/")
	if len(ssplit) == 1 {
		ssplit = strings.Split(ssplit[0], " ")
		ret, _ := network.NewIpv4(ssplit[0], utils.ChangeNetmaskToWildcard(ssplit[1]))
		return ret
	}
	prefNum, _ := strconv.ParseInt(ssplit[1], 10, 8)
	wildcardMask := utils.ChangeNetmaskToWildcard(utils.PrefixToDottedDecimal(int(prefNum)))
	ret, _ := network.NewIpv4(ssplit[0], wildcardMask)
	return ret
}
func searchKeyword(src string) (network.IPv4Addr, error) {
	if strings.Contains(src, "any") {
		return network.FromKeywordToIP("", "any")
	} else if strings.Contains(src, "host") {
		host := strings.Split(src, " ")[1]
		return network.FromKeywordToIP(host, "host")
	}
	return network.IPv4Addr{
		Prefix: -1,
	}, nil
}
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
	return string(s)
}
