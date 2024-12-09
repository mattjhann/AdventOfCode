package main

import (
	"fmt"
	"os"
	"strings"
)

type Tower struct {
	location  Vector
	frequency rune
	antinodes []Vector
}

func (a *Tower) eq(b Tower) bool {
	if a.location.x == b.location.x && a.location.y == b.location.y && a.frequency == b.frequency {
		return true
	} else {
		return false
	}
}

func (tower *Tower) findAntinodes(towers []Tower, grid []string) {
	// for all the other towers
	for i, _ := range towers {
		// if they're the same frequency
		if !tower.eq(towers[i]) && tower.frequency == towers[i].frequency {
			vector := tower.location.subtract(towers[i].location)
			tower.antinodes = append(tower.antinodes, tower.location)
			for antinode := tower.location.add(vector); antinode.inBounds(grid); antinode = antinode.add(vector) {
				tower.antinodes = append(tower.antinodes, antinode)
			}
		}
	}
}

type Vector struct {
	x, y int
}

func (vec *Vector) inBounds(grid []string) bool {
	return vec.x >= 0 && vec.x < len(grid[0]) && vec.y >= 0 && vec.y < len(grid)
}

func (a *Vector) subtract(b Vector) Vector {
	return Vector{
		x: a.x - b.x,
		y: a.y - b.y,
	}
}

func (a *Vector) add(b Vector) Vector {
	return Vector{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

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

	// loop through towers and find their antinode
	for i, _ := range towers {
		towers[i].findAntinodes(towers, grid)
		for _, antinode := range towers[i].antinodes {
			antinodes[antinode] = true
		}
	}

	fmt.Print(len(antinodes))
}
