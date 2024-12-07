package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func drawRoom(room [][]rune, loc []int) {
	var str string
	for x := range 10 {
		for y := range 10 {
			if x == 5 && y == 5 {
				str += "O"
			} else {
				if !checkBounds(room, loc[0]-5+x, loc[1]-5+y) {
					str += "."
				} else {
					str += string(room[loc[0]-5+x][loc[1]-5+y])
				}
			}
		}
		str += "\n"
	}
	fmt.Print(str)
	time.Sleep(50 * time.Millisecond)
}

func strToRune(strs []string) [][]rune {
	out := [][]rune{}
	for _, str := range strs {
		out = append(out, []rune(str))
	}
	return out
}

func findStart(room [][]rune) ([]int, rune) {
	for x, row := range room {
		for y, col := range row {
			if room[x][y] != '.' && room[x][y] != rune('#') {
				return []int{x, y}, col
			}
		}
	}
	return nil, '0'
}

func checkBounds(room [][]rune, x, y int) bool {
	if x < 0 || x >= len(room) {
		return false
	}
	if y < 0 || y >= len(room[x]) {
		return false
	}
	return true
}

// walk returns the location of the next obstacle, or (-1,-1) if the guard leaves the room
func walk(room [][]rune, dir []int, loc []int, locations map[string]bool, obstructions map[string]bool) []int {
	// if next location is not an obstacle
	if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
		return []int{-1, -1}
	}
	for room[loc[0]+dir[0]][loc[1]+dir[1]] != rune('#') {
		locations[fmt.Sprint(loc, dir)] = true
		loc[0] = loc[0] + dir[0]
		loc[1] = loc[1] + dir[1]

		// check if the guard will leave on the next step
		if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
			locations[fmt.Sprint(loc, dir)] = true
			return []int{-1, -1}
		}

		nextLoc := []int{loc[0] + dir[0], loc[1] + dir[1]}
		if !locations[fmt.Sprint(nextLoc, dir)] { // next loc is not in locations
			store := room[nextLoc[0]][nextLoc[1]]
			room[nextLoc[0]][nextLoc[1]] = rune('#')
			possibleLocs := make(map[string]bool)

			ploc := make([]int, len(loc))
			pdir := make([]int, len(dir))
			copy(ploc, loc)
			copy(pdir, dir)

			if checkLoops(room, pdir, ploc, locations, possibleLocs) {
				obstructions[fmt.Sprint(nextLoc)] = true
			}
			room[nextLoc[0]][nextLoc[1]] = store
		}

	}

	// otherwise return the guard's location
	locations[fmt.Sprint(loc, dir)] = true
	return loc
}

func checkLoops(room [][]rune, dir, loc []int, locations map[string]bool, possibleLocs map[string]bool) bool {
	dir = turnRight(dir)
	if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
		return false
	}

	for room[loc[0]+dir[0]][loc[1]+dir[1]] != rune('#') {
		//drawRoom(room, loc)

		possibleLocs[fmt.Sprint(loc, dir)] = true

		loc[0] = loc[0] + dir[0]
		loc[1] = loc[1] + dir[1]

		if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
			return false
		}
		if locations[fmt.Sprint(loc, dir)] {
			return true
		}
		if possibleLocs[fmt.Sprint(loc, dir)] {
			return true
		}
	}
	possibleLocs[fmt.Sprint(loc, dir)] = true
	return checkLoops(room, dir, loc, locations, possibleLocs)
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
	obstructions := make(map[string]bool)
	vectors := map[rune][]int{
		'^': []int{-1, 0},
		'>': []int{0, 1},
		'v': []int{1, 0},
		'<': []int{0, -1},
	}

	input, err := os.ReadFile("prod_input.txt")
	check(err)

	sroom := strings.Split(string(input), "\r\n")
	room := strToRune(sroom)

	loc, guard := findStart(room)
	dir := vectors[guard]
	fmt.Println("Starting at pos: ", loc, "\nTravelling in dir: ", vectors[guard])

	for loc[0] != -1 && loc[1] != -1 {
		loc = walk(room, dir, loc, locations, obstructions)
		dir = turnRight(dir)
	}

	fmt.Println("Squares walked: ", len(locations))
	fmt.Println("Obstructions for loop: ", len(obstructions))
}
