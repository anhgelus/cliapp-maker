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
	ParamsPassed  []ParamPassed
}

func (cmd Cmd) generateHelp() {
	fmt.Printf("%s - %s\n", cmd.Name, cmd.Help)
	params := ""
	for _, v := range cmd.Params {
		params += " " + v.Name + " (" + v.ParamType.String() + ")"
	}
	fmt.Printf("%s - %s\n", cmd.Name, params)
	for _, opt := range cmd.Options {
		if opt.TakeValue {
			fmt.Printf("--%s %s - %s\n", opt.Name, opt.OptType.String(), opt.Help)
			continue
		}
		fmt.Printf("-%s - %s\n", opt.Name, opt.Help)
	}
}
