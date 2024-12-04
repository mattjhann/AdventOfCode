package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("prod.txt")
	check(err)

	rows := strings.Split(string(data), "\r\n")

	var ws [][]rune
	for _, row := range rows {
		ws = append(ws, []rune(row))
	}

	results := doWordSearch(ws)

	fmt.Print(len(results))
}

func doWordSearch(ws [][]rune) [][]int {
	var results [][]int

	// for each row
	for irow := 1; irow < len(ws)-1; irow++ {
		// for each col in ws
		for icol := 1; icol < len(ws[irow])-1; icol++ {
			// if the letter is 'A'
			if ws[irow][icol] == rune('A') {
				// gather all diagonals
				diags := []rune{
					ws[irow+1][icol+1],
					ws[irow+1][icol-1],
					ws[irow-1][icol-1],
					ws[irow-1][icol+1],
				}

				if validateDiags(diags) {
					results = append(results, []int{irow, icol})
				}
			}
		}
	}

	return results
}

func validateDiags(diags []rune) bool {
	targets := [][]rune{
		[]rune{'M', 'M', 'S', 'S'},
		[]rune{'S', 'M', 'M', 'S'},
		[]rune{'S', 'S', 'M', 'M'},
		[]rune{'M', 'S', 'S', 'M'},
	}

	for _, target := range targets {
		if reflect.DeepEqual(target, diags) {
			return true
		}
	}

	return false
}
