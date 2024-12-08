package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

)

type equation struct {
	result int
	numbers []int
}

func main() {
	start1 := time.Now().UnixMicro()
	fmt.Println(SolvePart1("input"))
	end1 := time.Now().UnixMicro()
	fmt.Println("part 1 took: ", end1-start1)

	start2 := time.Now().UnixMicro()
	fmt.Println(SolvePart2("input"))
	end2 := time.Now().UnixMicro()
	fmt.Println("part 2 took: ", end2-start2)
}

func SolvePart1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ans := 0

	equations := []equation{}
	for scanner.Scan() {
		splitStr := strings.Split(scanner.Text(), ":")

		result, err := strconv.Atoi(splitStr[0])
		if err != nil {
			panic(err)
		}
		trimmed := strings.TrimSpace(splitStr[1])
		numbersStr := strings.Split(trimmed, " ")

		numbers := []int{}
		for i := range numbersStr {
			number, err := strconv.Atoi(numbersStr[i])
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}
		equation := equation{result: result, numbers: numbers}
		equations = append(equations, equation)
	}

	for i := range(equations) {
		if isValidEquationRTL(equations[i], "+", 0, equations[i].result, false) {
			ans += equations[i].result
			continue
		}
		if isValidEquationRTL(equations[i], "*", 0, equations[i].result, false) {
			ans += equations[i].result
			continue
		}
	}
	return ans
}

func SolvePart2(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ans := 0

	equations := []equation{}
	for scanner.Scan() {
		splitStr := strings.Split(scanner.Text(), ":")

		result, err := strconv.Atoi(splitStr[0])
		if err != nil {
			panic(err)
		}
		trimmed := strings.TrimSpace(splitStr[1])
		numbersStr := strings.Split(trimmed, " ")

		numbers := []int{}
		for i := range numbersStr {
			number, err := strconv.Atoi(numbersStr[i])
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}
		equation := equation{result: result, numbers: numbers}
		equations = append(equations, equation)
	}

	for i := range(equations) {
		if isValidEquationRTL(equations[i], "+", 0, equations[i].result, true) {
			ans += equations[i].result
			continue
		}
		if isValidEquationRTL(equations[i], "*", 0, equations[i].result, true) {
			ans += equations[i].result
			continue
		}
		if isValidEquationRTL(equations[i], "||", 0, equations[i].result, true) {
			ans += equations[i].result
			continue
		}
	}

	return ans
}

func isValidEquationLTR(e equation, op string, idx, result int, part2 bool) bool {

	if len(e.numbers)-1 == idx {
		if op == "+" {
			result = result + e.numbers[idx]
		}
		if op == "*" {
			result = result * e.numbers[idx]
		}
		if part2 && op == "||" {
			concat := fmt.Sprintf("%d%d", result, e.numbers[idx])
			concatInt, err := strconv.Atoi(concat)
			if err != nil {
				panic(err)
			}
			result = concatInt
		}
		if result == e.result {
			return true
		}
		return false
	}

	//pre
	if op == "+" {
		result = result + e.numbers[idx]
	}

	if op == "*" {
		if idx == 0 {
			result = e.numbers[idx]
		} else {
			result = result * e.numbers[idx]
		}
	}

	if part2 && op == "||" {
		concat := fmt.Sprintf("%d%d", result, e.numbers[idx])
		concatInt, err := strconv.Atoi(concat)
		if err != nil {
			panic(err)
		}
		result = concatInt
	}

	if isValidEquationLTR(e, "+", idx+1, result, part2) {
		return true
	}
	if isValidEquationLTR(e, "*", idx+1, result, part2) {
		return true
	}
	if part2 && isValidEquationLTR(e, "||", idx+1, result, part2) {
		return true
	}
	return false
}

func isValidEquationRTL(e equation, op string, idx, result int, part2 bool) bool {
	revIdx := len(e.numbers)-1-idx

	if revIdx == 0 {
		if result == e.numbers[revIdx] {
				return true
		} else {
			return false
		}
	}

	if op == "*" && result % e.numbers[revIdx] != 0 {
		return false
	}

	if part2 && op == "||" {
		resStr := fmt.Sprintf("%d", result)
		numStr := fmt.Sprintf("%d", e.numbers[revIdx])
		if len(resStr) <= len(numStr) {
			return false
		}
		lastResStr := resStr[len(resStr)-len(numStr):]
		if lastResStr != numStr {
			return false
		}
	}

	if op == "+" {
		result = result - e.numbers[revIdx]
	}

	if op == "*" {
		result = result / e.numbers[revIdx]
	}

	if part2 && op == "||" {
		resStr := fmt.Sprintf("%d", result)
		numStr := fmt.Sprintf("%d", e.numbers[revIdx])
		newResStr := resStr[:len(resStr)-len(numStr)]
		newRes, err := strconv.Atoi(newResStr)
		if err != nil {
			panic(err)
		}
		result = newRes
	}

	if isValidEquationRTL(e, "+", idx+1, result, part2) {
		return true
	}
	if isValidEquationRTL(e, "*", idx+1, result, part2) {
		return true
	}
	if part2 && isValidEquationRTL(e, "||", idx+1, result, part2) {
		return true
	}

	return false
}
