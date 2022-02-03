package main

import (
	"strings"
	"testing"
)

func TestSubmarineOxygenGeneratorRating(t *testing.T) {
	expected := int64(23)
	actual := CalcOxigenGeneratorRating(strings.Split(`00100
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
01010`, "\n"))
	if BinaryToDecimal(actual) != expected {
		msg := `
	Result: %v
	Expected result: %v`
		t.Fatalf(msg, actual, expected)
	}
}

func TestSubmarineC02ScrubberRating(t *testing.T) {
	expected := int64(10)
	actual := CalcC02ScrubberRating(strings.Split(`00100
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
01010`, "\n"))
	if BinaryToDecimal(actual) != expected {
		msg := `
	Result: %v
	Expected result: %v`
		t.Fatalf(msg, actual, expected)
	}
}

func TestSubmarineDiagnosticReportLifeSupportRating(t *testing.T) {
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
		t.Fatalf(msg, actual[1], expected)
	}
}
