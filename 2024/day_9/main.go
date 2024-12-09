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
	files := getFiles(input)
	emptySpace := getEmptySpace(input)

	for i := range files {
		moved := false
		for j := 0; j < len(emptySpace); j++ {
			if files[i].start < emptySpace[j].start {
				j = len(emptySpace)
				continue
			}
			if moved {
				j = len(emptySpace)
				continue
			}

			if emptySpace[j].size >= files[i].size {
				for k := 0; k < files[i].size; k++ {
					inputCpy[emptySpace[j].start+k] = inputCpy[files[i].start+k]
					inputCpy[files[i].start+k] = -1
				}
				emptySpace[j].size -= files[i].size
				emptySpace[j].start += files[i].size
				moved = true
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

func getFiles(input []int) []file {
	files := []file{}

	for i := len(input) - 1; i >= 0; i-- {
		fileStart := 0
		fileSize := 0
		if input[i] == -1 {
			continue
		}
		if len(files) > 0 && files[len(files)-1].id == input[i] {
			continue
		}
		fileStart = i
		for input[i] == input[i-fileSize] && i-fileSize > 0 {
			fileSize++
		}
		fileStart = fileStart - fileSize + 1
		file := file{id: input[fileStart], start: fileStart, size: fileSize}
		files = append(files, file)
	}
	return files
}

func getEmptySpace(input []int) []file {
	emptyBlocks := []file{}
	for j := 0; j < len(input); j++ {
		emptyStart := 0
		emptySize := 0

		if input[j] != -1 {
			continue
		}
		emptyStart = j

		for input[j+emptySize] == -1 && j+emptySize < len(input)-1 {
			emptySize++
		}
		emptyBlock := file{id: -1, size: emptySize, start: emptyStart}
		emptyBlocks = append(emptyBlocks, emptyBlock)
		j = j + emptySize
	}
	return emptyBlocks
}
