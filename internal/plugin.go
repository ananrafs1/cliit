package internal

import "github.com/ananrafs1/cliit/plugin"

func GetPluginList() map[string]plugin.Executor {
	return map[string]plugin.Executor{
		"Calculator": plugin.ExecDummy1{},
		"Mapping":    plugin.ExecDummy2{},
	}
}
