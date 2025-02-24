package logic

import (
	"fmt"
	"strconv"
)

func CatchMeLeftRightEqual(input string) (string, error) {
	result := "0"
	for idx, v := range input {
		symbol := string(v)
		switch symbol {
		case "L":
			result += "0"
			result = checkRune(input, result, idx)
		case "R":
			lastDigit, _ := strconv.Atoi(string(result[len(result)-1]))
			result += strconv.Itoa(lastDigit + 1)
		case "=":
			lastDigit := result[len(result)-1]
			result += string(lastDigit)
		default:
			return "", fmt.Errorf("invalid symbol [%s], Please input [L, R, =] only", symbol)
		}
	}
	return result, nil
}

func checkRune(input string, result string, currentIdxSymbol int) string {
	for i := currentIdxSymbol; i >= 0; i-- {
		newNum := 0
		numL, _ := strconv.Atoi(string(result[i]))
		numR, _ := strconv.Atoi(string(result[i+1]))
		symbol := string(input[i])
		switch symbol {
		case "L":
			if numL <= numR {
				newNum = numR + 1
			} else {
				newNum = numL
			}
		case "R":
			if numL >= numR {
				newNum = numL + 1
			} else {
				newNum = numL
			}
		case "=":
			newNum = numR
		default:
			continue
		}
		// อัปเดตผลลัพธ์
		result = result[:i] + strconv.Itoa(newNum) + result[i+1:]
	}
	return result
}
