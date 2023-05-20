package cliapp_maker

import (
	"fmt"
	"go/types"
	"testing"
)

var errorDuringTest error

func TestCliParsing(t *testing.T) {
	realCli := "bin cmd -o --value 1 param"
	args := genArgsForTest(realCli)
	cli := genCli(args)
	if cli != "cmd -o --value 1" {
		t.Errorf("bad generation of cli\nexcepted: %s\ngot: %s", "cmd -o --value 1", cli)
	}
	opts := []OptionPassed{
		{Value: "1", Option: Option{TakeValue: true, Global: Global{Name: "value"}}},
		{Value: "", Option: Option{TakeValue: false, Global: Global{Name: "o"}}},
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
			println(o.Name, o.Value)
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

func TestBasicApp(t *testing.T) {
	app := CliApp{}
	// Global
	app.SetName("Test App").SetHelp("Just a test app")
	// Version
	app.SetVersion("1.0.0").SetVersionNote("Test app version")
	// Commands
	test := Cmd{}
	test.SetName("test").SetHelp("Test")
	test.SetOptions([]Option{
		{
			Global: Global{
				Name: "test",
				Help: "Test",
			},
			TakeValue: true,
			OptType:   types.String,
		},
	}).SetParams([]Param{
		{
			Global: Global{
				Name: "try",
				Help: "Try 1",
			},
			ParamType: types.String,
		},
		{
			Global: Global{
				Name: "hello",
				Help: "Try 2",
			},
			ParamType: types.String,
		},
	})
	test.SetProcess(processTest)
	app.SetCommands([]Cmd{test})

	realCli := "bin test try1 try2"
	app.handle(genArgsForTest(realCli))
	testError(t)
	realCli = "bin -v test try1 try2"
	app.handle(genArgsForTest(realCli))
	testError(t)
	realCli = "bin test -v try1 try2"
	app.handle(genArgsForTest(realCli))
	testError(t)
	realCli = "bin test try1 -v try2"
	app.handle(genArgsForTest(realCli))
	testError(t)
	realCli = "bin test try1 --test hello try2"
	app.handle(genArgsForTest(realCli))
	testError(t)
	realCli = "bin test try1 --test hello -v try2"
	app.handle(genArgsForTest(realCli))
	testError(t)
}

func testError(t *testing.T) {
	if errorDuringTest != nil {
		t.Errorf(errorDuringTest.Error())
	}
	errorDuringTest = nil
}

func processTest(data CmdData) {
	if data.Line != "try1 try2" {
		errorDuringTest = fmt.Errorf("bad line\nexcepted: %s\ngot: %s", "try1 try2", data.Line)
		return
	}
	if data.Name != "test" {
		errorDuringTest = fmt.Errorf("bad name\nexcepted: %s\ngot: %s", "test", data.Name)
		return
	}
	for _, o := range data.OptionsPassed {
		if o.Name != "test" {
			return
		}
		if o.TakeValue == false {
			errorDuringTest = fmt.Errorf("bad option passed\nexcepted: %s\ngot: %s", "true", "false")
			return
		}
		if o.Value != "hello" {
			errorDuringTest = fmt.Errorf("bad option value passed\nexcepted: %s\ngot: %s", "hello", o.Value)
			return
		}
	}
}
