package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"time"
)

type face int

const (
	NORTH face = iota
	EAST
	SOUTH
	WEST
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

type region struct {
	val    string
	area   int
	edges  int
	fences []fence
}

type fence struct {
	point point
	face  face
}
type point struct {
	x, y int
}

func SolvePart1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ans := 0

	input := [][]string{}
	visited := [][]bool{}
	region := region{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		vRow := []bool{}
		for i := range line {
			row = append(row, string(line[i]))
			vRow = append(vRow, false)
		}
		input = append(input, row)
		visited = append(visited, vRow)
	}
	for i := range input {
		for j := range input[i] {
			if visited[i][j] {
				continue
			}
			region, visited = bfs(input, point{j, i}, visited)
			ans += region.area * region.edges
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
	visited := [][]bool{}
	region := region{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		vRow := []bool{}
		for i := range line {
			row = append(row, string(line[i]))
			vRow = append(vRow, false)

		}
		input = append(input, row)
		visited = append(visited, vRow)
	}
	for i := range input {
		for j := range input[i] {
			if visited[i][j] {
				continue
			}
			region, visited = bfs(input, point{j, i}, visited)
			sides := getAllSides(region)
			ans += region.area * sides
		}
	}

	return ans
}

func bfs(input [][]string, curr point, visited [][]bool) (region, [][]bool) {
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	region := region{}
	region.val = input[curr.y][curr.x]

	queue := []point{}
	queue = append(queue, curr)

	runBfs := true
	for runBfs {
		if len(queue) == 0 {
			runBfs = false
			continue
		}
		curr = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if visited[curr.y][curr.x] {
			continue
		} else {
			visited[curr.y][curr.x] = true
		}
		region.area++

		for i := range directions {
			x := directions[i][1]
			y := directions[i][0]
			newX := curr.x + x
			newY := curr.y + y
			oob := false

			if newX < 0 || newX >= len(input[0]) {
				point := point{x: newX, y: newY}
				f := fence{}
				if newX < 0 {
					f = fence{point: point, face: WEST}
				} else {
					f = fence{point: point, face: EAST}
				}
				region.fences = append(region.fences, f)
				oob = true
			}
			if newY < 0 || newY >= len(input) {
				point := point{x: newX, y: newY}
				f := fence{}
				if newY < 0 {
					f = fence{point: point, face: NORTH}
				} else {
					f = fence{point: point, face: SOUTH}
				}
				region.fences = append(region.fences, f)
				oob = true
			}

			if oob {
				region.edges++
				continue
			}

			visited := visited[newY][newX]
			if !visited && input[newY][newX] == region.val {
				queue = append(queue, point{y: curr.y + y, x: curr.x + x})
			}

			if input[newY][newX] != region.val {
				point := point{x: newX, y: newY}
				f := fence{}
				if newX < curr.x {
					f = fence{point: point, face: WEST}
				}
				if newX > curr.x {
					f = fence{point: point, face: EAST}
				}
				if newY < curr.y {
					f = fence{point: point, face: NORTH}
				}
				if newY > curr.y {
					f = fence{point: point, face: SOUTH}
				}
				region.fences = append(region.fences, f)
				region.edges++
			}
		}
	}
	return region, visited
}

/*
all fences are marked with face NORTH, WEST, EAST, or WEST
we calculate number of sides per face

calculating faces by sorting fences by y and then x coordinates for NORTH and SOUTH fences, wise versa for EAST and WEST fences
iterating over fences slice and checking wether the previous fence
is neighbourgh or not. If it's not - increase sides count
*/
func getAllSides(r region) int {
	sides := 0
	sides += getSidesForFace(r, NORTH)
	sides += getSidesForFace(r, SOUTH)
	sides += getSidesForFace(r, EAST)
	sides += getSidesForFace(r, WEST)
	return sides
}

func getSidesForFace(r region, f face) int {
	sides := 0
	fences := []fence{}
	for i := range r.fences {
		if r.fences[i].face == f {
			fences = append(fences, r.fences[i])
		}
	}
	if f == NORTH || f == SOUTH {
		slices.SortFunc(fences, func(a, b fence) int {
			return cmp.Or(
				cmp.Compare(a.point.y, b.point.y),
				cmp.Compare(a.point.x, b.point.x),
			)
		})
	} else {
		slices.SortFunc(fences, func(a, b fence) int {
			return cmp.Or(
				cmp.Compare(a.point.x, b.point.x),
				cmp.Compare(a.point.y, b.point.y),
			)
		})
	}
	for i := range fences {
		curr := fences[i].point
		if sides == 0 {
			sides++
			continue
		}

		if f == EAST || f == WEST {
			if curr.x == fences[i-1].point.x &&
				curr.y-fences[i-1].point.y == 1 {
				continue
			}
			sides++
		} else {
			if curr.y == fences[i-1].point.y &&
				curr.x-fences[i-1].point.x == 1 {
				continue
			}
			sides++
		}
	}
	return sides
}