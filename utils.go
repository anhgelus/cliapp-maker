package cliapp_maker

import "fmt"

// FormatHelp format the help
//
// name is the name of the command and help is the help information of the command
func FormatHelp(name string, help string) string {
	return fmt.Sprintf("%s - %s", primary(name), secondary(help))
}

// FormatStringMaxChars create a slice where each line is < max
//
// str is the string and max is the max chars per line
func FormatStringMaxChars(str string, max uint) []string {
	var formatted []string
	if uint(len(str)) > max {
		c := uint(0)
		i := uint(1)
		for uint(len(str)) > max {
			for j := i * max; j > c; j-- {
				if str[j] != ' ' && j-1 != c {
					continue
				}
				formatted = append(formatted, str[c:j])
				str = str[j:]
				i++
				c = j
			}
		}
	} else {
		formatted = append(formatted, str)
	}
	return formatted
}
