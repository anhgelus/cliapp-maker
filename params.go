package cliapp_maker

import "go/types"

type Param struct {
	Global
	ParamType types.Type
}

type ParamPassed struct {
	Name  string
	Value string
}
