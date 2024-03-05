package commands

import (
	"confdecl/commands/comfunc"
	"fmt"
)

type commands struct {
	cmd map[string]func(string) (string, error)
}

func (c commands) Get(com string) func(string) (string, error) {
	cmd, ok := c.cmd[com]
	if !ok {
		return func(str string) (string, error) {
			return fmt.Sprintf("The command %s is nonexistant, or you may have mispelled it.", com), fmt.Errorf("no valid command provided")
		}
	}
	return cmd
}

var command *commands = nil

func Commands() commands {
	if command == nil {
		mp := make(map[string]func(string) (string, error))
		mp["apply"] = comfunc.Apply
		command = &commands{
			cmd: mp,
		}
	}
	return *command
}
