package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start1 := time.Now()
	fmt.Println(SolvePart1("input"))
	fmt.Println("part 1 took: ", time.Since(start1))

	start2 := time.Now()
	fmt.Println(SolvePart2("input"))
	fmt.Println("part 1 took: ", time.Since(start2))
}

func SolvePart1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	rotations := []string{}
	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
		if err != nil {
			panic(err)
		}
	}
	ans := 0
	pos := 50
	for _, v := range rotations {
		rot, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}
		if v[0] == 'L' {
			pos -= rot
		}
		if v[0] == 'R' {
            pos += rot
		}
		if pos == 0 || pos % 100 == 0 {
			ans += 1
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

	rotations := []string{}
	for scanner.Scan() {
		rotations = append(rotations, scanner.Text())
		if err != nil {
			panic(err)
		}
	}
	ans := 0
	pos := 50
	for _, v := range rotations {
		rot, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}
		if v[0] == 'L' {
			for i := 0; i < rot; i++ {
				pos -= 1
				if pos == 0 || pos % 100 == 0 {
					ans += 1
				}
			}
		}
		if v[0] == 'R' {
			for i := 0; i < rot; i++ {
				pos += 1
				if pos == 0 || pos % 100 == 0 {
					ans += 1
				}
			}
		}
	}
	return ans
}
