package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/advent_of_code_2019/utils"
	"github.com/gitchander/permutation"
)

func ProcessInstruction(instruction int) (opCode int, parametersMode []int) {
	base := utils.AddPading(utils.IntToSlice(instruction), 5)
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

func RunAmplifier(label string, wg sync.WaitGroup, originalInstructions []int, input <-chan int, output chan<- int) {
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
			// fmt.Printf("%s is waiting for input\n", label)
			phase := <-input
			// fmt.Printf("%s input has arrived: %d\n", label, phase)
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
			close(output)
			wg.Done()
			return
		}
	}
}

func ExecuteAmplificationPipeline(originalInstructions []int, phases []int) int {

	chanToA := make(chan int, 1)
	chanAtoB := make(chan int, 1)
	chanBtoC := make(chan int, 1)
	chanCtoD := make(chan int, 1)
	chanDtoE := make(chan int, 1)
	chanEtoOutput := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(5)
	go RunAmplifier("A", wg, originalInstructions, chanToA, chanAtoB)
	go RunAmplifier("B", wg, originalInstructions, chanAtoB, chanBtoC)
	go RunAmplifier("C", wg, originalInstructions, chanBtoC, chanCtoD)
	go RunAmplifier("D", wg, originalInstructions, chanCtoD, chanDtoE)
	go RunAmplifier("E", wg, originalInstructions, chanDtoE, chanEtoOutput)

	// Startup with phases
	chanToA <- phases[0]
	chanAtoB <- phases[1]
	chanBtoC <- phases[2]
	chanCtoD <- phases[3]
	chanDtoE <- phases[4]
	chanToA <- 0

	return <-chanEtoOutput
}

func ExecuteAmplificationWithFeedbackLoop(originalInstructions []int, phases []int) int {

	chanToA := make(chan int, 1)
	chanAtoB := make(chan int, 1)
	chanBtoC := make(chan int, 1)
	chanCtoD := make(chan int, 1)
	chanDtoE := make(chan int, 1)
	chanEtoOutput := make(chan int, 1)

	var feedbackValue int
	var wg sync.WaitGroup

	wg.Add(5)
	go RunAmplifier("A", wg, originalInstructions, chanToA, chanAtoB)
	go RunAmplifier("B", wg, originalInstructions, chanAtoB, chanBtoC)
	go RunAmplifier("C", wg, originalInstructions, chanBtoC, chanCtoD)
	go RunAmplifier("D", wg, originalInstructions, chanCtoD, chanDtoE)
	go RunAmplifier("E", wg, originalInstructions, chanDtoE, chanEtoOutput)

	// Startup with phases
	chanToA <- phases[0]
	chanAtoB <- phases[1]
	chanBtoC <- phases[2]
	chanCtoD <- phases[3]
	chanDtoE <- phases[4]
	chanToA <- 0

	for {
		valueReceived, isActive := <-chanEtoOutput
		fmt.Printf("Feedback value: %d, isActive: %+v", feedbackValue, isActive)

		if isActive == true {
			feedbackValue = valueReceived
			chanToA <- feedbackValue
		} else {
			fmt.Printf("Solution: %d", feedbackValue)
			return feedbackValue
		}
	}

	wg.Wait()
	return 0
}

func part1(instructions []int) {
	phases := []int{0, 1, 2, 3, 4}
	phaseSequences := permutation.New(permutation.IntSlice(phases))
	outputs := []int{}
	for phaseSequences.Next() {
		fmt.Printf("\n")
		output := ExecuteAmplificationPipeline(instructions, phases)
		outputs = append(outputs, output)
	}
	fmt.Printf("Ejercicio 7 Parte1: %v\n", utils.Max(outputs))
}

func part2(instructions []int) {
	phases := []int{5, 6, 7, 8, 9}
	phaseSequences := permutation.New(permutation.IntSlice(phases))
	outputs := []int{}
	for phaseSequences.Next() {
		fmt.Printf("\n")
		output := ExecuteAmplificationWithFeedbackLoop(instructions, phases)
		outputs = append(outputs, output)
	}
	fmt.Printf("\nEjercicio 7 Parte2: %v\n", utils.Max(outputs))
}

func main() {
	instructions := utils.LoadInstructions("./day7/input")
	part1(instructions)
	part2(instructions)
}
