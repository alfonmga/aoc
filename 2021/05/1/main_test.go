package main

import "testing"

func TestCalcPointLinesOverlap(t *testing.T) {
	expected := 5
	actual := CalcPointLinesOverlap(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`)
	if actual != expected {
		msg := `
	Result: %v
	Expected result: %v`
		t.Fatalf(msg, actual, expected)
	}
}
