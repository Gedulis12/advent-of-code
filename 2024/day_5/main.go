package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	fmt.Println("part 2 took: ", end2-start2)
}

func SolvePart1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ans := 0

	rules := [][]int{}
	updates := [][]int{}

	for scanner.Scan() {
		if text := scanner.Text(); strings.Contains(text, "|") {
			rule := strings.Split(text, "|")

			rule_1, _ := strconv.Atoi(rule[0])
			if err != nil {
				panic(err)
			}

			rule_2, err := strconv.Atoi(rule[1])
			if err != nil {
				panic(err)
			}

			rules = append(rules, []int{rule_1, rule_2})
			continue
		}
		if scanner.Text() == "\n" || scanner.Text() == "" {
			continue
		}
		update := strings.Split(scanner.Text(), ",")

		updateInts := []int{}
		for _, v := range update {
			i, err := strconv.Atoi(v)

			if err != nil {
				panic(err)
			}
			updateInts = append(updateInts, i)
		}
		updates = append(updates, updateInts)
	}
	goodUpdates := getGoodUpdates(rules, updates)

	for _, u := range goodUpdates {
		idx := (len(u) - 1) / 2
		ans += u[idx]
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

	rules := [][]int{}
	updates := [][]int{}

	for scanner.Scan() {
		if text := scanner.Text(); strings.Contains(text, "|") {
			rule := strings.Split(text, "|")

			rule_1, _ := strconv.Atoi(rule[0])
			if err != nil {
				panic(err)
			}

			rule_2, err := strconv.Atoi(rule[1])
			if err != nil {
				panic(err)
			}

			rules = append(rules, []int{rule_1, rule_2})
			continue
		}
		if scanner.Text() == "\n" || scanner.Text() == "" {
			continue
		}
		update := strings.Split(scanner.Text(), ",")

		updateInts := []int{}
		for _, v := range update {
			i, err := strconv.Atoi(v)

			if err != nil {
				panic(err)
			}
			updateInts = append(updateInts, i)
		}
		updates = append(updates, updateInts)
	}
	badUpdates := getBadUpdates(rules, updates)
	fixedUpdates := fixBadUpdates(rules, badUpdates)

	for _, u := range fixedUpdates {
		idx := (len(u) - 1) / 2
		ans += u[idx]
	}

	return ans
}

func checkUpdate(rules [][]int, update []int) bool {
	isGood := true
	for _, rule := range rules {
		if slices.Contains(update, rule[0]) && slices.Contains(update, rule[1]) {
			if slices.Index(update, rule[0]) > slices.Index(update, rule[1]) {
				isGood = false
			}
		}
	}
	return isGood
}

func getGoodUpdates(rules, updates [][]int) [][]int {
	goodUpdates := [][]int{}
	for _, update := range updates {
		if checkUpdate(rules, update) {
			goodUpdates = append(goodUpdates, update)
		}
	}
	return goodUpdates
}

func getBadUpdates(rules, updates [][]int) [][]int {
	badUpdates := [][]int{}
	for _, update := range updates {
		if !checkUpdate(rules, update) {
			badUpdates = append(badUpdates, update)
		}
	}
	return badUpdates
}

func swap(update []int, s1, s2 int) []int {
	temp := update[s1]
	update[s1] = update[s2]
	update[s2] = temp
	return update
}

func fixUpdate(rules [][]int, update []int) []int {
	updateCpy := make([]int, len(update))
	copy(updateCpy, update)
	for _, rule := range rules {
		if slices.Contains(updateCpy, rule[0]) && slices.Contains(updateCpy, rule[1]) {
			if slices.Index(updateCpy, rule[0]) > slices.Index(updateCpy, rule[1]) {
				updateCpy = swap(updateCpy, slices.Index(updateCpy, rule[0]), slices.Index(updateCpy, rule[1]))
			}
		}
	}
	return updateCpy
}

func fixBadUpdates(rules, updates [][]int) [][]int {
	fixedUpdates := [][]int{}
	for _, update := range updates {
		updateCpy := make([]int, len(update))
		copy(updateCpy, update)

		c := checkUpdate(rules, updateCpy)
		for c != true {
			updateCpy = fixUpdate(rules, updateCpy)
			c = checkUpdate(rules, updateCpy)
		}
		fixedUpdates = append(fixedUpdates, updateCpy)
	}
	return fixedUpdates
}
