package cliapp_maker

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

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

func (app CliApp) Handle() {
	cli := ""
	for i := 1; i < len(os.Args)-2; i++ {
		cli += " " + os.Args[0]
	}
	option := regexp.MustCompile(`--[a-zA-Z\-]+ [a-zA-Z0-9 ]+`)
	simpleOption := regexp.MustCompile(`-[a-zA-Z]+`)
	opts := option.FindAllString(cli, -1)
	nCli := cli
	for _, o := range opts {
		nCli = strings.ReplaceAll(cli, o, "")
	}
	simpleOpts := option.FindAllString(nCli, -1)
	//TODO: get every options' struct
	//TODO: put them into options
	//TODO: parse the command and parameters
}

func (app CliApp) generateHelp() {
	println(app.Name)
	for _, cmd := range app.Cmds {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Help)
	}
}
