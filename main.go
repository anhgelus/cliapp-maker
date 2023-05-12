package cliapp_maker

type global struct {
	Name string
	Help string
}

type CliApp struct {
	global
	Version     string
	VersionNote string
	Cmds        []Cmd
}
