package mapping

import (
	"confdecl/conf"
	"gopkg.in/yaml.v2"
	"os"
)

type Conf interface {
	ReadConf(fileName string) error
}

type CiscoConf struct {
	Interfaces []conf.Interface `yaml:"interfaces,omitempty"`
	ACLs       []conf.ACL       `yaml:"ACLS,omitempty"`
}

func (c *CiscoConf) AddIface(i conf.Interface) {
	c.Interfaces = append(c.Interfaces, i)
}
func (c *CiscoConf) AddAcl(a conf.ACL) {
	c.ACLs = append(c.ACLs, a)
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
