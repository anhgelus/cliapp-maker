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
	if len(os.Args) == 1 {
		app.generateHelp()
		os.Exit(0)
	}
	cli := genCli(os.Args)
	options, nCli := parseOptions(cli)
	for _, cmd := range app.Cmds {
		if cmd.Name == os.Args[1] {
			cmd.Process(CmdData{Name: cmd.Name, OptionsPassed: options, Line: cmd.genLine(os.Args, nCli)})
			os.Exit(0)
		}
	}
	fmt.Printf("The command %s does not exist", os.Args[1])
	os.Exit(1)
}

func parseOptions(cli string) ([]OptionPassed, string) {
	option := regexp.MustCompile(`--[a-zA-Z\-]+ [a-zA-Z0-9\-_]`)
	simpleOption := regexp.MustCompile(`-[a-zA-Z]+`)
	opts := option.FindAllString(cli, -1)
	nCli := cli
	var options []OptionPassed
	for _, o := range opts {
		nCli = strings.ReplaceAll(nCli, " "+o, "")
		name := strings.ReplaceAll(strings.Split(o, " ")[0], "--", "")
		value := strings.ReplaceAll(o, "--"+name+" ", "")
		options = append(options, OptionPassed{
			Value: value,
			Option: Option{
				TakeValue: true,
				OptType:   nil,
				Global:    Global{Name: name},
			},
		})
	}
	simpleOpts := simpleOption.FindAllString(nCli, -1)
	for _, o := range simpleOpts {
		nCli = strings.ReplaceAll(nCli, " "+o, "")
		options = append(options, OptionPassed{
			Value: "",
			Option: Option{
				TakeValue: false,
				OptType:   nil,
				Global:    Global{Name: strings.Replace(o, "-", "", 1)},
			},
		})
	}
	return options, nCli
}

func genCli(args []string) string {
	cli := args[1]
	for i := 2; i < len(args)-1; i++ {
		cli += " " + args[i]
	}
	return cli
}

func (cmd Cmd) genLine(args []string, nCli string) string {
	return strings.ReplaceAll(nCli, cmd.Name+" ", "") + " " + args[len(args)-1]
}

func genLineForTest(name string, args []string, nCli string) string {
	return strings.ReplaceAll(nCli, name+" ", "") + " " + os.Args[len(os.Args)-1]
}

func genArgsForTest(realCli string) []string {
	return strings.Split(realCli, " ")
}

func (app CliApp) generateHelp() {
	println(app.Name)
	for _, cmd := range app.Cmds {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Help)
	}
}
