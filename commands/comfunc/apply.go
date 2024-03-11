package comfunc

import (
	"confdecl/mapping"
)

func Apply(fname string) (string, error) {
	cconf := mapping.CiscoConf{}
	err := cconf.ReadConf(fname)
	if err != nil {
		return "", err
	}
	return cconf.Configure()
}
