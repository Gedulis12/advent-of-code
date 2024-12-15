package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type DIR int

const (
	UP DIR = iota
	DOWN
	LEFT
	RIGHT
)

type point struct {
	x, y int
}
type box struct {
	left, right point
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

	grid := [][]string{}
	steps := []DIR{}
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "#") {
			row := []string{}
			for i := range text {
				row = append(row, string(text[i]))
			}
			grid = append(grid, row)
		} else {
			for i := range text {
				if text[i] == '^' {
					steps = append(steps, UP)
				}
				if text[i] == 'v' {
					steps = append(steps, DOWN)
				}
				if text[i] == '>' {
					steps = append(steps, RIGHT)
				}
				if text[i] == '<' {
					steps = append(steps, LEFT)
				}
			}
		}
	}

	for i := range steps {
		grid = move(grid, steps[i])
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "O" {
				ans += (100 * i) + j
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

	grid := [][]string{}
	steps := []DIR{}
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "#") {
			row := []string{}
			for i := range text {
				if text[i] == '#' {
					row = append(row, string(text[i]))
					row = append(row, string(text[i]))
				}
				if text[i] == 'O' {
					row = append(row, "[")
					row = append(row, "]")
				}
				if text[i] == '.' {
					row = append(row, string(text[i]))
					row = append(row, string(text[i]))
				}
				if text[i] == '@' {
					row = append(row, string(text[i]))
					row = append(row, ".")
				}
			}
			grid = append(grid, row)
		} else {
			for i := range text {
				if text[i] == '^' {
					steps = append(steps, UP)
				}
				if text[i] == 'v' {
					steps = append(steps, DOWN)
				}
				if text[i] == '>' {
					steps = append(steps, RIGHT)
				}
				if text[i] == '<' {
					steps = append(steps, LEFT)
				}
			}
		}
	}

	for i := range steps {
		grid = move2(grid, steps[i])
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "[" {
				ans += (100 * i) + j
			}
		}
	}

	return ans
}

func printGrid(grid [][]string) {
	for i := range grid {
		fmt.Println(grid[i])
	}
}

func getCurr(grid [][]string) point {
	curr := point{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "@" {
				curr = point{x: j, y: i}
			}
		}
	}
	return curr
}

func move(grid [][]string, dir DIR) [][]string {
	curr := getCurr(grid)
	offset := point{}
	if dir == UP {
		offset = point{x: 0, y: -1}
	}
	if dir == DOWN {
		offset = point{x: 0, y: 1}
	}
	if dir == LEFT {
		offset = point{x: -1, y: 0}
	}
	if dir == RIGHT {
		offset = point{x: 1, y: 0}
	}

	check := point{curr.x + offset.x, curr.y + offset.y}
	if grid[check.y][check.x] == "." {
		grid[check.y][check.x] = "@"
		grid[curr.y][curr.x] = "."
	}

	if grid[check.y][check.x] == "O" {
		hasSpace := false
		space := point{}
		for grid[check.y][check.x] != "#" {
			if grid[check.y][check.x] == "." {
				hasSpace = true
				space = point{x: check.x, y: check.y}
				break
			} else {
				check.x += offset.x
				check.y += offset.y
			}
		}
		if hasSpace {
			for grid[space.y][space.x] != "@" {
				grid[space.y][space.x] = grid[space.y-offset.y][space.x-offset.x]
				space.x -= offset.x
				space.y -= offset.y
			}
			grid[space.y][space.x] = "."
		}
	}
	return grid
}

func move2(grid [][]string, dir DIR) [][]string {
	curr := getCurr(grid)
	offset := point{}
	if dir == UP {
		offset = point{x: 0, y: -1}
	}
	if dir == DOWN {
		offset = point{x: 0, y: 1}
	}
	if dir == LEFT {
		offset = point{x: -1, y: 0}
	}
	if dir == RIGHT {
		offset = point{x: 1, y: 0}
	}

	check := point{curr.x + offset.x, curr.y + offset.y}
	if grid[check.y][check.x] == "." {
		grid[check.y][check.x] = "@"
		grid[curr.y][curr.x] = "."
	}

	if dir == LEFT || dir == RIGHT {
		if grid[check.y][check.x] == "[" || grid[check.y][check.x] == "]" {
			hasSpace := false
			space := point{}
			for grid[check.y][check.x] != "#" {
				if grid[check.y][check.x] == "." {
					hasSpace = true
					space = point{x: check.x, y: check.y}
					break
				} else {
					check.x += offset.x
					check.y += offset.y
				}
			}
			if hasSpace {
				for grid[space.y][space.x] != "@" {
					grid[space.y][space.x] = grid[space.y-offset.y][space.x-offset.x]
					space.x -= offset.x
					space.y -= offset.y
				}
				grid[space.y][space.x] = "."
			}
		}
	}
	if dir == UP || dir == DOWN {
		boxes := []box{}
		queue := []box{}

		if grid[check.y][check.x] == "[" {
			left := point{x: check.x, y: check.y}
			right := point{x: check.x + 1, y: check.y}
			box := box{left: left, right: right}
			boxes = append(boxes, box)
			queue = append(queue, box)
		}
		if grid[check.y][check.x] == "]" {
			left := point{x: check.x - 1, y: check.y}
			right := point{x: check.x, y: check.y}
			box := box{left: left, right: right}
			boxes = append(boxes, box)
			queue = append(queue, box)
		}
		for len(queue) > 0 {
			curr := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			// single box directly above current one
			if grid[curr.left.y+offset.y][curr.left.x] == "[" {
				left := point{x: curr.left.x, y: curr.left.y + offset.y}
				right := point{x: curr.right.x, y: curr.right.y + offset.y}
				box := box{left: left, right: right}
				boxes = append(boxes, box)
				queue = append(queue, box)
			}

			// box above the current one offset to the left
			if grid[curr.left.y+offset.y][curr.left.x] == "]" {
				left := point{x: curr.left.x - 1, y: curr.left.y + offset.y}
				right := point{x: curr.left.x, y: curr.left.y + offset.y}
				box := box{left: left, right: right}
				boxes = append(boxes, box)
				queue = append(queue, box)
			}

			// box above the current one offset to the right
			if grid[curr.right.y+offset.y][curr.right.x] == "[" {
				left := point{x: curr.right.x, y: curr.left.y + offset.y}
				right := point{x: curr.right.x + 1, y: curr.left.y + offset.y}
				box := box{left: left, right: right}
				boxes = append(boxes, box)
				queue = append(queue, box)
			}
			if grid[curr.left.y+offset.y][curr.left.x] == "#" ||
				grid[curr.right.y+offset.y][curr.right.x] == "#" {
				queue = []box{}
				boxes = []box{}
			}
		}

		// check if top (or bottom) row of boxes is not blocked by walls
		hasSpace := true
		if len(boxes) < 1 {
			hasSpace = false
		}
		for i := range boxes {
			if grid[boxes[i].left.y+offset.y][boxes[i].left.x] == "#" ||
				grid[boxes[i].right.y+offset.y][boxes[i].right.x] == "#" {
				hasSpace = false
			}
		}

		// move boxes and robot on the grid
		if hasSpace {
			for i := range boxes {
				grid[boxes[i].left.y][boxes[i].left.x] = "."
				grid[boxes[i].right.y][boxes[i].right.x] = "."
			}
			for i := range boxes {
				grid[boxes[i].left.y+offset.y][boxes[i].left.x] = "["
				grid[boxes[i].right.y+offset.y][boxes[i].right.x] = "]"
			}
			grid[curr.y][curr.x] = "."
			grid[curr.y+offset.y][curr.x] = "@"
		}
	}
	return grid
}
