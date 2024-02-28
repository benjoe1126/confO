package mapping

import (
	"confdecl/conf"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Conf interface {
	ReadConf(fileName string) error
}

type CiscoConf struct {
	Interfaces []conf.Interface `yaml:"interfaces,flow"`
	ACLs       []conf.ACL       `yaml:"ACLS,omitempty"`
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
	fmt.Print(c)
	return nil

}
