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
	fmt.Println("part 2 took: ", end2-start2)
}
type point struct {
	x, y int
}
type robot struct {
	position point
	velocity point
}

var BOUNDS point = point{x: 101, y:103}

func SolvePart1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ans := 0

	robots := []robot{}
	for scanner.Scan() {
		text := scanner.Text()
		pv := strings.Split(text, "v")
		p := strings.Split(pv[0], "=")[1]
		v := strings.Split(pv[1], "=")[1]
		pxs := strings.Trim(strings.Split(p, ",")[0], " ")
		pys := strings.Trim(strings.Split(p, ",")[1], " ")
		vxs := strings.Trim(strings.Split(v, ",")[0], " ")
		vys := strings.Trim(strings.Split(v, ",")[1], " ")
		px, _ := strconv.Atoi(pxs)
		py, _ := strconv.Atoi(pys)
		vx, _ := strconv.Atoi(vxs)
		vy, _ := strconv.Atoi(vys)

		position := point{x: px, y: py}
		velocity := point{x: vx, y: vy}
		robots = append(robots, robot{position: position, velocity: velocity})
	}

	var q1, q2, q3, q4 int
	for i := 0; i < len(robots); i ++ {
		moveRobot(&robots[i], 100)
		px := robots[i].position.x
		py := robots[i].position.y
		if px <= ((BOUNDS.x-1)/2)-1 && py <= ((BOUNDS.y-1)/2)-1 {
			q1++
		}
		if px >= ((BOUNDS.x-1)/2)+1 && py <= ((BOUNDS.y-1)/2)-1 {
			q2++
		}
		if px <= ((BOUNDS.x-1)/2)-1 && py >= ((BOUNDS.y-1)/2)+1 {
			q3++
		}
		if px >= ((BOUNDS.x-1)/2)+1 && py >= ((BOUNDS.y-1)/2)+1 {
			q4++
		}
	}
	ans = 1
	if q1 != 0 {
		ans = ans * q1
	}
	if q2 != 0 {
		ans = ans * q2
	}
	if q3 != 0 {
		ans = ans * q3
	}
	if q4 != 0 {
		ans = ans * q4
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


	robots := []robot{}
	for scanner.Scan() {
		text := scanner.Text()
		pv := strings.Split(text, "v")
		p := strings.Split(pv[0], "=")[1]
		v := strings.Split(pv[1], "=")[1]
		pxs := strings.Trim(strings.Split(p, ",")[0], " ")
		pys := strings.Trim(strings.Split(p, ",")[1], " ")
		vxs := strings.Trim(strings.Split(v, ",")[0], " ")
		vys := strings.Trim(strings.Split(v, ",")[1], " ")
		px, _ := strconv.Atoi(pxs)
		py, _ := strconv.Atoi(pys)
		vx, _ := strconv.Atoi(vxs)
		vy, _ := strconv.Atoi(vys)

		position := point{x: px, y: py}
		velocity := point{x: vx, y: vy}
		robots = append(robots, robot{position: position, velocity: velocity})
	}


	idx := 0
	check := true
	for check {
		grid := [][]string{}
		for i := 0; i < BOUNDS.y; i++ {
			row := []string{}
			for j := 0; j < BOUNDS.x; j++ {
				row = append(row, ".")
			}
			grid = append(grid, row)
		}

		idx++
		for i := 0; i < len(robots); i++ {
			moveRobot(&robots[i], 1)
			grid[robots[i].position.y][robots[i].position.x] = "X"
		}
		for i := range grid {
			for j := 9; j < len(grid[i]); j++ {
				nineX := true
				for x := 0; x < 9; x++ {
					if grid[i][j-x] != "X" {
						nineX = false
						continue
					}
				}
				if nineX {
					ans = idx
					check = false
					break
				}
			}
		}
		if !check {
			for x := range grid {
				fmt.Println(grid[x])
			}
		}
	}
	return ans
}

func moveRobot(r *robot, n int) {
	xmax := BOUNDS.x
	ymax := BOUNDS.y

	for i := 0; i < n; i++ {
		newx := r.position.x + r.velocity.x
		newy := r.position.y + r.velocity.y
		if newx >= xmax {
			newx = newx - xmax
		}
		if newx < 0 {
			newx = xmax + newx
		}
		if newy >= ymax {
			newy = newy - ymax
		}
		if newy < 0 {
			newy = ymax + newy
		}
		r.position.x = newx
		r.position.y = newy
	}
}
