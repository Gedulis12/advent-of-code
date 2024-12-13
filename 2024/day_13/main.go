package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type vec2 struct {
	m, n float64
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

	a := vec2{}
	b := vec2{}
	prize := vec2{}

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Trim(text, " ")
		if strings.HasPrefix(text, "Button A:") {
			split := strings.Split(text, ",")
			split1 := strings.Split(split[0], "+")
			split2 := strings.Split(split[1], "+")
			x, _ := strconv.Atoi(split1[len(split1)-1])
			y, _ := strconv.Atoi(split2[len(split2)-1])
			a.m = float64(x)
			a.n = float64(y)
			continue
		}
		if strings.HasPrefix(text, "Button B:") {
			split := strings.Split(text, ",")
			split1 := strings.Split(split[0], "+")
			split2 := strings.Split(split[1], "+")
			x, _ := strconv.Atoi(split1[len(split1)-1])
			y, _ := strconv.Atoi(split2[len(split2)-1])
			b.m = float64(x)
			b.n = float64(y)
			continue
		}
		if strings.HasPrefix(text, "Prize:") {
			split := strings.Split(text, ",")
			split1 := strings.Split(split[0], "=")
			split2 := strings.Split(split[1], "=")
			x, _ := strconv.Atoi(split1[len(split1)-1])
			y, _ := strconv.Atoi(split2[len(split2)-1])
			prize = vec2{m: float64(x), n: float64(y)}
		}
		if text == "" {
			a = vec2{}
			b = vec2{}
			prize = vec2{}
			continue
		}
		tokensX, tokensY := gaussianElim(a, b, prize)
		ans += tokensX * 3
		ans += tokensY * 1
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

	a := vec2{}
	b := vec2{}
	prize := vec2{}

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Trim(text, " ")
		if strings.HasPrefix(text, "Button A:") {
			split := strings.Split(text, ",")
			split1 := strings.Split(split[0], "+")
			split2 := strings.Split(split[1], "+")
			x, _ := strconv.Atoi(split1[len(split1)-1])
			y, _ := strconv.Atoi(split2[len(split2)-1])
			a.m = float64(x)
			a.n = float64(y)
			continue
		}
		if strings.HasPrefix(text, "Button B:") {
			split := strings.Split(text, ",")
			split1 := strings.Split(split[0], "+")
			split2 := strings.Split(split[1], "+")
			x, _ := strconv.Atoi(split1[len(split1)-1])
			y, _ := strconv.Atoi(split2[len(split2)-1])
			b.m = float64(x)
			b.n = float64(y)
			continue
		}
		if strings.HasPrefix(text, "Prize:") {
			split := strings.Split(text, ",")
			split1 := strings.Split(split[0], "=")
			split2 := strings.Split(split[1], "=")
			x, _ := strconv.Atoi(split1[len(split1)-1])
			y, _ := strconv.Atoi(split2[len(split2)-1])
			x += 10000000000000
			y += 10000000000000
			prize = vec2{m: float64(x), n: float64(y)}
		}
		if text == "" {
			a = vec2{}
			b = vec2{}
			prize = vec2{}
			continue
		}
		tokensX, tokensY := gaussianElim(a, b, prize)
		ans += tokensX * 3
		ans += tokensY * 1
	}

	return ans
}

func gaussianElim(a, b, prize vec2) (int, int) {
	matrix := [2][2 + 1]float64{{a.m, b.m, prize.m}, {a.n, b.n, prize.n}}

	// swap rows until one with largest first element is on top
	maxRow := 0
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] > matrix[maxRow][0] {
			matrix = swapRows(matrix, i, maxRow)
			maxRow = i
		}
	}
	if matrix[0][0] == 0 {
		matrix = swapRows(matrix, 0, 1)
	}

	div := matrix[0][0]
	for i := 0; i < len(matrix[0]); i++ {
		matrix[0][i] = matrix[0][i] / div
	}

	multiple := matrix[1][0]
	for i := 0; i < len(matrix[0]); i++ {
		matrix[1][i] = (matrix[1][i] - (matrix[0][i] * multiple))
	}

	div = matrix[1][1]
	for i := 0; i < len(matrix[0]); i++ {
		matrix[1][i] = matrix[1][i] / div
	}

	multiple = matrix[0][1]
	for i := 0; i < len(matrix[0]); i++ {
		matrix[0][i] = (matrix[0][i] - (matrix[1][i] * multiple))
	}

	ax := int(math.Round(matrix[0][2]))
	bx := int(math.Round(matrix[1][2]))

	valid := check(ax, bx, a, b, prize)

	if valid {
		return ax, bx
	}
	return 0, 0
}

func check(ax, bx int, a, b, prize vec2) bool {
	checkX := 0
	checkY := 0

	checkX += ax * int(a.m)
	checkX += bx * int(b.m)

	checkY += ax * int(a.n)
	checkY += bx * int(b.n)
	if checkX == int(prize.m) && checkY == int(prize.n) {
		return true
	}
	return false
}

func swapRows(matrix [2][3]float64, a, b int) [2][3]float64 {
	temp := matrix[a]
	matrix[a] = matrix[b]
	matrix[b] = temp
	return matrix
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
