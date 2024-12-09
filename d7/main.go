package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type Equation struct {
	target         int
	equation       []int
	operators      [][]string
	validOperators []string
}

func (eq *Equation) findOperators(operators []string) {
	if len(operators) == len(eq.equation)-1 {
		total := eq.equation[0]
		for i := 0; i < len(eq.equation)-1; i++ {
			if operators[i] == "*" {
				total *= eq.equation[i+1]
			} else if operators[i] == "+" {
				total += eq.equation[i+1]
			} else if operators[i] == "||" {
				total, _ = strconv.Atoi(fmt.Sprint(strconv.Itoa(total), strconv.Itoa(eq.equation[i+1])))
			} else {
				panic("operator not programmed")
			}
		}
		if total == eq.target {
			operatorsCopy := make([]string, len(operators))
			copy(operatorsCopy, operators)
			eq.operators = append(eq.operators, operatorsCopy)
		}
		return
	}

	for _, operator := range eq.validOperators {
		operators = append(operators, operator)
		eq.findOperators(operators)
		operators = operators[:len(operators)-1]
	}
}

func parseInput(input []byte) []Equation {
	result := []Equation{}
	in := string(input)
	lines := strings.Split(in, "\r\n")
	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		target, err := strconv.Atoi(splitLine[0])
		check(err)
		operandsStr := strings.Split(splitLine[1], " ")
		var operandsInt []int
		for _, i := range operandsStr {
			this, err := strconv.Atoi(i)
			check(err)
			operandsInt = append(operandsInt, this)
		}
		result = append(result, Equation{target: target, equation: operandsInt, validOperators: []string{"*", "+", "||"}})
	}
	return result
}

func main() {
	txtInput, err := os.ReadFile("prod_input.txt")
	check(err)

	equations := parseInput(txtInput)

	total := 0
	for i, _ := range equations {
		equations[i].findOperators([]string{})
		if len(equations[i].operators) > 0 {
			total += equations[i].target
		}
	}

	fmt.Println(total)
}
