package part1

import (
	"os"
	"strings"
)

func ParseText(text string) (Vector, map[Vector]rune, []Vector) {
	instructionParse := false
	movements := []Vector{}
	grid := make(map[Vector]rune)
	x := 0
	location := Vector{-1, -1}
	for _, v := range strings.Split(text, "\r\n") {
		if v == "" {
			instructionParse = true
		}
		if instructionParse {
			for _, instruction := range v {
				switch instruction {
				case '<':
					movements = append(movements, Vector{0, -1})
				case 'v':
					movements = append(movements, Vector{1, 0})
				case '>':
					movements = append(movements, Vector{0, 1})
				case '^':
					movements = append(movements, Vector{-1, 0})
				}
			}
		} else {
			for y, _ := range v {
				switch v[y] {
				case '@':
					location = Vector{x, y}
				case '#':
					grid[Vector{x, y}] = '#'
				case 'O':
					grid[Vector{x, y}] = 'O'
				}
			}
			x++
		}
	}
	return location, grid, movements
}

func DoPuzzle(file string) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	location, grid, movements := ParseText(string(text))

	return RunSim(location, grid, movements)
}

func RunSim(location Vector, grid map[Vector]rune, movements []Vector) int {
	for _, move := range movements {
		next := location.add(move)
		for grid[next] != '#' {
			if v, ok := grid[next]; v == 'O' {
				next = next.add(move)
			} else if ok == false {
				location = location.add(move)
				if grid[location] == 'O' {
					delete(grid, location)
					grid[next] = 'O'
				}
				break
			}
		}
		// printGrid(grid, location)
	}

	return getScore(grid)
}

func getScore(grid map[Vector]rune) int {
	count := 0
	for k, v := range grid {
		if v == 'O' {
			count += 100*k.x + k.y
		}
	}
	return count
}

// func printGrid(grid map[Vector]rune, location Vector) {
// 	for x := 0; x < 10; x++ {
// 		for y := 0; y < 10; y++ {
// 			if v, ok := grid[Vector{x, y}]; ok {
// 				fmt.Print(string(v))
// 			} else {
// 				if location.x == x && location.y == y {
// 					fmt.Print("@")
// 				} else {
// 					fmt.Print(".")
// 				}
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// }
