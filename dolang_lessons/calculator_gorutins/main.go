package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	inputString := "2+2, 3+6, 7*7, 9/3"
	var mChan = make(chan string)
	operationSlice := strings.Split(inputString, ",")
	for _, element := range operationSlice {
		go calcResult(element, mChan)
	}
	for i := 0; i < len(operationSlice); i++ {
		message := <-mChan
		fmt.Println(message)
	}
}

func calcResult(numbers string, mChan chan string) {
	if res, err := calculate(numbers); err == nil {
		mChan <- fmt.Sprintf("result of %s is %d", numbers, res)
	} else {
		mChan <- fmt.Sprintf("result of %s is error %e", numbers, err)
	}
}

func calculate(numbers string) (int, error) {
	operands := strings.FieldsFunc(numbers, arithmeticSigns)
	if len(operands) < 2 {
		return math.MinInt, nil
	}

	return getCalcResult(&operands, getSignOperation(getSign(numbers, &operands)))
}

func getSign(operation string, operators *[]string) rune {
	return []rune(operation)[len((*operators)[0])]

}

func getCalcResult(operators *[]string, signOperation func(arr *[]string) (int, error)) (int, error) {
	return signOperation(operators)
}

func getSignOperation(sign rune) func(arr *[]string) (int, error) {
	switch sign {
	case '+':
		return func(arr *[]string) (int, error) {

			val1, err := strconv.Atoi(strings.Trim((*arr)[0], " "))
			if err != nil {
				return 0, err
			}
			val2, err := strconv.Atoi(strings.Trim((*arr)[1], " "))
			if err != nil {
				return 0, err
			}
			return val1 + val2, err

		}
	case '-':
		return func(arr *[]string) (int, error) {

			val1, err := strconv.Atoi(strings.Trim((*arr)[0], " "))
			if err != nil {
				return 0, err
			}
			val2, err := strconv.Atoi(strings.Trim((*arr)[1], " "))
			if err != nil {
				return 0, err
			}
			return val1 - val2, err

		}
	case '*':
		return func(arr *[]string) (int, error) {

			val1, err := strconv.Atoi(strings.Trim((*arr)[0], " "))
			if err != nil {
				return 0, err
			}
			val2, err := strconv.Atoi(strings.Trim((*arr)[1], " "))
			if err != nil {
				return 0, err
			}
			return val1 * val2, err

		}
	case '/':
		return func(arr *[]string) (int, error) {

			val1, err := strconv.Atoi(strings.Trim((*arr)[0], " "))
			if err != nil {
				return 0, err
			}
			val2, err := strconv.Atoi(strings.Trim((*arr)[1], " "))
			if err != nil {
				return 0, err
			}

			return val1 / val2, err

		}

	}
	return nil
}

func arithmeticSigns(r rune) bool {
	signs := "/*-+"
	return strings.Index(signs, string(r)) != -1
}
