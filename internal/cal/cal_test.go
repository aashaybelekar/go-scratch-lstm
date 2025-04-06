package cal_test

import (
	"math"
	"testing"

	"github.com/aashaybelekar/go-scratch-lstm/internal/cal"
)

func floatEqual(a, b, tolerance float64) bool {
	return float64(math.Abs(float64(a-b))) <= tolerance
}

func TestSigmoid(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 0.5},
		{1, 1 / (1 + float64(math.Pow(math.E, 1)))},
		{-1, 1 / (1 + float64(math.Pow(math.E, -1)))},
	}

	for _, tt := range tests {
		got := cal.Sigmoid(tt.input)
		if !floatEqual(got, tt.want, 0.0001) {
			t.Errorf("sigmoid(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
