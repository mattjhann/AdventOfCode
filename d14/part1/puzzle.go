package part1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseText(text string) []Robot {
	robots := []Robot{}
	for _, v := range strings.Split(text, "\r\n") {
		r, _ := regexp.Compile("(?:[=,])([-\\d]+)")
		line := r.FindAllStringSubmatch(v, 4)

		posx, err := strconv.Atoi(line[0][1])
		posy, err := strconv.Atoi(line[1][1])
		movx, err := strconv.Atoi(line[2][1])
		movy, err := strconv.Atoi(line[3][1])

		if err != nil {
			panic(err.Error())
		}

		robots = append(robots, Robot{
			startPos: Vector{posx, posy},
			velocity: Vector{movx, movy},
		})
	}

	return robots
}

type Robot struct {
	startPos   Vector
	velocity   Vector
	currentPos Vector
}

func (r *Robot) moveBySeconds(seconds int, grid Vector) {
	r.currentPos.x = (r.startPos.x + seconds*r.velocity.x) % grid.x
	r.currentPos.y = (r.startPos.y + seconds*r.velocity.y) % grid.y
	if r.currentPos.x < 0 {
		r.currentPos.x = grid.x + r.currentPos.x
	}
	if r.currentPos.y < 0 {
		r.currentPos.y = grid.y + r.currentPos.y
	}
}

func DoPuzzle(file string) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	robots := ParseText(string(text))
	gridSize := Vector{11, 7}
	locations := make(map[string]int)

	for i, _ := range robots {
		for j := 0; j < 6; j++ {
			robots[i].moveBySeconds(j, gridSize)
			fmt.Println("Seconds: ", j, "\t Current pos: ", robots[i].currentPos)
		}
		robots[i].moveBySeconds(5, gridSize)
		locations[fmt.Sprint(robots[i].currentPos.x, robots[i].currentPos.y)] += 1
	}

	return len(locations)
}
