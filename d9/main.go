package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkSumDisk(disk []int) int {
	total := 0
	for i, id := range disk {
		if id >= 0 {
			total += i * id
		}
	}
	return total
}

func printDisk(disk []int) string {
	var result strings.Builder
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			result.WriteString(".")
		} else {
			result.WriteString(strconv.Itoa(disk[i]))
		}
	}
	return result.String()
}

func defragDisk(disk []int) []int {
	for i := slices.Index(disk, -1); i != -1; i = slices.Index(disk, -1) {
		j := getLastNum(disk)
		if i >= j {
			break
		}
		if j == -1 {
			break
		}
		disk = swapInt(disk, i, j)
	}
	return disk
}

func swapInt(disk []int, i1, i2 int) []int {
	store := disk[i1]
	disk[i1] = disk[i2]
	disk[i2] = store
	return disk
}

func getLastNum(disk []int) int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			return i
		}
	}
	return -1
}

func main() {
	file, _ := os.ReadFile("prod_input.txt")

	disk := []int{}
	id := 0
	for char, _ := range file {
		if char%2 == 0 {
			// file block loop number of times
			num, _ := strconv.Atoi(string(file[char]))
			for i := 0; i < num; i++ {
				disk = append(disk, id)
			}
			id++
		} else {
			num, _ := strconv.Atoi(string(file[char]))
			for i := 0; i < num; i++ {
				disk = append(disk, -1)
			}
		}
	}

	disk = defragDisk(disk)

	checkSum := checkSumDisk(disk)

	fmt.Println(checkSum)
}
