package main

import "testing"

func TestDepthMeasurementCounter(t *testing.T) {
	actual := DepthMeasurementCounter(`199
200
208
210
200
207
240
269
260
263`)
	if actual != 7 {
		msg := `
	Result: %v
	Expected result: %v`
		t.Fatalf(msg, actual, 7)
	}
}
