package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var directions = [][]int{
	{0, 1},   //r
	{1, 1},   //rd
	{1, 0},   //d
	{1, -1},  //dl
	{0, -1},  //l
	{-1, -1}, //ul
	{-1, 0},  //u
	{-1, 1},  //ur
}

var results []Result

type Result struct {
	start  []int
	finish []int
}

func main() {
	data, err := os.ReadFile("prod.txt")
	check(err)

	rows := strings.Split(string(data), "\r\n")

	var ws [][]rune
	for _, row := range rows {
		ws = append(ws, []rune(row))
	}

	var targetWords [][]rune
	targetWords = append(targetWords, []rune("XMAS"))

	doWordSearch(ws, targetWords)

	fmt.Print(len(results))
}

func doWordSearch(ws [][]rune, targetWords [][]rune) []Result {

	// for each target word
	for _, targetWord := range targetWords {
		// for each row in ws
		for irow, crow := range ws {
			// for each col in ws
			for icol, _ := range crow {
				// for each direction
				for _, dir := range directions {
					checkWord(ws, targetWord, []int{irow, icol}, 0, dir)
				}
			}
		}
	}

	return results
}

func checkWord(ws [][]rune, targetWord []rune, loc []int, cindex int, dir []int) {

	// new location is going to be the start location + the direction vector * the letter index
	newLoc := []int{loc[0] + (dir[0] * cindex), loc[1] + (dir[1] * cindex)}

	// if new location out of bounds, return
	if !checkBounds(ws, newLoc[0], newLoc[1]) {
		return
	}

	// if this new location matches the target word, then check the next letter
	if ws[newLoc[0]][newLoc[1]] == targetWord[cindex] {
		// if it's the last character, add the result to the results list and return
		if len(targetWord)-1 == cindex {
			results = append(results, Result{
				start:  []int{loc[0], loc[1]},
				finish: []int{newLoc[0], newLoc[1]},
			})
			return
		}

		cindex++
		checkWord(ws, targetWord, loc, cindex, dir)
	}
}

func checkBounds(ws [][]rune, x, y int) bool {
	if x < 0 || x >= len(ws) {
		return false
	}
	if y < 0 || y >= len(ws[x]) {
		return false
	}
	return true
}
