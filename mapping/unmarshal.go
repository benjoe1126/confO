package mapping

import (
	"confdecl/conf/cisco"
	"fmt"
)

func (c *CiscoConf) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type tempConf struct {
		Interfaces    []cisco.InterfaceCisco `yaml:"interfaces,omitempty"`
		SubInterfaces cisco.SubInterface     `yaml:"subInterfaces,omitempty"`
		ACL           []cisco.ACL            `yaml:"acl,omitempty"`
	}
	var tmp tempConf
	err := unmarshal(&tmp)
	if err != nil {
		return fmt.Errorf("can't unmarshal to cisco conf %w", err)
	}
	for i, iface := range c.Interfaces {
		for j, aci := range iface.ACLApplied {
			acl, err := c.GetAcl(aci.ACLApplied.Name, aci.ACLApplied.Number)
			if err != nil {
				return err
			}
			c.Interfaces[i].ACLApplied[j].ACLApplied = acl
		}
	}
	return nil

}
