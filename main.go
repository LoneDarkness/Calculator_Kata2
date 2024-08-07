package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var sign string
	var inputArray []string

	switch {
	case strings.Contains(input, " + "):
		inputArray = strings.Split(input, " + ")
		sign = "+"
	case strings.Contains(input, " - "):
		inputArray = strings.Split(input, " - ")
		sign = "-"
	case strings.Contains(input, " * "):
		inputArray = strings.Split(input, " * ")
		sign = "*"
	case strings.Contains(input, " / "):
		inputArray = strings.Split(input, " / ")
		sign = "/"
	default:
		panic("Error1")
	}

	if sign == "*" || sign == "/" {
		if strings.Contains(inputArray[1], "\"") {
			panic("Error2")
			return
		}
	}

	for i := range inputArray {
		inputArray[i] = strings.ReplaceAll(inputArray[i], "\"", "")
	}

	var strCount1 = strings.Count(inputArray[0], "")
	var strCount2 = strings.Count(inputArray[1], "")
	if strCount1 > 11 || strCount2 > 11 {
		panic("Error. Строка длинее 10 символов")
	}

	var result string
	switch sign {
	case "+":
		result = inputArray[0] + inputArray[1]
	case "*":
		multiplier, _ := strconv.Atoi(inputArray[1])
		if multiplier > 10 || multiplier < 1 {
			panic("Error. Введите число от 1 до 10")
		}
		result = ""
		for i := 0; i < multiplier; i++ {
			result += inputArray[0]
		}
	case "-":
		index := strings.Index(inputArray[0], inputArray[1])
		if index == -1 {
			result = inputArray[0]
		} else {
			result = inputArray[0][:index] + inputArray[0][index+len(inputArray[1]):]
		}
	case "/":
		div, _ := strconv.Atoi(inputArray[1])
		if div > 10 || div < 1 {
			panic("Error. Введите число от 1 до 10")
		}
		newLen := len(inputArray[0]) / div
		result = inputArray[0][:newLen]
	default:
		panic("Error2")
		return
	}

	var resultCount int
	resultCount = strings.Count(result, "")
	if resultCount > 40 {
		newresult := result[:39]
		fmt.Printf("\"%s...\"", newresult)
	} else {
		fmt.Printf("\"%s\"", result)
	}
}
