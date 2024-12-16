package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type DIRECTION int

const (
	UP DIRECTION = iota
	DOWN
	LEFT
	RIGHT
)

type edge struct {
	dst   *node
	score int
}

type point struct {
	x, y int
}

type node struct {
	id  int
	loc point
	dir DIRECTION
	adj []edge
}

func main() {
	start1 := time.Now().UnixMicro()
	fmt.Println(SolvePart1("example1"))
	end1 := time.Now().UnixMicro()
	fmt.Println("part 1 took: ", end1-start1)

}

func SolvePart1(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ans := 0

	maze := [][]byte{}
	for scanner.Scan() {
		text := scanner.Text()
		row := []byte{}
		for i := range text {
			row = append(row, text[i])
		}
		maze = append(maze, row)
	}
	graph := graphCreate(maze)
	start := findStartPoint(maze)
	end := findEndPoint(maze)
	src := findVertInGraph(graph, start, RIGHT)
	dst := []*node{
		findVertInGraph(graph, end, LEFT),
		findVertInGraph(graph, end, RIGHT),
		findVertInGraph(graph, end, UP),
		findVertInGraph(graph, end, DOWN),
	}

	ans = dijkstra(graph, src, dst)

	return ans
}

func dijkstra(graph []*node, src *node, dst []*node) int {
	dist := make(map[int]int)
	visited := make(map[int]bool)
	prev := make(map[int]int)

	for i, _ := range graph {
		dist[graph[i].id] = MaxInt
		visited[graph[i].id] = false
	}
	dist[src.id] = 0

	for i := 0; i < len(graph); i++ {
		fmt.Println("checking: ", i, "/", len(graph))
		u := minDist(dist, visited)
		if u == -1 {
			break
		}
		current := getVertById(graph, u)
		visited[u] = true

		for _, edge := range current.adj {
			if !visited[edge.dst.id] {
				newDist := dist[u] + edge.score
				if newDist < dist[edge.dst.id] {
					dist[edge.dst.id] = newDist
					prev[edge.dst.id] = u
				}
			}
		}
	}

	// get minimum distance
	minimum := dist[dst[0].id]
	for i := range dst {
		if dist[dst[i].id] < minimum {
			minimum = dist[dst[i].id]
		}
	}

	path := []point{}
	for i := range dst {
		if dist[dst[i].id] == minimum {
			currentId := dst[i].id
			for currentId != src.id {
				current := getVertById(graph, currentId)
				path = append([]point{current.loc}, path...)
				currentId = prev[currentId]
			}
			path = append([]point{src.loc}, path...)
		}
	}

	return minimum
}
func minDist(dist map[int]int, visited map[int]bool) int {
	mindst := MaxInt
	nodeId := -1

	for id, d := range dist {
		if !visited[id] && d < mindst {
			mindst = d
			nodeId = id
		}
	}
	return nodeId
}

func getVertById(graph []*node, id int) *node {
	for i := range graph {
		if graph[i].id == id {
			return graph[i]
		}
	}
	return nil
}

func printGraph(graph []*node) {
	for i := range graph {
		fmt.Println("ID: ", graph[i].id, " X, Y : ", graph[i].loc.x, graph[i].loc.y, " DIR: ", graph[i].dir, "EDGES:")
		for j := range graph[i].adj {
			fmt.Println("\t", "ID: ", graph[i].id, " X, Y : ", graph[i].adj[j].dst.loc.x, graph[i].adj[j].dst.loc.y, " DIR: ", graph[i].adj[j].dst.dir, " SCORE: ", graph[i].adj[j].score)
		}
	}

}

func findStartPoint(maze [][]byte) point {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == 'S' {
				return point{x: j, y: i}
			}
		}
	}
	return point{x: -1, y: -1}
}
func findEndPoint(maze [][]byte) point {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == 'E' {
				return point{x: j, y: i}
			}
		}
	}
	return point{x: -1, y: -1}
}

func graphCreate(maze [][]byte) []*node {
	graph := []*node{}
	for i := len(maze) - 1; i >= 0; i-- {
		for j := range maze[i] {
			if maze[i][j] == '#' {
				continue
			}

			var l, r, u, d *node

			if !locHasVert(graph, point{x: j, y: i}) {
				l, r, u, d = addVertToGraph(&graph, point{x: j, y: i})
			} else {
				l = findVertInGraph(graph, point{x: j, y: i}, LEFT)
				r = findVertInGraph(graph, point{x: j, y: i}, RIGHT)
				u = findVertInGraph(graph, point{x: j, y: i}, UP)
				d = findVertInGraph(graph, point{x: j, y: i}, DOWN)
			}

			if maze[i][j-1] != '#' {
				newp := point{x: j - 1, y: i}
				if !locHasVert(graph, newp) {
					addVertToGraph(&graph, newp)
				}
				n := findVertInGraph(graph, newp, LEFT)
				if !l.edgeHas(n) {
					l.edgeAdd(n, 1)
				}
			}

			if maze[i][j+1] != '#' {
				newp := point{x: j + 1, y: i}
				if !locHasVert(graph, newp) {
					addVertToGraph(&graph, newp)
				}
				n := findVertInGraph(graph, newp, RIGHT)
				if !r.edgeHas(n) {
					r.edgeAdd(n, 1)
				}
			}

			if maze[i-1][j] != '#' {
				newp := point{x: j, y: i - 1}
				if !locHasVert(graph, newp) {
					addVertToGraph(&graph, newp)
				}
				n := findVertInGraph(graph, newp, UP)
				if !u.edgeHas(n) {
					u.edgeAdd(n, 1)
				}
			}

			if maze[i+1][j] != '#' {
				newp := point{x: j, y: i + 1}
				if !locHasVert(graph, newp) {
					addVertToGraph(&graph, newp)
				}
				n := findVertInGraph(graph, newp, DOWN)
				if !d.edgeHas(n) {
					d.edgeAdd(n, 1)
				}
			}
		}
	}
	return graph
}

func locHasVert(graph []*node, location point) bool {
	for i := range graph {
		if graph[i].loc == location {
			return true
		}
	}
	return false
}

func (node *node) edgeHas(dst *node) bool {
	for i := range node.adj {
		if node.adj[i].dst == dst {
			return true
		}
	}
	return false
}

func printMaze(maze [][]byte) {
	for i := range maze {
		for j := range maze[i] {
			fmt.Print(string(maze[i][j]))
		}
		fmt.Println()
	}
}

func nodeCreate(loc point, dir DIRECTION, id int) *node {
	node := node{loc: loc, dir: dir, id: id}
	return &node
}

func (n *node) edgeAdd(dst *node, score int) {
	edge := edge{dst: dst, score: score}
	n.adj = append(n.adj, edge)
}

func findVertInGraph(graph []*node, l point, d DIRECTION) *node {
	for i := range graph {
		if graph[i].loc == l && graph[i].dir == d {
			return graph[i]
		}
	}
	return nil
}

func getMaxGraphId(graph []*node) int {
	id := 0
	for i := range graph {
		if graph[i].id > id {
			id = graph[i].id
		}
	}
	return id
}

func addVertToGraph(graph *[]*node, l point) (*node, *node, *node, *node) {
	id := getMaxGraphId(*graph) + 1
	nodeLeft := nodeCreate(point{x: l.x, y: l.y}, LEFT, id)
	id++
	nodeRight := nodeCreate(point{x: l.x, y: l.y}, RIGHT, id)
	id++
	nodeUp := nodeCreate(point{x: l.x, y: l.y}, UP, id)
	id++
	nodeDown := nodeCreate(point{x: l.x, y: l.y}, DOWN, id)
	id++

	nodeLeft.edgeAdd(nodeUp, 1000)
	nodeLeft.edgeAdd(nodeDown, 1000)
	nodeRight.edgeAdd(nodeUp, 1000)
	nodeRight.edgeAdd(nodeDown, 1000)
	nodeUp.edgeAdd(nodeLeft, 1000)
	nodeUp.edgeAdd(nodeRight, 1000)
	nodeDown.edgeAdd(nodeLeft, 1000)
	nodeDown.edgeAdd(nodeRight, 1000)

	*graph = append(*graph, nodeLeft)
	*graph = append(*graph, nodeRight)
	*graph = append(*graph, nodeUp)
	*graph = append(*graph, nodeDown)

	return nodeLeft, nodeRight, nodeUp, nodeDown
}