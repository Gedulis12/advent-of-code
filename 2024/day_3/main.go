package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start1 := time.Now().UnixMicro()
	fmt.Println(SolvePart1("input"))
	end1 := time.Now().UnixMicro()
	fmt.Println("part 1 took: ", end1-start1)


	start2 := time.Now().UnixMicro()
	fmt.Println(SolvePart2("input"))
	end2 := time.Now().UnixMicro()
	fmt.Println("part 1 took: ", end2-start2)
}

func SolvePart1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	ans := 0
	validExpressions := []string{}

	for scanner.Scan() {
		expr := re.FindAllString(scanner.Text(), -1)
		for _, v := range(expr) {
			validExpressions = append(validExpressions, v)
		}
	}
	for i, _ := range(validExpressions) {
		validExpressions[i] = strings.Replace(validExpressions[i], "mul(", "", -1)
		validExpressions[i] = strings.Replace(validExpressions[i], ")", "", -1)
		mults := strings.Split(validExpressions[i], ",")
		mult1, err := strconv.Atoi(mults[0])
		if err != nil {
			panic(err)
		}
		mult2, err := strconv.Atoi(mults[1])
		if err != nil {
			panic(err)
		}
		ans += (mult1*mult2)
	}
	return ans
}

func SolvePart2(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)|do\(\)|don't\(\)`)
	ans := 0
	validExpressions := []string{}

	enabled := true
	for scanner.Scan() {
		instructions := re.FindAllString(scanner.Text(), -1)
		for _, instruction := range(instructions) {
			if instruction == "do()" {
				enabled = true
				continue
			} 
			if instruction == "don't()" {
				enabled = false
			}

			if enabled {
				validExpressions = append(validExpressions, instruction)
			}
		}
	}
	for i, _ := range(validExpressions) {
		validExpressions[i] = strings.Replace(validExpressions[i], "mul(", "", -1)
		validExpressions[i] = strings.Replace(validExpressions[i], ")", "", -1)
		mults := strings.Split(validExpressions[i], ",")

		mult1, err := strconv.Atoi(mults[0])
		if err != nil {
			panic(err)
		}
		mult2, err := strconv.Atoi(mults[1])
		if err != nil {
			panic(err)
		}
		ans += mult1 * mult2
	}
	return ans
}
