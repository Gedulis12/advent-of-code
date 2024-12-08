package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type point struct {
	x, y int
}

var DIRECTIONS = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

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

	inputMap := [][]string{}
	for scanner.Scan() {
		inputRow := []string{}
		for _, v := range scanner.Text() {
			inputRow = append(inputRow, string(v))
		}
		inputMap = append(inputMap, inputRow)
	}

	// i = y, j = x
	rows := len(inputMap)
	cols := len(inputMap[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if inputMap[i][j] == "X" {
				for dir := range DIRECTIONS {
					if found := walk(inputMap, point{j, i}, "XMAS", "", []point{}, 0, DIRECTIONS[dir]); found > 0 {
						ans += found
					}
				}
			}
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

	inputMap := [][]string{}
	for scanner.Scan() {
		inputRow := []string{}
		for _, v := range scanner.Text() {
			inputRow = append(inputRow, string(v))
		}
		inputMap = append(inputMap, inputRow)
	}

	rows := len(inputMap)
	cols := len(inputMap[0])
	for i := 0; i < rows-2; i++ {
		for j := 0; j < cols-2; j++ {

			xcheck := [][]string{}

			for _, v := range inputMap[i : i+3] {
				xcheck = append(xcheck, v[j:j+3])
			}

			if xcheck[1][1] != "A" {
				continue
			}
			if !((xcheck[0][0] == "M" && xcheck[2][2] == "S") || (xcheck[0][0] == "S" && xcheck[2][2] == "M")) {
				continue
			}
			if !((xcheck[0][2] == "M" && xcheck[2][0] == "S") || (xcheck[0][2] == "S" && xcheck[2][0] == "M")) {
				continue
			}
			ans += 1
		}
	}

	// i = y, j = x
	return ans
}

func walk(input [][]string, curr point, lookFor, currentString string, path []point, found int, dir []int) int {

	// found XMAS
	if currentString == lookFor {
		return found
	}

	// off the map
	if curr.x < 0 || curr.x >= len(input[0]) || curr.y < 0 || curr.y >= len(input) {
		return -1
	}

	// not part of XMAS string
	if len(currentString) <= len(lookFor) {
		if currentString != lookFor[0:len(currentString)] {
			return -1
		}
	}
	// pre
	path = append(path, curr)
	currentString = fmt.Sprintf("%s%s", currentString, input[curr.y][curr.x])

	// recurse
	x := dir[0]
	y := dir[1]
	if f := walk(input, point{curr.x + x, curr.y + y}, lookFor, currentString, path, found, dir); f != -1 {
		return found + 1
	}

	// post
	return -1
}
