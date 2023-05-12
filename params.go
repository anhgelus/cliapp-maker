package cliapp_maker

import "go/types"

type Param struct {
	global
	ParamType types.Type
}

type ParamPassed struct {
	Name  string
	Value string
}
