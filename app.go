package main

import (
	"confdecl/conf"
	"fmt"
	"os"
)

var (
	help = map[string]string{
		"apply": "OPTIONS:\n-f <FILENAME> specify file to generate config form, must be .yml or .yaml file",
		"clear": "OPTIONS:\nall - clear all config files generated (files that end with .cconf",
	}

	commands = [...]string{
		"apply [OPTIONS][FILENAME] - generates confgis from file, use -h for further details",
		"clear [all | FILENAME] - clears configs, either all of FILENAME specified"}
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
	} else if contains("apply") {
		fmt.Print("A")
	}

	g0 := conf.NewInterface("g0/0/0")
	ipv4, _ := conf.NewIpv4("192.168.1.1", "255.255.255.0")
	g0.AddIpv4(ipv4)
	ivp4_2, _ := conf.NewIpv4("172.16.0.6", "255.252.0.0")
	ivp4_2.Print()
	g0.AddIpv4(ivp4_2)
	cnf, _ := g0.Configure()
	g1 := conf.NewInterface("g0/0/1")
	ipv4, _ = conf.NewIpv4("10.0.0.24", "255.0.0.0")
	g1.AddIpv4(ipv4)
	cnf2, _ := g1.Configure()
	fmt.Println(cnf)
	fmt.Println(cnf2)
	fmt.Print(commands[0])
}
