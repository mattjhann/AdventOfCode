package part3

import (
	"fmt"
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

func applyRules(n int, value int) int {

	if n == 0 {
		memo[fmt.Sprint(n, value)] = 1
		return 1
	}
	if v, ok := memo[fmt.Sprint(n, value)]; ok {
		return v
	}
	if value == 0 {
		result := applyRules(n-1, 1)
		memo[fmt.Sprint(n, value)] = result
		return result
	}
	str := strconv.Itoa(value)
	if len(str)%2 == 0 {
		i1, _ := strconv.Atoi(str[:len(str)/2])
		i2, _ := strconv.Atoi(str[len(str)/2:])
		result := applyRules(n-1, i1) + applyRules(n-1, i2)
		memo[fmt.Sprint(n, value)] = result
		return result
	}
	result := applyRules(n-1, value*2024)
	memo[fmt.Sprint(n, value)] = result
	return result
}

func DoPuzzle(file string, iterations int) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	list := parseText(string(text))

	result := 0
	for i := 0; i < len(list); i++ {
		result += applyRules(iterations, list[i])
	}

	return result
}
