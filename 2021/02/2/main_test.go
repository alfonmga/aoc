package main

import "testing"

func TestSubmarineMovementCmdHandler(t *testing.T) {
	expected := 900
	actual := SubmarineMovementCmdHandler(`forward 5
down 5
forward 8
up 3
down 8
forward 2`)
	if actual != expected {
		msg := `
	Result: %v
	Expected result: %v`
		t.Fatalf(msg, actual, expected)
	}
}
