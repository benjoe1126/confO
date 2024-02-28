package mapping

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Conf interface {
	readConf(fileName string) error
}

type CiscoConf struct {
}

func (c *CiscoConf) readConf(fileName string) error {
	yamlfile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	yamline := yaml_v2.yaml.Unmarshall
}
