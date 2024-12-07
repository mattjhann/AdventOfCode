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

func findStart(room [][]rune) []int {
	for x, row := range room {
		for y, _ := range row {
			if room[x][y] != '.' && room[x][y] != rune('#') {
				room[x][y] = '.'
				return []int{x, y}
			}
		}
	}
	return nil
}

func strToRune(strs []string) [][]rune {
	out := [][]rune{}
	for _, str := range strs {
		out = append(out, []rune(str))
	}
	return out
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

var route map[string]bool
var routeFull map[string]bool
var obstructions [][]int

func patrol(room [][]rune, loc, dir []int) {
	j := 0
	for {
		// check next step not out of bounds
		if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
			return
		}

		// turn right at obstruction
		for room[loc[0]+dir[0]][loc[1]+dir[1]] == rune('#') {
			dir = turnRight(dir)
			if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
				return
			}
		}

		// take a step
		loc[0] = loc[0] + dir[0]
		loc[1] = loc[1] + dir[1]

		// add the new location to the list
		route[fmt.Sprint(loc)] = true
		routeFull[fmt.Sprint(loc, dir)] = true

		// create a copy of the room with an obstruction
		if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
			return
		}
		playDir := make([]int, len(dir))
		playLoc := make([]int, len(loc))
		copy(playDir, dir)
		copy(playLoc, loc)

		store := room[loc[0]+dir[0]][loc[1]+dir[1]]
		room[loc[0]+dir[0]][loc[1]+dir[1]] = rune('#')

		// check for loops
		if checkLoop(room, playLoc, playDir) {
			obstructions = append(obstructions, loc)
			fmt.Println(loc)
		}
		room[loc[0]+dir[0]][loc[1]+dir[1]] = store
		j++
	}
}

func checkLoop(room [][]rune, loc, dir []int) bool {
	potentialRoute := make(map[string]bool)

	j := 0
	for {
		// check next step not out of bounds
		if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
			return false
		}

		// turn right at obstruction
		for room[loc[0]+dir[0]][loc[1]+dir[1]] == rune('#') {
			dir = turnRight(dir)
			if !checkBounds(room, loc[0]+dir[0], loc[1]+dir[1]) {
				return false
			}
		}

		// take a step
		loc[0] = loc[0] + dir[0]
		loc[1] = loc[1] + dir[1]

		// if location/direction already in potentialRoute return
		if potentialRoute[fmt.Sprint(loc, dir)] {
			return true
		}
		if route[fmt.Sprint(loc, dir)] {
			return true
		}

		// add the new location to the list
		potentialRoute[fmt.Sprint(loc, dir)] = true

		//error handling
		if j > len(room)*len(room[0]) {
			panic("what the hell")
		}
		if len(potentialRoute) > len(room)*len(room[0]) {
			panic("what the hell")
		}
		j++
	}
}

func main() {
	route = make(map[string]bool)
	routeFull = make(map[string]bool)

	input, err := os.ReadFile("prod_input.txt")
	check(err)

	sroom := strings.Split(string(input), "\r\n")
	room := strToRune(sroom)

	loc := findStart(room)

	patrol(room, loc, []int{-1, 0})

	for i, x := range room {
		for j, _ := range x {
			if route[fmt.Sprint("[", i, " ", j, "]")] {
				fmt.Print("X")
			} else {
				fmt.Print(string(room[i][j]))
			}
		}
		fmt.Print("\n")
	}

	fmt.Println(len(route))
	fmt.Println(len(obstructions))
}
