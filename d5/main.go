package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getData(input string) (map[int][]int, [][]int) {
	data, err := os.ReadFile(input)
	check(err)

	var strRules []string
	var strPages []string
	rows := strings.Split(string(data), "\r\n")
	for i := range rows {
		if rows[i] == "" {
			strRules = rows[0:i]
			strPages = rows[i+1:]
			break
		}
	}

	rules, err := splitToMap(strRules)
	check(err)
	pages, err := splitToIntSlice(strPages)
	check(err)

	return rules, pages
}

func splitToMap(lines []string) (map[int][]int, error) {
	dict := make(map[int][]int)
	for _, str := range lines {
		parts := strings.Split(str, "|")
		var pair []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			pair = append(pair, num)
		}
		dict[pair[0]] = append(dict[pair[0]], pair[1])
	}

	return dict, nil
}

func splitToIntSlice(lines []string) ([][]int, error) {
	var result [][]int
	for _, str := range lines {
		parts := strings.Split(str, ",")
		var pair []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			pair = append(pair, num)
		}
		result = append(result, pair)
	}
	return result, nil
}

func validatePage(rules map[int][]int, page []int) bool {
	for pageIndex, pageNum := range page {
		for _, ruleNum := range rules[pageNum] {
			indexOfSecondNum := slices.Index(page, ruleNum)
			if pageIndex > indexOfSecondNum && indexOfSecondNum != -1 {
				return false
			}

		}
	}
	return true
}

func fixOrder(rules map[int][]int, page []int) []int {
	// while invalid
	for !validatePage(rules, page) {
		page = bubblePass(rules, page)
	}

	return page
}

func bubblePass(rules map[int][]int, page []int) []int {
	slices.Reverse(page)
	// for each number in the page
	for i := 0; i < len(page)-1; i++ {
		// for each rule for that number
		for _, rule := range rules[page[i]] {
			// if that rule matches
			incorrect := slices.Index(page, rule)
			if incorrect > i {
				page = slices.Insert(page, i, page[incorrect])
				page = removeElement(page, incorrect+1)
			}
		}
	}
	slices.Reverse(page)
	return page
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	rules, pages := getData("prod_input.txt")
	part1count := 0
	part2count := 0
	for _, page := range pages {
		if !validatePage(rules, page) {
			fixOrder(rules, page)
			part2count += page[(len(page)-1)/2]
		} else {
			part1count += page[(len(page)-1)/2]
		}
	}
	fmt.Println("Part 1: ", part1count)
	fmt.Println("Part 2: ", part2count)
}
