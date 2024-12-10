package main

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
