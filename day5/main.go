package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func loadNumbers() []int {
	var numbers []int
	bytes, _ := ioutil.ReadFile("./day5/input")
	for _, str := range strings.Split(string(bytes), ",") {
		integer, _ := strconv.Atoi(str)
		numbers = append(numbers, integer)
	}
	return numbers
}

func intToSlice(input int) []int {
	s := strconv.Itoa(input)
	slice := []int{}

	for i := 0; i < len(s); i++ {
		value, _ := strconv.Atoi(string(s[i]))
		slice = append(slice, value)
	}
	return slice
}

func addPading(slice []int, digits int) []int {
	finalSlice := make([]int, digits)
	offset := digits - len(slice)

	for i := 0; i < len(slice); i++ {
		finalSlice[offset+i] = slice[i]
	}
	return finalSlice
}

func ProcessInstruction(instruction int) (opCode int, parametersMode []int) {
	base := addPading(intToSlice(instruction), 5)
	parametersMode = base[0:3]
	opCode = base[3]*10 + base[4]
	return
}

func getValue(instructions []int, parameterMode int, value int) int {
	switch parameterMode {
	case 0:
		return instructions[value]
	case 1:
		return value
	default:
		log.Panic("Something wrong is happening")
	}

	return -1
}

func runPart1(instructions []int) {
	input := 1
	var value1 int
	var value2 int
	var value int
	i := 0

	for {
		// fmt.Printf("%d\n", instructions[i])
		opCode, parameterMode := ProcessInstruction(instructions[i])
		// fmt.Printf("\n%+v\t", instructions[i])
		// fmt.Printf("%d\t %+v", opCode, parameterMode)

		switch opCode {
		case 1:
			value1 = getValue(instructions, parameterMode[2], instructions[i+1])
			value2 = getValue(instructions, parameterMode[1], instructions[i+2])
			instructions[instructions[i+3]] = value1 + value2
			i = i + 4
		case 2:
			value1 = getValue(instructions, parameterMode[2], instructions[i+1])
			value2 = getValue(instructions, parameterMode[1], instructions[i+2])
			instructions[instructions[i+3]] = value1 * value2
			i = i + 4
		case 3:
			instructions[instructions[i+1]] = input
			i = i + 2
		case 4:
			value = getValue(instructions, parameterMode[2], instructions[i+1])
			if value != 0 {
				fmt.Printf("Part1 Solution:   %d   ", value)
			}
			i = i + 2
		case 99:
			print("\nChao\n")
			return
		}
	}

}

func runPart2(instructions []int) {
	input := 5
	var value1 int
	var value2 int
	var value int
	i := 0

	for {
		// fmt.Printf("%d\n", instructions[i])
		opCode, parameterMode := ProcessInstruction(instructions[i])
		// fmt.Printf("\n%+v\t", instructions[i])
		// fmt.Printf("%d\t %+v", opCode, parameterMode)

		switch opCode {
		case 1:
			value1 = getValue(instructions, parameterMode[2], instructions[i+1])
			value2 = getValue(instructions, parameterMode[1], instructions[i+2])
			instructions[instructions[i+3]] = value1 + value2
			i = i + 4
		case 2:
			value1 = getValue(instructions, parameterMode[2], instructions[i+1])
			value2 = getValue(instructions, parameterMode[1], instructions[i+2])
			instructions[instructions[i+3]] = value1 * value2
			i = i + 4
		case 3:
			instructions[instructions[i+1]] = input
			i = i + 2
		case 4:
			value = getValue(instructions, parameterMode[2], instructions[i+1])
			if value != 0 {
				fmt.Printf("Part2 Solution:   %d   ", value)
			}
			i = i + 2
		case 5: // Jump-if-true
			value1 = getValue(instructions, parameterMode[2], instructions[i+1])
			value2 = getValue(instructions, parameterMode[1], instructions[i+2])
			if value1 != 0 {
				i = value2
			} else {
				i = i + 3
			}
		case 6: // Jump-if-false
			value1 = getValue(instructions, parameterMode[2], instructions[i+1])
			value2 = getValue(instructions, parameterMode[1], instructions[i+2])
			if value1 == 0 {
				i = value2
			} else {
				i = i + 3
			}
		case 7: // Less than
			value1 = getValue(instructions, parameterMode[2], instructions[i+1])
			value2 = getValue(instructions, parameterMode[1], instructions[i+2])
			if value1 < value2 {
				instructions[instructions[i+3]] = 1
			} else {
				instructions[instructions[i+3]] = 0
			}
			i = i + 4
		case 8: // Less than
			value1 = getValue(instructions, parameterMode[2], instructions[i+1])
			value2 = getValue(instructions, parameterMode[1], instructions[i+2])
			if value1 == value2 {
				instructions[instructions[i+3]] = 1
			} else {
				instructions[instructions[i+3]] = 0
			}
			i = i + 4
		case 99:
			print("\nChao\n")
			return
		}
	}

}

func main() {
	instructions := loadNumbers()

	part1Instructions := make([]int, len(instructions))
	copy(part1Instructions, instructions)
	runPart1(part1Instructions)

	runPart2(instructions)
}
