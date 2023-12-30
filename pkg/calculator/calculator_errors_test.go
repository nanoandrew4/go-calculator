package calculator

import (
	"errors"
	"testing"
)

func TestCalculatorErrors(t *testing.T) {
	testsToRun := map[string]error{
		"":            genericError,
		"+":           genericError,
		"1+":          genericError,
		"123/12(":     invalidNumOfParenthesisErr,
		"123/12)":     genericError,
		"123/(12+2))": genericError,

		"-5!":    genericError,
		"-2!-3!": genericError,
	}

	for testStr, expectedErr := range testsToRun {
		calculatedResult, err := Calculate(testStr)
		if calculatedResult != "" {
			t.Fatalf("%s expected not to return anything, returned: %s", testStr, calculatedResult)
		}
		if !errors.Is(err, expectedErr) {
			t.Fatalf("%s expected %s - returned %s", testStr, expectedErr, err)
		}
	}
}
