package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type IDRange struct {
	start int
	end int
}

func main() {
	start1 := time.Now()
	fmt.Println(SolvePart1("input"))
	fmt.Println("part 1 took: ", time.Since(start1))

	start2 := time.Now()
	fmt.Println(SolvePart2("input"))
	fmt.Println("part 2 took: ", time.Since(start2))
}

func SolvePart1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ranges := []string{}
	for scanner.Scan() {
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, strings.Split(scanner.Text(), ",")...)
	}
	IDRanges := []IDRange{}
	for _, v := range ranges {
		if v != "" {
			idRange := strings.Split(v, "-")
			start, err := strconv.Atoi(idRange[0])
			if err != nil {
				panic(err)
			}

			end, err := strconv.Atoi(idRange[1])
			if err != nil {
				panic(err)
			}
			IDRanges = append(IDRanges, IDRange{start: start, end: end})
		}
	}
	invalidIds := []int{}
	for _, r := range IDRanges {
		for i := r.start; i <= r.end; i++ {
			id := strings.TrimLeft(strconv.Itoa(i), "0")
			if len(id) % 2 != 0 {
				continue
			}
			if id[0:len(id)/2] == id[len(id)/2:] {
				invalidId, err := strconv.Atoi(id)
				if err != nil {
					panic(err)
				}
				invalidIds = append(invalidIds, invalidId)
			}
		}
	}
	ans := 0
	for _, v := range invalidIds {
		ans += v
	}
	return ans
}

func SolvePart2(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ranges := []string{}
	for scanner.Scan() {
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, strings.Split(scanner.Text(), ",")...)
	}
	IDRanges := []IDRange{}
	for _, v := range ranges {

		if v != "" {
			idRange := strings.Split(v, "-")
			start, err := strconv.Atoi(idRange[0])
			if err != nil {
				panic(err)
			}

			end, err := strconv.Atoi(idRange[1])
			if err != nil {
				panic(err)
			}
			IDRanges = append(IDRanges, IDRange{start: start, end: end})
		}
	}
	invalidIds := []int{}
	for _, r := range IDRanges {
		for i := r.start; i <= r.end; i++ {
			id := strings.TrimLeft(strconv.Itoa(i), "0")
			for split := 1; split < len(id); split++ {
				if len(id) % split != 0 {
					continue
				}
				valid := splitsValid(id, split)
				if !valid {
					invalidIds = append(invalidIds, i)
					break
				}
			}
		}
	}
	ans := 0
	for _, v := range invalidIds {
		ans += v
	}
	return ans
}

func splitsValid(id string, splitSize int) bool {
	for i := 0; i < len(id) / splitSize; i++ {
		for j := 0; j < len(id) / splitSize; j++ {
			if j <= i {
				continue
			}
			if i*splitSize+splitSize > len(id) || j*splitSize+splitSize > len(id) {
				continue
			}
			if id[i*splitSize:i*splitSize+splitSize] != id[j*splitSize:j*splitSize+splitSize] {
				return true
			}
		}
	}
	return false
}
