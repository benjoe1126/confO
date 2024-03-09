package main

import (
	"confdecl/commands"
	"confdecl/utils"
	"fmt"
	"os"
)

var (
	help = map[string]string{
		"apply": "OPTIONS:\n-f <FILENAME> specify file to generate config form, must be .yml or .yaml file",
		"clear": "OPTIONS:\nall - clear all config files generated (files that end with .cconf",
	}

	/*commands = [...]string{
	"apply [OPTIONS][FILENAME] - generates confgis from file, use -h for further details",
	"clear [all | FILENAME] - clears configs, either all of FILENAME specified"}
	*/
)

func contains(str string) bool {
	for _, el := range os.Args {
		if str == el {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Errorf("%s\n", "No command provided, please provide one of the following commands")
	} else {
		str := commands.Commands().Get(os.Args[1])
		strin, err := str(os.Args[2])
		if err != nil {
			panic(err)
		}
		fmt.Println(strin)
		fmt.Println(utils.PrefixToDottedDecimal(17))
	} /*
		g0 := cisco.NewInterface("g0/0/0")
		ipv4, _ := network.NewIpv4("192.168.1.1", "255.255.255.0")
		g0.AddIpv4(ipv4)
		ivp4_2, _ := network.NewIpv4("172.16.0.6", "255.252.0.0")
		g0.AddIpv4(ivp4_2)
		//cnf, _ := g0.Configure()
		g1 := cisco.NewInterface("g0/0/1")
		ipv4, _ = network.NewIpv4("10.0.0.24", "255.0.0.0")
		g1.AddIpv4(ipv4)
		//cnf2, _ := g1.Configure()
		//fmt.Println(cnf)
		cconf := mapping.CiscoConf{}
		cconf.AddIface(&g0)
		cconf.AddIface(&g1)
		sub, _ := cisco.NewSubinterface("g0/0/1.2", 2)
		str, err := sub.Configure()
		if err != nil {
			panic(err)
		}
		strr, _ := yaml.Marshal(cconf)
		fmt.Println(string(strr))
		fmt.Println(str)
		cconf.AddSubIface(sub)

		str, err = g0.Configure()
		fmt.Println(str)*/

}

//TODO megy snmp, lehetne több command, get int desc és egyéb shitek :D
//TODO csak fellépni eszközökre és configokat kiadni ott, would be very nice soonTM
//TODO notice me senpai >:333 UwU CoC ^w^ OuO (O<~~>O) §^§
