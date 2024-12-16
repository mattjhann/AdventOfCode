package part1

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Reindeer struct {
	location Vector
}

func ParseText(text string) (Reindeer, map[Vector]rune, Vector) {
	grid := make(map[Vector]rune)
	x := 0
	finish := Vector{-1, -1}
	reindeer := Reindeer{Vector{-1, -1}}
	for _, v := range strings.Split(text, "\r\n") {
		for y, _ := range v {
			switch v[y] {
			case '#':
				grid[Vector{x, y}] = '#'
			case 'S':
				reindeer.location = Vector{x, y}
			case 'E':
				finish = Vector{x, y}
			}
		}
		x++
	}
	return reindeer, grid, finish
}

func DoPuzzle(file string) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	reindeer, maze, finish := ParseText(string(text))

	return DoMaze(reindeer, maze, finish)
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

func (s Step) turnRight(grid map[Vector]rune) (Step, bool) {
	// increment direction
	s.direction = s.direction + 1
	if s.direction > 3 {
		s.direction = 0
	}
	// take a step
	s.location = s.location.add(directions[s.direction])
	if _, ok := grid[s.location]; ok { // fail if it's a wall
		return s, false
	}
	s.score += 1001
	return s, true
}

func (s Step) turnLeft(grid map[Vector]rune) (Step, bool) {
	s.direction = s.direction - 1
	if s.direction < 0 {
		s.direction = 3
	}
	s.location = s.location.add(directions[s.direction])
	if _, ok := grid[s.location]; ok {
		return s, false
	}
	s.score += 1001
	return s, true
}

func (s Step) stepForward(grid map[Vector]rune) (Step, bool) {
	s.location = s.location.add(directions[s.direction])
	s.score += 1
	if _, ok := grid[s.location]; ok {
		return s, false
	}
	return s, true
}

func DoMaze(reindeer Reindeer, grid map[Vector]rune, finish Vector) int {
	history := make(map[string]int)
	queue := []Step{}
	initial := Step{location: reindeer.location, score: 0, direction: 0}
	queue = append(queue, initial)

	for len(queue) > 0 {
		// pop from queue
		next := queue[0]
		queue = queue[1:]

		// check for finish condition
		if next.location.eq(finish) {
			return next.score
		}

		if v, _ := history[fmt.Sprint(next.location, next.direction)]; v == 0 || v > next.score { // if not visited, or the new score is better
			history[fmt.Sprint(next.location, next.direction)] = next.score // add to history

			// check steps and add to queue
			if v, ok := next.stepForward(grid); ok {
				queue = append(queue, v)
			}
			if v, ok := next.turnLeft(grid); ok {
				queue = append(queue, v)
			}
			if v, ok := next.turnRight(grid); ok {
				queue = append(queue, v)
			}
		}

		// sort the queue, so we're always evaluating the lowest score first
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].score < queue[j].score
		})
	}

	return -1
}

// PRINTGRID: CAN PLACE ANYWHERE IN THE LOOP TO SEE THE GRID PRINTED TO THE CONSOLE
func printGrid(grid map[Vector]rune, history map[Vector]int, location, finish Vector) {
	for x := 0; x < 141; x++ {
		for y := 0; y < 141; y++ {
			if location.x == x && location.y == y {
				fmt.Print("@")
			} else if finish.x == x && finish.y == y {
				fmt.Print("E")
			} else if v, ok := grid[Vector{x, y}]; ok {
				fmt.Print(string(v))
			} else if _, ok := history[Vector{x, y}]; ok {
				fmt.Print("^")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
