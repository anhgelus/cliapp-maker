package cliapp_maker

import "fmt"

type global struct {
	Name string
	Help string
}

type CliApp struct {
	global
	Version     string
	VersionNote string
	Cmds        []Cmd
}

func (app CliApp) generateHelp() {
	println(app.Name)
	for _, cmd := range app.Cmds {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Help)
	}
}
