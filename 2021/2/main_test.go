package main

import "testing"

func TestDepthMeasurementCounter(t *testing.T) {
	expected := 5
	actual := DepthMeasurementCounter(`607
618
618
617
647
716
769
792`)
	if actual != expected {
		msg := `
	Result: %v
	Expected result: %v`
		t.Fatalf(msg, actual, expected)
	}
}

/* func BenchmarkHey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Hey(tt.input)
		}
	}
} */
