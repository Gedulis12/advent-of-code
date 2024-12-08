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

type antena struct {
	x, y int
	val string
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

	freqMap := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for i := range(line) {
			row = append(row, string(line[i]))
		}
		freqMap = append(freqMap, row)
	}

	antenas := []antena{}
	for i :=range freqMap {
		for j := range freqMap[i] {
			if freqMap[i][j] != "." {
				antena := antena{val: freqMap[i][j], x: j, y: i}
				antenas = append(antenas, antena)
			}
		}
	}
	antinodes := map[point]int{}
	dimensions := point{y: len(freqMap)-1, x: len(freqMap[0])-1}
	for i := range antenas {
		antinodes = findAntinodes(antenas[i], antenas, antinodes, dimensions)
	}
	ans = len(antinodes)

	return ans
}

func SolvePart2(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ans := 0

	freqMap := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for i := range(line) {
			row = append(row, string(line[i]))
		}
		freqMap = append(freqMap, row)
	}

	antenas := []antena{}
	for i :=range freqMap {
		for j := range freqMap[i] {
			if freqMap[i][j] != "." {
				antena := antena{val: freqMap[i][j], x: j, y: i}
				antenas = append(antenas, antena)
			}
		}
	}
	antinodes := map[point]int{}
	dimensions := point{y: len(freqMap)-1, x: len(freqMap[0])-1}
	for i := range antenas {
		antinodes = findResonantAntinodes(antenas[i], antenas, antinodes, dimensions)
	}
	ans = len(antinodes)

	return ans
}

func findAntinodes(a antena, al []antena, antinodes map[point]int, dimensions point) map[point]int{

	for i := range al {
		xDist := a.x - al[i].x
		yDist := a.y - al[i].y
		var xAnti, yAnti int

		if xDist < 0 {
			xDist = xDist * - 1
		}
		if yDist < 0 {
			yDist = yDist * - 1
		}

		if a.val != al[i].val { // Antenas missmatch
			continue
		}

		if al[i].x == a.x && al[i].y == a.y { // on source antena
			continue
		}
		if a.x < al[i].x && xDist > a.x {
			continue
		}
		if a.x > al[i].x && xDist > dimensions.x - a.x {
			continue
		}
		if a.y < al[i].y && yDist > a.y {
			continue
		}
		if a.y > al[i].y && yDist > dimensions.y - a.y {
			continue
		}

		if a.x < al[i].x {
			xAnti = a.x - xDist
		} else {
			xAnti = a.x + xDist
		}

		if a.y < al[i].y {
			yAnti = a.y - yDist
		} else {
			yAnti = a.y + yDist
		}
		antinode := point{x: xAnti, y: yAnti}
		antinodes[antinode]++
	}

	return antinodes
}

func findResonantAntinodes(a antena, al []antena, antinodes map[point]int, dimensions point) map[point]int{

	for i := range al {

		// if any other antena with the same frequency is on the map, current antena becomes antinode as well
		if al[i].val == a.val && al[i].x != a.x && al[i].y != a.y {
			antinode := point{x: a.x, y: a.y}
			antinodes[antinode]++
		}

		xDist := a.x - al[i].x
		yDist := a.y - al[i].y
		var xAnti, yAnti int

		if xDist < 0 {
			xDist = xDist * - 1
		}
		if yDist < 0 {
			yDist = yDist * - 1
		}

		if a.val != al[i].val { // Antenas missmatch
			continue
		}

		if al[i].x == a.x && al[i].y == a.y { // on source antena
			continue
		}

		// only check antenas that provide frequencies within the map
		if a.x < al[i].x && xDist > a.x {
			continue
		}
		if a.x > al[i].x && xDist > dimensions.x - a.x {
			continue
		}
		if a.y < al[i].y && yDist > a.y {
			continue
		}
		if a.y > al[i].y && yDist > dimensions.y - a.y {
			continue
		}

		idx := 1
		resonates := true
		for resonates {
			newXDist := xDist * idx
			newYDist := yDist * idx


			if a.x < al[i].x && a.x - newXDist < 0 {
				resonates = false
				continue
			}
			if a.x > al[i].x && a.x + newXDist > dimensions.x {
				resonates = false
				continue
			}

			if a.y < al[i].y && a.y - newYDist < 0 {
				resonates = false
				continue
			}
			if a.y > al[i].y && a.y + newYDist > dimensions.y {
				resonates = false
				continue
			}
			idx++

			if a.x < al[i].x {
				xAnti = a.x - newXDist
			} else {
				xAnti = a.x + newXDist
			}

			if a.y < al[i].y {
				yAnti = a.y - newYDist
			} else {
				yAnti = a.y + newYDist
			}
			antinode := point{x: xAnti, y: yAnti}
			antinodes[antinode]++
			anti1 := point{x: a.x, y: a.y}
			anti2 := point{x: al[i].x, y: al[i].y}
			antinodes[anti1]++
			antinodes[anti2]++
		}
	}
	return antinodes
}
