package logic

import (
	"encoding/json"
	"fmt"
)

func FindTheMostValuablePath(input string) (float64, error) {
	// you can hard code input.

	// convert data string to array
	var array [][]float64
	err := json.Unmarshal([]byte(input), &array)
	if err != nil {
		return 0, fmt.Errorf("invalid format array")
	}
	// find the most valuable path
	var sum float64
	var currentIndexNum int
	for _, subArray := range array {
		if len(subArray) == 1 {
			sum += subArray[currentIndexNum]
			continue
		}
		if len(subArray) > 0 {
			maxNum := 0.0
			num1 := subArray[currentIndexNum]
			num2 := subArray[currentIndexNum+1]
			if num1 > num2 {
				maxNum = num1
			} else {
				maxNum = num2
				currentIndexNum += 1
			}
			sum += maxNum
		}
	}
	//fmt.Println(arrayOfMostNum)
	return sum, nil
}
