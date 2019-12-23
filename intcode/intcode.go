package intcode

import (
	"fmt"
	"log"
	"sync"

	"github.com/advent_of_code_2019/utils"
)

func getValue(instructions []int, parameterMode int, value int, relativeBase int) int {
	switch parameterMode {
	case 0:
		return instructions[value]
	case 1:
		return value
	case 2:
		return instructions[relativeBase+value]
	default:
		log.Panic("Something wrong is happening")
	}
	return -1
}

func getPosition(parameterMode int, value int, relativeBase int) int {
	switch parameterMode {
	case 0:
		return value
	case 2:
		return relativeBase + value
	default:
		log.Panic("Something wrong is happening")
	}
	return -1
}

func ProcessInstruction(instruction int) (opCode int, parametersMode []int) {
	base := utils.AddPading(utils.IntToSlice(instruction), 5)
	parametersMode = base[0:3]
	opCode = base[3]*10 + base[4]
	return
}

func RunAmplifier(label string, wg sync.WaitGroup, inputInstructions []int, input <-chan int, output chan<- int) {
	var value1 int
	var value2 int
	var value int
	i := 0

	relativeBase := 0
	memorySpace := 1000000
	instructions := make([]int, memorySpace)
	copy(instructions, inputInstructions)

	for {
		opCode, parameterMode := ProcessInstruction(instructions[i])
		fmt.Printf("\n%+v\t", instructions[i])
		fmt.Printf("%d\t %+v", opCode, parameterMode)

		switch opCode {
		case 1: // Sum
			value1 = getValue(instructions, parameterMode[2], instructions[i+1], relativeBase)
			value2 = getValue(instructions, parameterMode[1], instructions[i+2], relativeBase)
			destPosition := getPosition(parameterMode[0], instructions[i+3], relativeBase)
			// fmt.Printf(" Value1: %d, value2: %d, destPosition: %d", value1, value2, destPosition)
			instructions[destPosition] = value1 + value2
			i = i + 4
		case 2: // Multiply
			value1 = getValue(instructions, parameterMode[2], instructions[i+1], relativeBase)
			value2 = getValue(instructions, parameterMode[1], instructions[i+2], relativeBase)
			destPosition := getPosition(parameterMode[0], instructions[i+3], relativeBase)
			fmt.Printf(" Value1: %d, value2: %d, destPosition: %d", value1, value2, destPosition)
			instructions[destPosition] = value1 * value2
			i = i + 4
		case 3: // Input
			fmt.Printf("Waiting for intput:")
			phase := <-input
			destPosition := getPosition(parameterMode[2], instructions[i+1], relativeBase)
			fmt.Printf(" input: %d, destPosition: %d", phase, destPosition)
			instructions[destPosition] = phase
			i = i + 2
		case 4: // Output
			value = getValue(instructions, parameterMode[2], instructions[i+1], relativeBase)
			fmt.Printf("%s is waiting to sent his output: %d\n", label, value)
			output <- value
			i = i + 2
		case 5: // Jump-if-true
			value1 = getValue(instructions, parameterMode[2], instructions[i+1], relativeBase)
			value2 = getValue(instructions, parameterMode[1], instructions[i+2], relativeBase)
			if value1 != 0 {
				i = value2
			} else {
				i = i + 3
			}
		case 6: // Jump-if-false
			value1 = getValue(instructions, parameterMode[2], instructions[i+1], relativeBase)
			value2 = getValue(instructions, parameterMode[1], instructions[i+2], relativeBase)
			if value1 == 0 {
				i = value2
			} else {
				i = i + 3
			}
		case 7: // Less than
			value1 = getValue(instructions, parameterMode[2], instructions[i+1], relativeBase)
			value2 = getValue(instructions, parameterMode[1], instructions[i+2], relativeBase)
			destPosition := getPosition(parameterMode[0], instructions[i+3], relativeBase)

			if value1 < value2 {
				instructions[destPosition] = 1
			} else {
				instructions[destPosition] = 0
			}
			i = i + 4
		case 8: // Less than
			value1 = getValue(instructions, parameterMode[2], instructions[i+1], relativeBase)
			value2 = getValue(instructions, parameterMode[1], instructions[i+2], relativeBase)
			destPosition := getPosition(parameterMode[0], instructions[i+3], relativeBase)

			if value1 == value2 {
				instructions[destPosition] = 1
			} else {
				instructions[destPosition] = 0
			}
			i = i + 4
		case 9: // Less than
			value1 = getValue(instructions, parameterMode[2], instructions[i+1], relativeBase)
			relativeBase = relativeBase + value1
			i = i + 2
		case 99:
			// print("\nChao\n")
			close(output)
			wg.Done()
			return
		}
	}
}

func WaitAndAccumulateOutput(outputChannel <-chan int) []int {
	output := []int{}
	for {
		valueReceived, isActive := <-outputChannel
		if isActive {
			output = append(output, valueReceived)
		} else {
			break
		}
	}

	return output
}
