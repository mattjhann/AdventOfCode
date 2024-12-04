package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("prod.txt")
	check(err)
	fmt.Println(string(dat))

	matcher, err := regexp.Compile("(do\\(\\)|don't\\(\\)|mul\\(\\d+,\\d+\\))")
	check(err)

	matches := matcher.FindAll(dat, -1)

	var ops [][]int
	matching := true
	for _, match := range matches {
		s := string(match)
		matcher, err = regexp.Compile("\\d+")
		check(err)

		if s == "do()" {
			matching = true
			continue
		} else if s == "don't()" {
			matching = false
			continue
		}

		if !matching {
			continue
		}

		digits := matcher.FindAll([]byte(s), 2)
		var mul []int
		for _, digit := range digits {
			d, err := strconv.Atoi(string(digit))
			check(err)
			mul = append(mul, d)
		}
		ops = append(ops, mul)
	}

	count := 0
	for _, op := range ops {
		count += op[0] * op[1]
	}

	fmt.Println(count)
}
