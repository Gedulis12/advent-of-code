package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

type point struct {
	x, y int
}

var NOOBST = point{-1, -1}

type step struct {
	coordinate point
	direction  string
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

	input := [][]string{}
	for scanner.Scan() {
		row := []string{}
		text := scanner.Text()
		for i := range text {
			row = append(row, string(text[i]))
		}
		input = append(input, row)
	}
	start := getStart(input)
	if start.x == -1 || start.y == -1 {
		panic("start not found in a given input")
	}
	out := walk(input, start, NOOBST, "U", map[step]int{})
	for i := range out {
		for j := range out[i] {
			if out[i][j] == "V" {
				ans += 1
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

	input := [][]string{}
	for scanner.Scan() {
		row := []string{}
		text := scanner.Text()
		for i := range text {
			row = append(row, string(text[i]))
		}
		input = append(input, row)
	}
	start := getStart(input)

	if start.x == -1 || start.y == -1 {
		panic("start not found in a given input")
	}

	inputCpy := [][]string{}
	for i := range input {
		inputCpyRow := []string{}
		for j := range input[i] {
			inputCpyRow = append(inputCpyRow, input[i][j])
		}
		inputCpy = append(inputCpy, inputCpyRow)
	}

	visited := walk(input, start, NOOBST, "U", map[step]int{})


	cords := []point{}
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] == "V" {
				cord := point{x: j, y: i}
				cords = append(cords, cord)
			}
		}
	}

	ncpu := runtime.NumCPU()
	workSplits := len(cords) / ncpu
	var wg sync.WaitGroup

	for i := 0; i < ncpu; i++ {
		startCord := i*workSplits
		endCord := i*workSplits+workSplits

		if endCord > len(cords) {
			endCord = len(cords)
		}
		if i == ncpu-1 {
			endCord = len(cords)
		}

			wg.Add(1)
			go func() {
				defer wg.Done()
				for  _, cord := range(cords[startCord:endCord]) {
					if walk(input, start, cord, "U", map[step]int{}) == nil {
						ans += 1
					}
				}
			}()
		wg.Wait()
	}
//	for _, cord := range cords {
//		if walk(input, start, cord, "U", map[step]int{}) == nil {
//			ans += 1
//		}
//	}

	return ans

}

func getStart(input [][]string) point {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == "^" {
				return point{j, i}
			}
		}
	}
	return point{-1, -1}
}

func walk(input [][]string, curr, obsticle point, direction string, steps map[step]int) [][]string {
	//found exit
	if curr.x < 0 || curr.x >= len(input[0]) || curr.y < 0 || curr.y >= len(input) {
		return input
	}

	step := step{direction: direction, coordinate: curr}
	steps[step]++
	if steps[step] > 1 {
		return nil
	}

	//pre
	if input[curr.y][curr.x] == "#" || (curr == obsticle) {
		if direction == "U" {
			curr.y = curr.y + 1 // go back one step
			direction = "R"
		} else if direction == "R" {
			curr.x = curr.x - 1 // go back one step
			direction = "D"
		} else if direction == "D" {
			curr.y = curr.y - 1 // go back one step
			direction = "L"
		} else if direction == "L" {
			curr.x = curr.x + 1 // go back one step
			direction = "U"
		} else {
			fmt.Println("Direction invalid")
		}
	}
	input[curr.y][curr.x] = "V" // visited

	//recurse
	if direction == "U" {
		input = walk(input, point{curr.x, curr.y - 1}, obsticle, direction, steps)
	} else if direction == "R" {
		input = walk(input, point{curr.x + 1, curr.y}, obsticle, direction, steps)
	} else if direction == "D" {
		input = walk(input, point{curr.x, curr.y + 1}, obsticle, direction, steps)
	} else if direction == "L" {
		input = walk(input, point{curr.x - 1, curr.y}, obsticle, direction, steps)
	} else {
		fmt.Println("Direction invalid")
	}

	// post
	return input
}
