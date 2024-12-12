package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
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

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), " ")
		for i := range nums {
			num, err := strconv.Atoi(nums[i])
			if err != nil {
				panic(err)
			}
			input := []int{num}
			for i := 0; i < 25; i++ {
				input = blink(input)
			}
			ans += len(input)
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

	cache := map[int]map[int]int{}
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), " ")

		for _, n := range nums {
			num, _ := strconv.Atoi(n)
			stones := map[int]int{}
			stones[num] = 1
			stones = blink25(stones, cache)
			stones2 := blink25(stones, cache)
			stones3 := blink25(stones2, cache)
			for _, v := range stones3 {
				ans += v
			}
		}
	}

	return ans
}

// caching the map of the stone after 25 blinks
func blink25(input map[int]int, cache map[int]map[int]int) map[int]int {
	out := map[int]int{}
	for k, v := range input {
		if _, ok := cache[k]; ok {
			for ck, sv := range cache[k] {
				out[ck] += sv * v
			}
		} else {
			ci := map[int]int{}
			ci[k] = 1
			for j := 0; j < 25; j++ {
				ci = blinkToMap(ci)
			}
			cache[k] = ci
			for ck, sv := range cache[k] {
				out[ck] += sv * v
			}
		}
	}
	return out
}

// using maps
func blinkToMap(input map[int]int) map[int]int {
	out := map[int]int{}
	for k, v := range input {
		for i := 0; i < v; i++ {
			if k == 0 {
				out[1]++
				continue
			}
			numStr := strconv.Itoa(k)
			if numLen := len(numStr); numLen%2 == 0 {
				split1 := numStr[:numLen/2]
				split2 := numStr[numLen/2:]
				s1, _ := strconv.Atoi(split1)
				s2, _ := strconv.Atoi(split2)
				out[s1]++
				out[s2]++
				continue
			}
			out[k*2024]++
		}
	}
	return out
}

// naive approach:
func blink(input []int) []int {
	out := []int{}
	ncpu := runtime.NumCPU()
	workSplits := len(input) / ncpu
	var wg sync.WaitGroup
	for i := 0; i < ncpu; i++ {
		start := i * workSplits
		end := i*workSplits + workSplits

		if end > len(input) {
			end = len(input)
		}
		if i == ncpu-1 {
			end = len(input)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, v := range input[start:end] {
				if v == 0 {
					out = append(out, 1)
					continue
				}
				numStr := strconv.Itoa(v)
				if numLen := len(numStr); numLen%2 == 0 {
					split1 := numStr[:numLen/2]
					split2 := numStr[numLen/2:]
					split1Int, err := strconv.Atoi(split1)
					if err != nil {
						panic(err)
					}
					split2Int, err := strconv.Atoi(split2)
					if err != nil {
						panic(err)
					}
					out = append(out, split1Int)
					out = append(out, split2Int)
					continue
				}
				out = append(out, v*2024)
			}
		}()
		wg.Wait()
	}
	return out
}

/* Elephant graveyard below. blinkR - recursive implementation, nthBlink - itterative. Attempted to use memozation but was still too slow. Keeping as a reference for futer */
/*
func blinkR(blinks, blinksorig int, curr, start int, count *int, hm map[int]twoInts) bool {

    if blinks == 0 {
        return true
    }

    curr1 := -1
    if val, ok := hm[curr]; ok {
        curr = val.first
        curr1 = val.second
        if val.second != -1  {
            *count++
        }
    } else {
        newCurr := 0
        newCurr1 := -1
        numStr := strconv.Itoa(curr)
        numLen := len(numStr)

        if numLen % 2 != 0 && curr != 0 {
            newCurr = curr * 2024
        }

        if numLen % 2 == 0 && curr != 0 {
            numStr := strconv.Itoa(curr)
            numLen := len(numStr)
            split1 := numStr[:numLen/2]
            split2 := numStr[numLen/2:]
            newCurr, _ = strconv.Atoi(split1)
            newCurr1, _ = strconv.Atoi(split2)
            *count++
        }

        if curr == 0 {
            newCurr = 1
        }
        hm[curr] = twoInts{first: newCurr, second: newCurr1}
        curr = newCurr
        curr1 = newCurr1
    }

    if curr1 == -1 && blinkR(blinks-1, blinksorig, curr, start, count, hm) {
        return true
    }
    return blinkR(blinks-1, blinksorig, curr, start, count, hm) && blinkR(blinks-1, blinksorig, curr1, start, count, hm)
}

func nthBlink(start int, n int, hm map[int]twoInts) int {
    res := 1
    curr1 := start
    curr2 := -1
    queue := []int{}

    for i := 0; i < n; i++ {
        fmt.Println(i)
        ql := len(queue)
        for j := 0; j < ql; j++  {
            if val, ok := hm[queue[j]]; ok {
                queue[j] = val.first
                if val.second != -1  {
                    queue = append(queue, val.second)
                    res++
                }
                continue
            }

            if queue[j] == 0 {
                hm[queue[j]] = twoInts{first: 1, second: -1}
                queue[j] = 1
                continue
            }

            numStr := strconv.Itoa(queue[j])
            numLen := len(numStr)

            if numLen % 2 != 0 {
                hm[queue[j]] = twoInts{first: queue[j]*2024, second: -1}
                queue[j] = queue[j] * 2024
                continue
            }

            if numLen % 2 == 0 && queue[j] != 0 {
                numStr := strconv.Itoa(queue[j])
                numLen := len(numStr)
                split1 := numStr[:numLen/2]
                split2 := numStr[numLen/2:]
                s1, _ := strconv.Atoi(split1)
                s2, _ := strconv.Atoi(split2)
                hm[queue[j]] = twoInts{first: s1, second: s2}
                queue[j] = s1
                queue = append(queue, s2)
                res++
            }
        }

        if val, ok := hm[curr1]; ok {
            curr1 = val.first
            curr2 = val.second
            if val.second != -1  {
                res++
                queue = append(queue, val.second)
            }
            continue
        }
        if curr1 == 0 {
            hm[curr1] = twoInts{first: 1, second: -1}
            curr1 = 1
            continue
        }
        numStr := strconv.Itoa(curr1)
        numLen := len(numStr)
        if numLen % 2 != 0 && curr1 != 0 {
            hm[curr1] = twoInts{first: curr1*2024, second: -1}
            curr1 = curr1 * 2024
            continue
        }
        if numLen % 2 == 0 && curr1 != 0 {
            numStr := strconv.Itoa(curr1)
            numLen := len(numStr)
            split1 := numStr[:numLen/2]
            split2 := numStr[numLen/2:]
            c1, _ := strconv.Atoi(split1)
            c2, _ := strconv.Atoi(split2)
            hm[curr1] = twoInts{first: c1, second: c2}
            curr1 = c1
            curr2 = c2
            queue = append(queue, curr2)
            res++
        }
    }
    return res
}
*/