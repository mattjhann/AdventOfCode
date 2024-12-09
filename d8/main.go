package main

import (
	"fmt"
	"os"
	"strings"
)

func parseText(text string) ([]Tower, []string) {
	towers := []Tower{}
	rows := strings.Split(text, "\r\n")
	for y, row := range rows {
		for x, char := range row {
			if char != '.' {
				towers = append(towers, Tower{location: Vector{x: x, y: y}, frequency: char})
			}
		}
	}
	return towers, rows
}

func main() {
	text, _ := os.ReadFile("prod_input.txt")

	towers, grid := parseText(string(text))

	antinodes := make(map[Vector]bool)

	// loop through towers and find their antinodes and add to map (to remove duplicates)
	for i, _ := range towers {
		towers[i].findAntinodes(towers, grid)
		for _, antinode := range towers[i].antinodes {
			antinodes[antinode] = true
		}
	}

	fmt.Print(len(antinodes))
}
