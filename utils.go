package cliapp_maker

import "fmt"

func FormatHelp(name string, help string) string {
	return fmt.Sprintf("%s - %s", primary(name), secondary(help))
}
