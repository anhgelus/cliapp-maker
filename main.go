package cliapp_maker

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Global struct {
	Name string
	Help string
}

type CliApp struct {
	Global
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
	simpleOpts := simpleOption.FindAllString(nCli, -1)
	var options []OptionPassed
	for _, o := range opts {
		nCli = strings.ReplaceAll(nCli, o, "")
		name := strings.ReplaceAll(strings.Split(o, " ")[0], "--", "")
		value := strings.ReplaceAll(o, name+" ", "")
		options = append(options, OptionPassed{
			Value: value,
			Option: Option{
				TakeValue: true,
				OptType:   nil,
				Global:    Global{Name: name},
			},
		})
	}
	for _, o := range simpleOpts {
		nCli = strings.ReplaceAll(nCli, o, "")
		options = append(options, OptionPassed{
			Value: "",
			Option: Option{
				TakeValue: false,
				OptType:   nil,
				Global:    Global{Name: o},
			},
		})
	}
	//TODO: parse the command and parameters
}

func (app CliApp) generateHelp() {
	println(app.Name)
	for _, cmd := range app.Cmds {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Help)
	}
}
