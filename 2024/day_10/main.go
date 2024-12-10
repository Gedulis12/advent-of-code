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

var DIRECTIONS = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

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

	input := [][]int{}
	for scanner.Scan() {
		row := []int{}
		text := scanner.Text()
		for i := range text {
			num := int(text[i] - '0')
			row = append(row, num)
		}
		input = append(input, row)
	}
	trailheads := []point{}
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 0 {
				trailhead := point{x: j, y: i}
				trailheads = append(trailheads, trailhead)
			}
		}
	}
	for i := range trailheads {
		found := map[point]int{}
		walk(input, trailheads[i], 0, found)
		ans += len(found)
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

	input := [][]int{}
	for scanner.Scan() {
		row := []int{}
		text := scanner.Text()
		for i := range text {
			num := int(text[i] - '0')
			row = append(row, num)
		}
		input = append(input, row)
	}
	trailheads := []point{}
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 0 {
				trailhead := point{x: j, y: i}
				trailheads = append(trailheads, trailhead)
			}
		}
	}
	found := map[point]int{}
	for i := range trailheads {
		walk(input, trailheads[i], 0, found)
	}
	for _, v := range found {
		ans += v
	}

	return ans
}

func walk(input [][]int, curr point, next int, found map[point]int) bool {

	if curr.x < 0 || curr.x >= len(input[0]) || curr.y < 0 || curr.y >= len(input) {
		return false
	}

	if input[curr.y][curr.x] == 9 && next == 9 {
		found[curr]++
		return true
	}

	if input[curr.y][curr.x] != next {
		return false
	}

	for i := range DIRECTIONS {
		x := DIRECTIONS[i][0]
		y := DIRECTIONS[i][1]
		walk(input, point{curr.x + x, curr.y + y}, next+1, found)
	}
	return false
}