package cliapp_maker

type Cmd struct {
	global
	Options []Option
	Params  []Param
	Process func(data CmdData)
}

type CmdData struct {
	Name          string
	OptionsPassed []OptionPassed
	ParamsPassed  []ParamPassed
}
