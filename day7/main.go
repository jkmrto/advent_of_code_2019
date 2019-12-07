package main

import (
	"fmt"
	"github.com/gitchander/permutation"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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

func RunAmplifier(label string, originalInstructions []int, input <-chan int, output chan<- int) {
	var value1 int
	var value2 int
	var value int
	i := 0

	instructions := make([]int, len(originalInstructions))
	copy(instructions, originalInstructions)

	for {
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
			fmt.Printf("%s is waiting for input\n", label)
			phase := <-input
			fmt.Printf("%s input has arrived: %d\n", label, phase)
			instructions[instructions[i+1]] = phase
			i = i + 2
		case 4:
			value = getValue(instructions, parameterMode[2], instructions[i+1])
			fmt.Printf("%s is waiting to sent his output\n", label)
			output <- value
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
			// print("\nChao\n")
			return
		}
	}
}

func LoadInstructions(path string) []int {
	var numbers []int
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		log.Panic(err)
	}
	for _, str := range strings.Split(string(bytes), ",") {
		integer, _ := strconv.Atoi(str)
		numbers = append(numbers, integer)
	}
	return numbers
}

func ExecuteAmplificationPipeline(originalInstructions []int, phases []int) int {

	chanToA := make(chan int, 1)
	chanAtoB := make(chan int, 1)
	chanBtoC := make(chan int, 1)
	chanCtoD := make(chan int, 1)
	chanDtoE := make(chan int, 1)
	chanEtoOutput := make(chan int, 1)

	go RunAmplifier("A", originalInstructions, chanToA, chanAtoB)
	go RunAmplifier("B", originalInstructions, chanAtoB, chanBtoC)
	go RunAmplifier("C", originalInstructions, chanBtoC, chanCtoD)
	go RunAmplifier("D", originalInstructions, chanCtoD, chanDtoE)
	go RunAmplifier("E", originalInstructions, chanDtoE, chanEtoOutput)

	// Startup with phases

	chanToA <- phases[0]
	chanAtoB <- phases[1]
	chanBtoC <- phases[2]
	chanCtoD <- phases[3]
	chanDtoE <- phases[4]
	chanToA <- 0

	return <-chanEtoOutput
}

func max(values []int) int {
	max := 0

	for _, value := range values {
		if value > max {
			max = value
		}
	}

	return max
}

func main() {
	instructions := LoadInstructions("./day7/input")
	phases := []int{0, 1, 2, 3, 4}
	phaseSequences := permutation.New(permutation.IntSlice(phases))
	outputs := []int{}
	for phaseSequences.Next() {
		fmt.Printf("\n")
		output := ExecuteAmplificationPipeline(instructions, phases)
		outputs = append(outputs, output)
	}
	fmt.Printf("Ejercicio 7: %v\n", max(outputs))
}
