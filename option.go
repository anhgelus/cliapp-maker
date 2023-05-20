package cliapp_maker

import "go/types"

type Option struct {
	Global
	TakeValue bool
	OptType   types.BasicKind
}

type OptionPassed struct {
	Option
	Value string
}

func (o *Option) SetTakeValue(b bool) *Option {
	o.TakeValue = b
	return o
}

func (o *Option) SetType(t types.BasicKind) *Option {
	o.OptType = t
	return o
}
