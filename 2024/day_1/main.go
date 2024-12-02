package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(SolvePart1("input"))
	fmt.Println(SolvePart2("input"))
}

func SolvePart1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}
	for scanner.Scan() {

		first, err := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(strings.Split(scanner.Text(), " ")[3])
		if err != nil {
			panic(err)
		}

		left = append(left, first)
		right = append(right, second)

	}

	if len(left) != len(right) {
		return 0
	}

	slices.Sort(left)
	slices.Sort(right)

	ans := 0
	for i, _ := range(left) {
		diff := left[i] - right[i]

		if diff < 0 {
			diff = right[i] - left[i]
		}
		ans += diff
	}
	return ans
}

func SolvePart2(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}
	for scanner.Scan() {

		first, err := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(strings.Split(scanner.Text(), " ")[3])
		if err != nil {
			panic(err)
		}

		left = append(left, first)
		right = append(right, second)
	}

	freqMap := make(map[int]int)

	for _, v := range(right) {
		freqMap[v] += 1
	}
	ans := 0
	for _, v := range(left) {
		ans += (v * freqMap[v])
	}

	return ans
}