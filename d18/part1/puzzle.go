package part1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseText(text string) (Vector, []Vector, Vector) {
	grid := []Vector{}
	finish := Vector{70, 70}
	start := Vector{0, 0}
	for _, v := range strings.Split(text, "\r\n") {
		s := strings.Split(v, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		grid = append(grid, Vector{x, y})
	}
	return start, grid, finish
}

func DoPuzzle(file string) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	start, grid, finish := ParseText(string(text))

	res, ok := shortestRoute(start, finish, grid, 1024)
	if !ok {
		panic("failure")
	}

	fmt.Println(res)

	res, ok = InterestingshortestRoute(start, finish, grid)
	if !ok {
		panic("failure")
	}
	fmt.Println(res)

	return res
}

type Step struct {
	location  Vector
	score     int
	direction int // a pointer for the list of directions below
}

var directions []Vector = []Vector{
	Vector{0, 1},
	Vector{1, 0},
	Vector{0, -1},
	Vector{-1, 0},
}

func contains(s []Vector, e Vector) bool {
	for _, a := range s {
		if a.eq(e) {
			return true
		}
	}
	return false
}

func (s Step) makeStep(grid []Vector, dir int, size Vector) (Step, bool) {
	s.location = s.location.add(directions[dir])
	s.score += 1
	if contains(grid, s.location) || !s.location.inBounds(size) {
		return s, false
	}
	return s, true
}

func shortestRoute(start, finish Vector, grid []Vector, nanoSeconds int) (int, bool) {
	queue := []Step{}
	queue = append(queue, Step{start, 0, 0})
	visited := make(map[Vector]bool)
	newGrid := grid[:nanoSeconds]
	for len(queue) > 0 {
		// pop queue
		next := queue[0]
		queue = queue[1:]
		if next.location.eq(finish) {
			return next.score, true
		}
		if !visited[next.location] {
			for i := 0; i < len(directions); i++ {
				if v, ok := next.makeStep(newGrid, i, finish); ok {
					queue = append(queue, v)
					visited[next.location] = true
				}
			}
		}
	}

	return -1, false
}

func InterestingshortestRoute(start, finish Vector, grid []Vector) (int, bool) {
	queue := []Step{}
	queue = append(queue, Step{start, 0, 0})
	visited := make(map[Vector]bool)
	for len(queue) > 0 {
		// pop queue
		next := queue[0]
		queue = queue[1:]
		newGrid := grid[:next.score]
		if next.location.eq(finish) {
			return next.score, true
		}
		if !visited[next.location] {
			for i := 0; i < len(directions); i++ {
				if v, ok := next.makeStep(newGrid, i, finish); ok {
					queue = append(queue, v)
					visited[next.location] = true
				}
			}
		}
	}

	return -1, false
}

func PrintGrid(newGrid []Vector, visited map[Vector]bool, finish Vector) {
	grid := make([][]rune, finish.x+1)
	for i, _ := range grid {
		grid[i] = make([]rune, finish.y+1)
	}
	for v := range visited {
		grid[v.y][v.x] = '.'
	}
	for _, v := range newGrid {
		grid[v.y][v.x] = '#'
	}
	for y := 0; y < len(grid); y++ {
		fmt.Println(string(grid[y]))
	}
	fmt.Println("--------------")
}
