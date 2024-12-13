package part2

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Target struct {
	A     Vector
	B     Vector
	Prize Vector
}

func ParseText(s string) []Target {
	lines := strings.Split(s, "\r\n")
	targets := []Target{}
	target := Target{}
	for _, line := range lines {
		if v := strings.Split(line, ":"); v[0] == "Button A" {
			matcher, _ := regexp.Compile("(?:X)([+-]\\d+)")
			target.A.x, _ = strconv.Atoi(matcher.FindStringSubmatch(v[1])[1])
			matcher, _ = regexp.Compile("(?:Y)([+-]\\d+)")
			target.A.y, _ = strconv.Atoi(matcher.FindStringSubmatch(v[1])[1])
		} else if v[0] == "Button B" {
			matcher, _ := regexp.Compile("(?:X)([+-]\\d+)")
			target.B.x, _ = strconv.Atoi(matcher.FindStringSubmatch(v[1])[1])
			matcher, _ = regexp.Compile("(?:Y)([+-]\\d+)")
			target.B.y, _ = strconv.Atoi(matcher.FindStringSubmatch(v[1])[1])
		} else if v[0] == "Prize" {
			matcher, _ := regexp.Compile("(?:X=)(\\d+)")
			target.Prize.x, _ = strconv.Atoi(matcher.FindStringSubmatch(v[1])[1])
			matcher, _ = regexp.Compile("(?:Y=)(\\d+)")
			target.Prize.y, _ = strconv.Atoi(matcher.FindStringSubmatch(v[1])[1])
			targets = append(targets, target)
			target = Target{}
		}
	}

	return targets
}

func DoPuzzle(file string) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	targets := ParseText(string(text))

	cost := 0
	for _, target := range targets {
		minimum := -1
		// loop 1-100 times
		for presses := 1; presses <= 200; presses++ {
			// loop for each combination
			for pressA := 0; pressA <= presses; pressA++ {
				pressB := (presses - pressA)
				finish := Vector{0, 0}
				finish = finish.add(Vector{pressA * target.A.x, pressA * target.A.y})
				finish = finish.add(Vector{pressB * target.B.x, pressB * target.B.y})
				if target.Prize.eq(finish) {
					if v := pressA*3 + pressB; v < minimum || minimum == -1 {
						minimum = v
					}
				}
			}
		}
		if minimum != -1 {
			cost += minimum
		}
	}

	return cost
}
