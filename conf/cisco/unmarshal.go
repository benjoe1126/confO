package cisco

import (
	"confdecl/network"
	"confdecl/utils"
	"fmt"
	"strconv"
	"strings"
)

// Rule unmarshal
func (r *Rule) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type defaultConf struct {
		Source      string `yaml:"source"`
		Destination string `yaml:"destination,omitempty"`
		Protocol    string `yaml:"protocol,omitempty"`
		Port        int    `yaml:"port"`
		SeqNum      int    `yaml:"seqNum,omitempty"`
	}
	var o defaultConf
	err := unmarshal(&o)
	if err != nil {
		return fmt.Errorf("error parsing Rule %w", err)
	}
	ipSrc, err := searchKeyword(o.Source)
	if err != nil {
		return err
	}
	if ipSrc.Prefix == -1 {
		ipSrc = strToIpv4(o.Source)
	}
	r.Source = ipSrc
	r.Protocol = o.Protocol
	r.Port = o.Port
	r.SeqNum = utils.Max(o.SeqNum, 5)
	if o.Destination == "" {
		return nil
	}
	ipDst, err := searchKeyword(o.Destination)
	if err != nil {
		return err
	}
	if ipDst.Prefix == -1 {
		ipDst = strToIpv4(o.Destination)
	}
	r.Destination = ipDst
	return nil
}

// TODO fix problem with sequence numbers
// Unmarshal ACL-s
func (acl *ACL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type tempAcl struct {
		Number int    `yaml:"number,omitempty"`
		Name   string `yaml:"name,omitempty"`
		Permit []Rule `yaml:"permit,omitempty"`
		Deny   []Rule `yaml:"deny,omitempty"`
	}
	var tmp tempAcl
	err := unmarshal(&tmp)
	if err != nil {
		return fmt.Errorf("cant unmarshal ACL %w", err)
	}
	acl.Number = tmp.Number
	acl.Name = tmp.Name
	if len(tmp.Permit) == 0 && len(tmp.Permit) == 0 {
		return fmt.Errorf("acl should have at least one rules")
	}
	for _, permit := range tmp.Permit {
		acl.PermitRule(permit)
	}
	for _, deny := range tmp.Deny {
		acl.DenyRule(deny)
	}
	return nil
}

func (i *InterfaceCisco) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type tempInt struct {
		Name          string             `yaml:"id"`
		Ipv4Addresses []network.IPv4Addr `yaml:"ipv4,omitempty"`
		Ipv6Addresses []network.IPv6Addr `yaml:"ipv6,omitempty"`
		AclApplied    []string           `yaml:"acl,omitempty"`
	}
	var tmp tempInt
	err := unmarshal(&tmp)
	if err != nil {
		return fmt.Errorf("can't unmarshal cisco interface %w", err)
	}
	i.Ipv4Addresses = tmp.Ipv4Addresses
	i.Ipv6Addresses = tmp.Ipv6Addresses
	i.ACLApplied = make([]aclInt, 0)
	for _, str := range tmp.AclApplied {
		split := strings.Split(str, " ")
		if split[1] != "in" && split[1] != "out" {
			return fmt.Errorf("invalid direction: %s", split[1])
		}
		toInt, err := strconv.Atoi(split[0])
		if err != nil {
			toInt = -1
		}
		i.ACLApplied = append(i.ACLApplied, aclInt{ACLApplied: ACL{Name: str, Number: toInt}, Direction: split[1]})
	}
	return nil
}
