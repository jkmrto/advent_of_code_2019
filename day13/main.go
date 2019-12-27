package main

import (
	"fmt"
	"sync"

	"github.com/advent_of_code_2019/intcode"
	"github.com/advent_of_code_2019/utils"
)

func Part1(instructions []int) int {
	var wg sync.WaitGroup
	wg.Add(1)

	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 1)

	go intcode.RunAmplifier("", wg, instructions, inputChannel, outputChannel)

	var x, y, tileId int
	var isActive bool

	counter := 0
	for {
		x, isActive = <-outputChannel

		if !isActive {
			break
		}

		y = <-outputChannel
		tileId = <-outputChannel

		if tileId == 2 {
			counter++
		}

		fmt.Sprintf("(%d, %d, %d)", x, y, tileId)
	}

	return counter
}

func main() {
	instructions := utils.LoadInstructions("./day13/input")

	fmt.Printf("Part 1 solution: %d", Part1(instructions))

	// fmt.Printf("%+v", instructions)

}
