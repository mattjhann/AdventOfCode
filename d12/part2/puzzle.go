package part2

import (
	"fmt"
	"os"
	"strings"
)

func ParseText(text string) [][]Square {
	lines := strings.Split(text, "\r\n")
	grid := make([][]Square, len(lines))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]Square, len(lines[0]))
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j].value = rune(lines[i][j])
			grid[i][j].region = 0
		}
	}
	return grid
}

var dirs []Vector = []Vector{
	Vector{1, 0}, // down
	Vector{0, 1}, // right
	Vector{-1, 0},// up
	Vector{0, -1},// left
}

type Square struct {
	value  rune
	region int
	perims int
}

func MakeFields(grid [][]Square) (map[int]int, map[int]int) {
	regionCounter := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col].region == 0 {
				regionCounter++
				makeField(grid, row, col, regionCounter)
				walkPerim(grid, row, col)
			}
		}
	}
	areas := make(map[int]int)
	perimiters := make(map[int]int)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			areas[grid[row][col].region] += 1
			perimiters[grid[row][col].region] += grid[row][col].perims
		}
	}
	return perimiters, areas
}

func walkPerim(grid [][]Square, startRow, startCol int) (int, int) {
	dirPointer := 0
	row := startRow
	col := startCol
	for fmt.Sprint(dirPointer, row, col) != fmt.Sprint(0, startRow, startCol) {
		dirPointer = (dirPointer+1)%4
		if next := dirs[dirPointer].add(Vector{row, col}); next.inBounds(grid) {
			if grid[next.x][next.y].region == grid[row][col]
		}
	}
}

func makeField(grid [][]Square, row, col, id int) {
	grid[row][col].region = id
	for _, dir := range dirs {
		newLoc := dir.add(Vector{row, col})
		// if out of bounds add a fence
		outBounds := newLoc.x < 0 || newLoc.x >= len(grid) || newLoc.y < 0 || newLoc.y >= len(grid[0])
		if outBounds {
			grid[row][col].perims += 1
			// if the region isn't set
		} else if grid[newLoc.x][newLoc.y].region == 0 {
			//
			if grid[newLoc.x][newLoc.y].value == grid[row][col].value {
				makeField(grid, newLoc.x, newLoc.y, id)
			}
		}
		if !outBounds && grid[newLoc.x][newLoc.y].value != grid[row][col].value {
			grid[row][col].perims += 1
		}
	}
}

func DoPuzzle(file string) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	grid := ParseText(string(text))
	perimiters, areas := MakeFields(grid)

	keys := make([]int, 0, len(perimiters))
	for k := range perimiters {
		keys = append(keys, k)
	}

	// for _, v := range grid {
	// 	for _, i := range v {
	// 		fmt.Print(string(i.value), i.perims, "\t")
	// 	}
	// 	fmt.Print("\n")
	// }

	result := 0
	// totalAreas := 0
	for _, key := range keys {
		result += perimiters[key] * areas[key]
		// totalAreas += areas[key]
		// fmt.Println(key, ": ", areas[key], " * ", perimiters[key], " = ", perimiters[key]*areas[key])
	}

	return result
}
