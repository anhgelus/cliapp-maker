package cliapp_maker

import "go/types"

type Option struct {
	Global
	TakeValue bool
	OptType   types.BasicKind
	// Process is the function called when the option is called
	// Return false if the command must be stopped
	// Needed for global options
	Process func(data *OptionPassed) bool
}

type OptionPassed struct {
	Option
	Value string
	Context
}

func (o *Option) SetTakeValue(b bool) *Option {
	o.TakeValue = b
	return o
}

func (o *Option) SetType(t types.BasicKind) *Option {
	o.OptType = t
	return o
}

func (o *Option) SetProcess(fn func(data *OptionPassed) bool) *Option {
	o.Process = fn
	return o
}
