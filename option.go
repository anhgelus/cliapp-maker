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
