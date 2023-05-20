package cliapp_maker

import "go/types"

type Param struct {
	Global
	ParamType types.BasicKind
}

func (p *Param) SetType(t types.BasicKind) *Param {
	p.ParamType = t
	return p
}
