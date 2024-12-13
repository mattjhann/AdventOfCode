package part2

import (
	"math"
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
		target.Prize.x = target.Prize.x
		target.Prize.y = target.Prize.y

		a1 := target.A.x
		b1 := -1 * target.A.y
		c1 := 0
		a2 := target.B.x
		b2 := -1 * target.B.y
		c2 := target.Prize.y / target.B.y

		// calculate intersection of two lines
		x := (b1*c2 - b2*c1) / (a1*b2 - a2*b1)

		// calculate number of times A pressed to get to intersection
		pressA := float64(x) / float64(target.A.x)

		// calculate number of times B pressed to get from intersection to target
		pressB := (float64(target.Prize.x) - float64(x)) / float64(target.B.x)

		if pressA == math.Trunc(pressA) && pressB == math.Trunc(pressB) {
			cost += 3*int(pressA) + int(pressB)
		}
	}

	return cost
}
