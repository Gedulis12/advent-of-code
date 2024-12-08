package main

import (
	"bufio"
	"fmt"
	"os"
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

	ans := 0
	for scanner.Scan() {

		report := strings.Split(scanner.Text(), " ")
		reportInts := sliceToInts(report)

		if checkSafeAsc(reportInts) || checkSafeDsc(reportInts) {
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

	ans := 0
	for scanner.Scan() {

		report := strings.Split(scanner.Text(), " ")
		reportInts := sliceToInts(report)

		if checkSafeAsc(reportInts) || checkSafeDsc(reportInts) {
			ans += 1
		} else {
			damps := getDamps(reportInts)
			if checkDamps(damps) {
				ans += 1
			}
		}
	}
	return ans
}

func sliceToInts(report []string) []int {
	reportInts := []int{}
	for i, _ := range report {
		a, err := strconv.Atoi(report[i])
		if err != nil {
			panic(err)
		}

		reportInts = append(reportInts, a)
	}
	return reportInts
}

func getDamps(report []int) [][]int {
	dampenedReports := [][]int{}
	for i, _ := range report {
		damp := make([]int, len(report))
		copy(damp, report)
		damp = append(damp[:i], damp[i+1:]...)
		dampenedReports = append(dampenedReports, damp)
	}

	return dampenedReports
}

func checkDamps(damps [][]int) bool {
	for _, v := range damps {
		if checkSafeAsc(v) || checkSafeDsc(v) {
			return true
		}
	}
	return false
}

func checkSafeAsc(report []int) bool {
	checked := 0
	for i, _ := range report {
		if i == 0 {
			continue
		}
		diff := report[i] - report[i-1]
		if diff > 0 && diff <= 3 {
			checked += 1
		}
		if checked == len(report)-1 {
			return true
		}
	}
	return false
}

func checkSafeDsc(report []int) bool {
	checked := 0
	for i, _ := range report {
		if i == 0 {
			continue
		}
		diff := report[i-1] - report[i]
		if diff > 0 && diff <= 3 {
			checked += 1
		}
		if checked == len(report)-1 {
			return true
		}
	}
	return false
}
