package src

import (
	"runtime/debug"
	"testing"
)

func TestCalculator(t *testing.T) {
	testsToRun := map[string]string{
		// Basic arithmetic
		"1":        "1",
		"1+2":      "3",
		"5-3":      "2",
		"3*4":      "12",
		"12/5":     "2.4",
		"12 / 5":   "2.4",
		"4/3":      "1.333333333333",
		"-1.25*40": "-50",

		// Basic operations with parenthesis + nested parenthesis
		"2 + (5 /2)":     "4.5",
		"2 + ((5 /2))":   "4.5",
		"23.5 +((5 /2))": "26",

		// Chained basic operations
		"2+3*4":       "14",
		"2+3*4/5":     "4.4",
		"2+3*4/5+2":   "6.4",
		"2+3*4/5+2-6": "0.4",

		// Negative numbers
		"-1":                       "-1",
		"-(-1)":                    "1",
		"-(-(-1))":                 "-1",
		"2 * ((5/2+2) - (-3*1.5))": "18",

		// Multiplication assumption when number precedes parenthesis
		"2(3+4)":        "14",
		"57(0.5(-0.5))": "-14.25",

		// Factorial
		"5!":           "120",
		"5!(.5)":       "60",
		"0!":           "1",
		"0!*1!*2!-3":   "-1",
		"0!1!2!-3":     "-1",
		"-(5!)":        "-120",
		"-(2!)(-(3!))": "12",

		// Powers
		"5^2":    "25",
		"25^1/2": "12.5",
		"25^0.5": "5",
		"25^-1":  "0.04",

		// Absolute value
		//"|6|": "6",
		//"|-6|": "6",
		//"|-1| + | -2 + 2 - 2 |": "3",
		//"|-1| / |2|": "0.5",
		// "||−4|*−3+2|−|6|": "4",
		//"|−1*|−4|*−3+2|−|6|": "8",
		//"|−1|−4|*−3+2|−|6|": "8", // would require lookahead
		//"|(−1)*|−4|−3+2|−|6|": "-1",

		// e
		// pi
		// mod
		// √ sqrt
		// sin
		// cos
		// tan
		// sinh
		// cosh
		// tanh
		// ln
		// log
	}

	for testStr, expectedResult := range testsToRun {
		func() {
			defer func() {
				if err := recover(); err != nil {
					debug.PrintStack()
					t.Errorf("test for expr \"%s\" panicked", testStr)
				}
			}()

			calculatedResult, err := calculate(testStr)
			if err != nil {
				t.Errorf("%s failed with error: %s", testStr, err.Error())
			}
			if calculatedResult != expectedResult {
				t.Errorf("%s expected %s - returned %s", testStr, expectedResult, calculatedResult)
			}
		}()
	}
}
