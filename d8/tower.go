package main

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
		// if they're the same frequency and not the same tower
		if !tower.eq(towers[i]) && tower.frequency == towers[i].frequency {
			vector := tower.location.subtract(towers[i].location)
			tower.antinodes = append(tower.antinodes, tower.location)
			// loop through all possible antinodes in that direction
			for antinode := tower.location.add(vector); antinode.inBounds(grid); antinode = antinode.add(vector) {
				tower.antinodes = append(tower.antinodes, antinode)
			}
		}
	}
}
