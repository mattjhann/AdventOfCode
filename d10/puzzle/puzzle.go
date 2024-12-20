package puzzle

import (
	"os"
	"strconv"
	"strings"
)

func TextToGrid(text string) ([][]int, []Vector) {
	rows := strings.Split(text, "\r\n")
	grid := [][]int{}
	trailHeads := []Vector{}
	for y, row := range rows {
		irow := []int{}
		for x, char := range row {
			digit, _ := strconv.Atoi(string(char))
			irow = append(irow, digit)
			if digit == 0 {
				trailHeads = append(trailHeads, Vector{x: x, y: y})
			}
		}
		grid = append(grid, irow)
	}
	return grid, trailHeads
}

var directions []Vector = []Vector{
	Vector{x: 1, y: 0},
	Vector{x: 0, y: 1},
	Vector{x: -1, y: 0},
	Vector{x: 0, y: -1},
}

func findTrail(grid [][]int, start, loc Vector, num int, dict map[Vector]int) {
	if num == 9 {
		dict[start] += 1
		return
	}
	for _, dir := range directions {
		newLoc := loc.add(dir)

		if checkBounds(grid, newLoc.x, newLoc.y) && grid[newLoc.y][newLoc.x] == num+1 {
			num++
			findTrail(grid, start, newLoc, num, dict)
			num--
		}
	}
}

func checkBounds(grid [][]int, x, y int) bool {
	if x < 0 || x >= len(grid) {
		return false
	}
	if y < 0 || y >= len(grid[x]) {
		return false
	}
	return true
}

func DoPuzzle(file string) int {
	text, _ := os.ReadFile(file)
	grid, trailHeads := TextToGrid(string(text))

	score := 0
	for _, start := range trailHeads {
		trailTails := make(map[Vector]int)
		findTrail(grid, start, start, 0, trailTails)
		score += trailTails[start]
	}
	return score
}
