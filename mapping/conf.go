package mapping

import (
	"confdecl/conf/cisco"
	"confdecl/network"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Conf interface {
	ReadConf(fileName string) error
	AddIface(i network.InterfaceNetwork)
}

type CiscoConf struct {
	Interfaces    []cisco.InterfaceCisco `yaml:"interfaces,omitempty"`
	SubInterfaces []cisco.SubInterface   `yaml:"subInterfaces,omitempty"`
	ACL           []cisco.ACL            `yaml:"acl,omitempty"`
}

func (c *CiscoConf) AddIface(i network.InterfaceNetwork) {
	ciscoInt := i.(*cisco.InterfaceCisco)
	c.Interfaces = append(c.Interfaces, *ciscoInt)
}
func (c *CiscoConf) AddAcl(a cisco.ACL) {
	c.ACL = append(c.ACL, a)
}
func (c *CiscoConf) AddSubIface(s cisco.SubInterface) {
	c.SubInterfaces = append(c.SubInterfaces, s)
}

func (c *CiscoConf) ReadConf(fileName string) error {
	yamlfile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlfile, c)
	if err != nil {
		return err
	}
	return nil

}

func (c *CiscoConf) Configure() (string, error) {
	retList := make([]string, 0)
	if len(c.ACL) != 0 && c.ACL != nil {
		for _, acl := range c.ACL {
			str, err := acl.Configure()
			if err != nil {
				return "", err
			}
			retList = append(retList, str)
		}
	}
	if len(c.Interfaces) != 0 && c.Interfaces != nil {
		for _, intf := range c.Interfaces {
			str, err := intf.Configure()
			if err != nil {
				return "", err
			}
			retList = append(retList, str)
		}
	}
	return strings.Join(retList, "\n"), nil
}
