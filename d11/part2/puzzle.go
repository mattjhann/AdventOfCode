package part2

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func parseText(text string) []int {
	list := strings.Split(text, " ")

	// create each successive element
	result := []int{}
	for i := 0; i < len(list); i++ {
		integer, _ := strconv.Atoi(list[i])
		result = append(result, integer)
	}

	return result
}

func getDivisor(num int) int {
	return int(math.Log10(float64(num))) + 1
}

var memo map[string]int = make(map[string]int)

func applyRules(n int, value []int) int {
	memo := make(map[int]int)
	// add each value into the memo
	for _, i := range value {
		memo[i] += 1
	}

	for i := 0; i < n; i++ {
		newMemo := make(map[int]int)
		// go through each value in the memo
		for val, count := range memo {
			nextLayer := []int{}
			// create the next layer
			if val == 0 {
				nextLayer = append(nextLayer, 1)
			} else if digits := getDivisor(val); digits%2 == 0 {
				divisor := int(math.Pow10(digits / 2))
				nextLayer = append(nextLayer, val/divisor)
				nextLayer = append(nextLayer, val%divisor)
			} else {
				nextLayer = append(nextLayer, val*2024)
			}
			for _, newVal := range nextLayer {
				newMemo[newVal] = newMemo[newVal] + count
			}
		}
		memo = newMemo
	}
	result := 0
	for _, count := range memo {
		result += count
	}
	return result
}

func DoPuzzle(file string, iterations int) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	list := parseText(string(text))

	result := 0
	result += applyRules(iterations, list)

	return result
}
