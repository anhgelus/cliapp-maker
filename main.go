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

type Context struct {
	App *CliApp
	// CmdCalled can be nil if there is no command
	CmdCalled *Cmd
}

var (
	globalOptions = make([]Option, 2)
)

func init() {
	help := Option{}
	help.SetTakeValue(false).SetProcess(handleHelp).SetName("h").SetHelp("Show the help")
	version := Option{}
	version.SetTakeValue(false).SetProcess(handleVersion).SetName("v").SetHelp("Show the version")
	globalOptions = append(globalOptions, help, version)
}

// AddGlobalOption add an option to the global option's array
func AddGlobalOption(o Option) {
	globalOptions = append(globalOptions, o)
}

func handleVersion(data *OptionPassed) bool {
	fmt.Printf("Version: %s\nNotes: %s", data.Context.App.Version, data.Context.App.VersionNote)
	return false
}

func handleHelp(data *OptionPassed) bool {
	if data.CmdCalled != nil {
		data.App.generateHelp()
	} else {
		data.CmdCalled.generateHelp()
	}
	return false
}

func (g *Global) SetName(s string) *Global {
	g.Name = s
	return g
}

func (g *Global) SetHelp(s string) *Global {
	g.Help = s
	return g
}

func (app *CliApp) SetVersion(s string) *CliApp {
	app.Version = s
	return app
}

func (app *CliApp) SetVersionNote(s string) *CliApp {
	app.VersionNote = s
	return app
}

func (app *CliApp) SetCommands(cmds []Cmd) *CliApp {
	app.Cmds = cmds
	return app
}

func (app *CliApp) Handle() {
	app.handle(os.Args)
}

func (app *CliApp) handle(args []string) error {
	if len(args) == 1 {
		app.generateHelp()
		return nil
	}
	cli := genCli(args)
	options, nCli := parseOptions(cli)

	for _, cmd := range app.Cmds {
		if cmd.Name == args[1] {
			if !app.handleOptions(&cmd, options) {
				return nil
			}
			cmd.Process(CmdData{Name: cmd.Name, OptionsPassed: options, Line: cmd.genLine(args, nCli)})
			return nil
		}
	}
	if !app.handleOptions(nil, options) {
		return nil
	}
	fmt.Printf("The command %s does not exist", args[1])
	return fmt.Errorf("The command %s does not exist", args[1])
}

func (app *CliApp) handleOptions(cmd *Cmd, opts []OptionPassed) bool {
	for _, opt := range opts {
		opt.Context = Context{App: app, CmdCalled: cmd}
		for _, o := range globalOptions {
			if o.Name == opt.Name && o.TakeValue == opt.TakeValue {
				return o.Process(&opt)
			}
		}
	}
	return false
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

func (cmd *Cmd) genLine(args []string, nCli string) string {
	return strings.ReplaceAll(nCli, cmd.Name+" ", "") + " " + args[len(args)-1]
}

func genLineForTest(name string, args []string, nCli string) string {
	cmd := Cmd{Global: Global{Name: name}}
	return cmd.genLine(args, nCli)
}

func genArgsForTest(realCli string) []string {
	return strings.Split(realCli, " ")
}
