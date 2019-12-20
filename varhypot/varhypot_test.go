package varhypot_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/texttheater/golang-variadic-hypot/varhypot"
)

var testCases = []struct {
	arguments []float64
	expected  float64
}{
	{
		arguments: []float64{3, 4},
		expected:  5,
	},
	{
		arguments: []float64{5, 12},
		expected:  13,
	},
	{
		arguments: []float64{3, 4, 5},
		expected:  7.0710678118654755,
	},
	{
		arguments: []float64{-5},
		expected:  5,
	},
}

func TestVarhypot(t *testing.T) {
	for _, testCase := range testCases {
		result := varhypot.Hypot(testCase.arguments...)
		if result != testCase.expected {
			var buf bytes.Buffer
			buf.WriteString("Hypot(")
			if len(testCase.arguments) > 0 {
				buf.WriteString(fmt.Sprintf("%f", testCase.arguments[0]))
				for _, arg := range testCase.arguments[1:] {
					buf.WriteString(", ")
					buf.WriteString(fmt.Sprintf("%f", arg))
				}
			}
			buf.WriteString(") should be ")
			buf.WriteString(fmt.Sprintf("%f", testCase.expected))
			buf.WriteString(", got ")
			buf.WriteString(fmt.Sprintf("%f", result))
			t.Log(buf.String())
			t.Fail()
		}
	}
}
