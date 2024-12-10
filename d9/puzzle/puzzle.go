package puzzle

import (
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
	for num := getLastNot(disk); num >= 0; num-- {
		blockStart, blockEnd := getBlock(disk, num)
		freeSpaceStart := getFreeSpaceOfLen(disk, blockEnd-blockStart)
		if freeSpaceStart != -1 && freeSpaceStart < blockStart {
			j := freeSpaceStart
			for i := blockStart; i < blockEnd; i++ {
				swapInt(disk, i, j)
				j++
			}
		}
	}
	return disk
}

func getFreeSpaceOfLen(disk []int, length int) int {
	for i, _ := range disk {
		if disk[i] == -1 {
			index := i
			counter := 0
			for i < len(disk) && disk[i] == -1 {
				counter++
				i++
				if counter == length {
					return index
				}
			}
		}
	}
	return -1
}

func getBlock(disk []int, num int) (int, int) {
	start := slices.Index(disk, num)
	slices.Reverse(disk)
	end := len(disk) - slices.Index(disk, num)
	slices.Reverse(disk)
	return start, end
}

func swapInt(disk []int, i1, i2 int) []int {
	store := disk[i1]
	disk[i1] = disk[i2]
	disk[i2] = store
	return disk
}

func getLastNot(disk []int) int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			return disk[i]
		}
	}
	return -1
}

func DoPuzzle(file string) int {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	disk := []int{}
	id := 0
	for char, _ := range text {
		if char%2 == 0 {
			// file block loop number of times
			num, _ := strconv.Atoi(string(text[char]))
			for i := 0; i < num; i++ {
				disk = append(disk, id)
			}
			id++
		} else {
			num, _ := strconv.Atoi(string(text[char]))
			for i := 0; i < num; i++ {
				disk = append(disk, -1)
			}
		}
	}

	disk = defragDisk(disk)

	checkSum := checkSumDisk(disk)

	return checkSum
}
