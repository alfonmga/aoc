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

/* func BenchmarkHey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Hey(tt.input)
		}
	}
} */
