package cliapp_maker

import "testing"

func TestFormatter(t *testing.T) {
	excepted := primary("a") + " - " + secondary("b")
	got := FormatHelp("a", "b")
	if got != excepted {
		t.Errorf("error when formatting the help\nexcepted: %s\ngot: %s", excepted, got)
	}
}
