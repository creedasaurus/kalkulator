package stacker

import (
	"math"
	"testing"
)

func TestProcessStatement(t *testing.T) {
	tests := []struct {
		problem string
		answer  float64
	}{
		{
			"1 + 3 + 4",
			8.0,
		},
		{
			"1++3 + -4 / 8.4",
			3.5238,
		},
		{
			"55 / 5 ^ 3 + (6 / 3)",
			2.4400,
		},
		{
			"5 * (4 /   (1 / 2)) + 99 / 1",
			139.0,
		},
	}

	for _, test := range tests {
		ans, err := ProcessStatement(test.problem)
		if err != nil {
			t.Error(err)
		}

		// check to see if number, rounded to 4 decimal places, is the same
		if math.Round(ans*10000)/10000 != test.answer {
			t.Logf("Expecting %f, got %f", test.answer, ans)
			t.Fail()
		}
		t.Log("Success")
	}
}
