package cliapp_maker

import "go/types"

type Option struct {
	Global
	TakeValue bool
	OptType   types.Type
}

type OptionPassed struct {
	Option
	Value string
}
