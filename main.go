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

func (g *Global) SetName(s string) *Global {
	g.Name = s
	return g
}

func (g *Global) SetHelp(s string) *Global {
	g.Help = s
	return g
}

func (app *CliApp) Handle() {
	app.handle(os.Args)
}

func (app *CliApp) handle(args []string) {
	if len(args) == 1 {
		app.generateHelp()
		return
	}
	cli := genCli(args)
	options, nCli := parseOptions(cli)
	for _, o := range options {
		if o.Name == "v" && o.TakeValue == false {
			fmt.Printf("Version: %s\nNotes: %s", app.Version, app.VersionNote)
			return
		}
	}
	for _, cmd := range app.Cmds {
		if cmd.Name == args[1] {
			for _, o := range options {
				if o.Name == "h" && o.TakeValue == false {
					println(cmd.Help)
					return
				}
			}
			cmd.Process(CmdData{Name: cmd.Name, OptionsPassed: options, Line: cmd.genLine(args, nCli)})
			return
		}
	}
	fmt.Printf("The command %s does not exist", args[1])
	return
}

func (app *CliApp) generateHelp() {
	println(app.Name)
	for _, cmd := range app.Cmds {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Help)
	}
}

func parseOptions(cli string) ([]OptionPassed, string) {
	option := regexp.MustCompile(`--[a-zA-Z\-]+ [a-zA-Z0-9\-_]+`)
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
	return Cmd{Global: Global{Name: name}}.genLine(args, nCli)
}

func genArgsForTest(realCli string) []string {
	return strings.Split(realCli, " ")
}
