package main

import (
	"fmt"
	"sync"

	"github.com/advent_of_code_2019/intcode"
	"github.com/advent_of_code_2019/utils"
)

func main() {
	instructions := utils.LoadInstructions("./day9/input")

	var wg sync.WaitGroup
	wg.Add(1)

	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 1)
	go intcode.RunAmplifier("", wg, instructions, inputChannel, outputChannel)

	// Start test mode
	inputChannel <- 1

	output := intcode.WaitAndAccumulateOutput(outputChannel)
	fmt.Printf("\n%+v", output)
}
