package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type file struct {
	id    int
	start int
	size  int
}

func main() {
	start1 := time.Now().UnixMicro()
	fmt.Println(SolvePart1("example"))
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

	for scanner.Scan() {
		text := scanner.Text()
		input := []int{}
		for i := range text {
			input = append(input, int(text[i]-'0'))
		}
		decoded := decode(input)
		compacted := compact(decoded)
		ans = checksum(compacted)
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

	for scanner.Scan() {
		text := scanner.Text()
		input := []int{}
		for i := range text {
			input = append(input, int(text[i]-'0'))
		}
		decoded := decode(input)
		defraged := defrag(decoded)
		ans = checksum(defraged)
	}
	return ans
}

func decode(input []int) []int {
	var out []int
	var expanded int
	idx := 0
	for i := range input {
		if ((i + 1) % 2) == 0 {
			expanded = -1
		} else {
			expanded = idx
			idx++
		}
		repeat := input[i]

		for j := 0; j < repeat; j++ {
			out = append(out, expanded)
		}
	}
	return out
}

func compact(input []int) []int {
	inputCpy := input
	out := []int{}
	idx := len(inputCpy) - 1
	for i := range input {
		if idx < i {
			break
		}
		if input[i] != -1 {
			out = append(out, input[i])
		} else {
			endByte := input[idx]
			for endByte == -1 {
				idx--
				endByte = input[idx]
			}
			out = append(out, endByte)
			idx--
		}
	}
	return out
}

func defrag(input []int) []int {
	inputCpy := input
	files := []file{}

	for i := len(inputCpy) - 1; i >= 0; i-- {
		fileStart := 0
		fileSize := 0
		if inputCpy[i] == -1 {
			continue
		}
		if len(files) > 0 && files[len(files)-1].id == inputCpy[i] {
			continue
		}
		fileStart = i
		for inputCpy[i] == inputCpy[i-fileSize] && i-fileSize > 0 {
			fileSize++
		}
		fileStart = fileStart - fileSize + 1
		file := file{id: inputCpy[fileStart], start: fileStart, size: fileSize}
		files = append(files, file)
	}

	for i := range files {
		moved := false
		for j := 0; j < len(input); j++ {
			emptyStart := 0
			emptySize := 0

			if files[i].start < j {
				j = len(input)
				continue
			}

			if moved {
				j = len(input)
				continue
			}

			if inputCpy[j] != -1 {
				continue
			}
			emptyStart = j

			for inputCpy[j+emptySize] == -1 && j+emptySize < len(inputCpy)-1 {
				emptySize++
			}

			if emptySize >= files[i].size && files[i].start > emptyStart {
				for k := 0; k < files[i].size; k++ {
					inputCpy[emptyStart+k] = inputCpy[files[i].start+k]
					inputCpy[files[i].start+k] = -1
				}
				moved = true
			} else {
				emptyStart = 0
				emptySize = 0
			}
		}
	}
	return inputCpy
}

func checksum(c []int) int {
	out := 0
	for i := range c {
		if c[i] == -1 {
			continue
		}
		out += (i * c[i])
	}
	return out
}