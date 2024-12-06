package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findStart(room []string) ([]int, rune) {
	for x, row := range room {
		for y, col := range row {
			if room[x][y] != '.' && room[x][y] != '#' {
				return []int{x, y}, col
			}
		}
	}
	return nil, '0'
}

func checkBounds(room []string, x, y int) bool {
	if x < 0 || x >= len(room) {
		return false
	}
	if y < 0 || y >= len(room[x]) {
		return false
	}
	return true
}

// walk returns the location of the next obstacle, or (-1,-1) if the guard leaves the room
func walk(room []string, dir []int, loc []int, locations map[string]bool) []int {
	// if next location is not an obstacle
	for room[loc[0]+dir[0]][loc[1]+dir[1]] != '#' {
		locations[fmt.Sprint(loc)] = true
		loc[0] = loc[0] + dir[0]
		loc[1] = loc[1] + dir[1]

		// check if the guard will leave on the next step
		if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
			locations[fmt.Sprint(loc)] = true
			return []int{-1, -1}
		}
	}

	// otherwise return the guard's location
	locations[fmt.Sprint(loc)] = true
	return loc
}

func turnRight(dir []int) []int {
	dirs := [][]int{
		[]int{-1, 0}, // up
		[]int{0, 1},  // right
		[]int{1, 0},  // down
		[]int{0, -1}, // left
	}

	for i, d := range dirs {
		if reflect.DeepEqual(d, dir) {
			return dirs[(i+1)%len(dirs)]
		}
	}

	return []int{}
}

func main() {
	locations := make(map[string]bool)
	vectors := map[rune][]int{
		'^': []int{-1, 0},
		'>': []int{0, 1},
		'v': []int{1, 0},
		'<': []int{0, -1},
	}

	input, err := os.ReadFile("prod_input.txt")
	check(err)

	room := strings.Split(string(input), "\r\n")

	loc, guard := findStart(room)
	dir := vectors[guard]
	fmt.Println("Starting at pos: ", loc, "\nTravelling in dir: ", vectors[guard])

	for loc[0] != -1 && loc[1] != -1 {
		loc = walk(room, dir, loc, locations)
		dir = turnRight(dir)
	}

	fmt.Println("Squares walked: ", len(locations))
}
