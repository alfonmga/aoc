package main

import "testing"

func TestSubmarineDiagnosticReport(t *testing.T) {
	expected := int64(230)
	actual := SubmarineDiagnosticReport(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)
	if actual[1] != expected {
		msg := `
	Result: %v
	Expected result: %v`
		t.Fatalf(msg, actual, expected)
	}
}
