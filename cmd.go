package cliapp_maker

import "fmt"

type Cmd struct {
	Global
	Options []Option
	Params  []Param
	Process func(data CmdData)
}

type CmdData struct {
	Name          string
	OptionsPassed []OptionPassed
	Line          string
}

func (cmd *Cmd) SetOptions(o []Option) *Cmd {
	cmd.Options = o
	return cmd
}

func (cmd *Cmd) SetParams(p []Param) *Cmd {
	cmd.Params = p
	return cmd
}

func (cmd *Cmd) SetProcess(fn func(data CmdData)) *Cmd {
	cmd.Process = fn
	return cmd
}

func (cmd *Cmd) GenerateHelp(app *CliApp) {
	f := FormatHelp(cmd.Name, cmd.Help)
	fLen := len(f)
	formatted := FormatStringMaxChars(f, app.CharsMax)
	str := ""
	for _, fo := range formatted {
		if fLen < len(fo) {
			fLen = len(fo)
		}
		str += fo + "\n"
	}
	app.PrintHeader(len(f))
	fmt.Println(str)
	params := ""
	for i, v := range cmd.Params {
		if i == 0 {
			params += fmt.Sprintf("%s %s", primary(v.Name), notice("("+string(rune(v.ParamType))+")"))
			continue
		}
		params += fmt.Sprintf(" %s %s", primary(v.Name), notice("("+string(rune(v.ParamType))+")"))
	}
	fmt.Println(params)
	for _, opt := range cmd.Options {
		if opt.TakeValue {
			fmt.Printf("%s %s - %s\n", primary("--"+opt.Name), notice(string(rune(opt.OptType))), secondary(opt.Help))
			continue
		}
		fmt.Printf("%s - %s\n", primary("-"+opt.Name), secondary(opt.Help))
	}
}
