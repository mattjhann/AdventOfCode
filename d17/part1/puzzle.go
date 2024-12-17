package part1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseText(text string) ([]uint64, []uint64) {
	parseInst := false
	registers := []uint64{}
	instructions := []uint64{}
	for _, v := range strings.Split(text, "\r\n") {
		if !parseInst {
			if v == "" {
				parseInst = true
			} else {
				s, _ := strconv.Atoi(strings.Split(v, ": ")[1])
				registers = append(registers, uint64(s))
			}
		} else {
			for _, y := range strings.Split(strings.Split(v, ": ")[1], ",") {
				s, _ := strconv.Atoi(y)
				instructions = append(instructions, uint64(s))
			}
		}
	}
	return registers, instructions
}

func DoPuzzle(file string) {
	text, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	registers, instructions := ParseText(string(text))

	fmt.Print(registers, "\n", instructions, "\n\n")

	output := RunComputer(registers, instructions)

	for _, i := range output {
		fmt.Print(i, ",")
	}

	return
}

func RunComputer(registers []uint64, instructions []uint64) []uint64 {
	output := []uint64{}
	for i := 1; i < len(instructions); i += 2 {
		var operand uint64 = 0
		switch instructions[i] {
		case 0:
			operand = 0
		case 1:
			operand = 1
		case 2:
			operand = 2
		case 3:
			operand = 3
		case 4:
			operand = registers[0]
		case 5:
			operand = registers[1]
		case 6:
			operand = registers[2]
		case 7:
			panic("instruction 7 not implemented")
		}
		switch instructions[i-1] {
		case uint64(0):
			registers[0] = registers[0] / (uint64(1) << operand)
		case uint64(1):
			registers[1] = registers[1] ^ instructions[i]
		case uint64(2):
			registers[1] = operand % 8
		case uint64(3):
			if registers[0] != 0 {
				i = int(instructions[i]) - 1
			}
		case uint64(4):
			registers[1] = registers[1] ^ registers[2]
		case uint64(5):
			output = append(output, operand%8)
		case uint64(6):
			registers[1] = registers[0] / (uint64(1) << operand)
		case uint64(7):
			registers[2] = registers[0] / (uint64(1) << operand)
		}
	}
	return output
}
