package main

import (
	"fmt"
	"main/part1"
	"main/part2"
)

func main() {
	fmt.Println(part1.DoPuzzle("part1/test_input.txt"))
	fmt.Println(part1.DoPuzzle("part1/prod_input.txt"))
	fmt.Println(part2.DoPuzzle("part2/test_input.txt"))
}
