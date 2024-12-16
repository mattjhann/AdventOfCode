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
	gridSize := Vector{101, 103}
	locations := make(map[string]int)

	for j := 0; j < 1000; j++ {
		for i, _ := range robots {
			robots[i].moveBySeconds(j, gridSize)
			locations[fmt.Sprint(robots[i].currentPos.x, robots[i].currentPos.y)] += 1
		}
		printRobots(locations, gridSize)
		println("Seconds: ", j)
		for k := range locations {
			delete(locations, k)
		}
	}

	return getSafetyScore(locations, gridSize)
}

func getSafetyScore(locations map[string]int, gridSize Vector) int {
	var quadrants []int = make([]int, 4)

	for y := 0; y < gridSize.y; y++ {
		for x := 0; x < gridSize.x; x++ {
			if x < gridSize.x/2 {
				if y < gridSize.y/2 {
					quadrants[0] += locations[fmt.Sprint(x, y)]
					fmt.Print(locations[fmt.Sprint(x, y)])
				} else if y >= gridSize.y-gridSize.y/2 {
					quadrants[2] += locations[fmt.Sprint(x, y)]
					fmt.Print(locations[fmt.Sprint(x, y)])
				} else {
					fmt.Print(" ")
				}
			} else if x >= gridSize.x-gridSize.x/2 {
				if y < gridSize.y/2 {
					quadrants[1] += locations[fmt.Sprint(x, y)]
					fmt.Print(locations[fmt.Sprint(x, y)])
				} else if y >= gridSize.y-gridSize.y/2 {
					quadrants[3] += locations[fmt.Sprint(x, y)]
					fmt.Print(locations[fmt.Sprint(x, y)])
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("--------------------")

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func printRobots(locations map[string]int, gridSize Vector) {
	for y := 0; y < gridSize.y; y++ {
		for x := 0; x < gridSize.x; x++ {
			if v := locations[fmt.Sprint(x, y)]; v > 0 {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("--------------------")
}
