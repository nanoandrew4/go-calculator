package src

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var (
	genericError               = errors.New("syntax error")
	numberParseErr             = errors.New("error parsing number")
	invalidNumOfParenthesisErr = errors.New("unclosed parenthesis")
)

func calculate(strToCalculate string) (string, error) {
	var numberStartIdx = -1
	var nextNumberNegative, lastTokenWasNumber, lastTokenWasFactorial bool
	var numbers []float64
	var ops []uint8

	for idx := 0; idx < len(strToCalculate); {
		c := strToCalculate[idx]
		if isNumber(c) {
			if lastTokenWasFactorial {
				ops = append(ops, '*')
			}

			lastTokenWasFactorial = false

			if numberStartIdx == -1 {
				numberStartIdx = idx
			}
			if len(strToCalculate) <= idx+1 || !isNumber(strToCalculate[idx+1]) {
				numToParse := strToCalculate[numberStartIdx : idx+1]
				parsedNum, err := strconv.ParseFloat(numToParse, 64)
				if err != nil {
					return "", numberParseErr
				}
				if nextNumberNegative {
					parsedNum *= -1
					nextNumberNegative = false
				}
				numbers = append(numbers, parsedNum)
				numberStartIdx = -1
				lastTokenWasNumber = true
			}
			idx++
		} else if c == '!' {
			lastNumPtr := &numbers[len(numbers)-1]
			if *lastNumPtr == 0 {
				*lastNumPtr = 1
			} else if *lastNumPtr < 0 {
				return "", genericError
			}
			for i := *lastNumPtr - 1; i > 0; i-- {
				*lastNumPtr *= i
			}
			lastTokenWasFactorial = true
			idx++
		} else if c == '|' {
			//if lastTokenWasNumber || lastTokenWasFactorial {
			//	ops = append(ops, '*')
			//}
			//
			//lastTokenWasNumber = false
			//lastTokenWasFactorial = false
			//
			//var nestedAbsoluteValues int
			//for pIdx := idx + 1; pIdx < len(strToCalculate); pIdx++ {
			//	if isNumber(strToCalculate[pIdx]) {
			//		lastTokenWasNumber = true
			//	} else if lastTokenWasNumber && strToCalculate[pIdx] == '|' {
			//		if nestedAbsoluteValues == 0 {
			//			// evaluate
			//		} else {
			//			nestedAbsoluteValues--
			//		}
			//	} else if !lastTokenWasNumber &&
			//}
			//
			//lastTokenWasNumber = false
		} else if c == '(' {
			if lastTokenWasNumber || lastTokenWasFactorial {
				ops = append(ops, '*')
			}

			lastTokenWasNumber = false
			lastTokenWasFactorial = false

			var numOfMatchedParenthesis, closingParenthesisIdx int
			for pIdx := idx + 1; pIdx < len(strToCalculate); pIdx++ {
				if strToCalculate[pIdx] == '(' {
					numOfMatchedParenthesis--
				} else if strToCalculate[pIdx] == ')' {
					numOfMatchedParenthesis++
				}
				if numOfMatchedParenthesis == 1 {
					closingParenthesisIdx = pIdx
					break
				}
			}
			if numOfMatchedParenthesis == 1 {
				resultFromParenthesis, err := calculate(strToCalculate[idx+1 : closingParenthesisIdx])
				if err != nil {
					return "", err
				}
				strToCalculate = strToCalculate[:idx] + resultFromParenthesis + strToCalculate[closingParenthesisIdx+1:]
			} else {
				return "", invalidNumOfParenthesisErr
			}
		} else if c != ' ' {
			if c == '+' || c == '-' || c == '*' || c == '/' || c == '^' {
				if c == '-' && len(numbers) == len(ops) {
					nextNumberNegative = !nextNumberNegative
				} else {
					ops = append(ops, c)
				}
				lastTokenWasNumber = false
				lastTokenWasFactorial = false
				idx++
			} else {
				return "", genericError
			}
		} else {
			idx++
		}
	}

	var result float64
	if len(numbers) != len(ops)+1 {
		return "", genericError
	}

	var secondIterOps []uint8
	secondIterNums := []float64{numbers[0]}
	for i := 1; i < len(numbers); i++ {
		if ops[i-1] == '^' {
			secondIterNums[len(secondIterNums)-1] = math.Pow(secondIterNums[len(secondIterNums)-1], numbers[i])
		} else if ops[i-1] == '*' {
			secondIterNums[len(secondIterNums)-1] = secondIterNums[len(secondIterNums)-1] * numbers[i]
		} else if ops[i-1] == '/' {
			secondIterNums[len(secondIterNums)-1] = secondIterNums[len(secondIterNums)-1] / numbers[i]
		} else {
			secondIterNums = append(secondIterNums, numbers[i])
			secondIterOps = append(secondIterOps, ops[i-1])
		}
	}

	result = secondIterNums[0]
	for i := 1; i < len(secondIterNums); i++ {
		if secondIterOps[i-1] == '+' {
			result += secondIterNums[i]
		} else if secondIterOps[i-1] == '-' {
			result -= secondIterNums[i]
		}
	}

	return formatResult(result), nil
}

func formatResult(result float64) string {
	formattedResult := fmt.Sprintf("%0.12f", result)
	var cutoff int
	for cutoff = len(formattedResult) - 1; ; {
		if cutoff > 0 && formattedResult[cutoff] == '0' {
			cutoff--
		} else if cutoff > 0 && formattedResult[cutoff] == '.' {
			cutoff--
			break
		} else {
			break
		}
	}
	return formattedResult[:cutoff+1]
}

func isNumber[T uint8 | int32](c T) bool {
	return c >= 48 && c <= 57 || c == '.'
}
