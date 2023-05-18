package cliapp_maker

import (
	"testing"
)

func TestCliParsing(t *testing.T) {
	realCli := "bin cmd -o --value 1 param"
	args := genArgsForTest(realCli)
	cli := genCli(args)
	if cli != "cmd -o --value 1" {
		t.Errorf("bad generation of cli\nexcepted: %s\ngot: %s", "cmd -o --value 1", cli)
	}
	opts := []OptionPassed{
		{Value: "1", Option: Option{TakeValue: true, OptType: nil, Global: Global{Name: "value"}}},
		{Value: "", Option: Option{TakeValue: false, OptType: nil, Global: Global{Name: "o"}}},
	}
	gotOpts, nCli := parseOptions(cli)
	if len(opts) != len(gotOpts) {
		for _, o := range gotOpts {
			println(o.Name, o.Value, o.TakeValue)
		}
		t.Errorf("bad parsing of options\nexcepted length: %d\ngot length: %d", len(opts), len(gotOpts))
	}
	for _, o := range gotOpts {
		valid := false
		for _, o2 := range opts {
			if o.Name == o2.Name && o.Value == o2.Value {
				valid = true
			}
		}
		if !valid {
			t.Errorf("bad parsing of options, not enough information")
		}
	}
	if nCli != "cmd" {
		t.Errorf("bad generation of new cli\nexcepted: %s\ngot: %s", "cmd", nCli)
	}
	line := genLineForTest("cmd", args, nCli)
	if line != "cmd param" {
		t.Errorf("bad generation of line\nexcepted: %s\ngot: %s", "cmd param", line)
	}
}
