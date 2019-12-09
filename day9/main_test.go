package main

import (
	"sync"
	"testing"

	"github.com/advent_of_code_2019/intcode"
	"github.com/advent_of_code_2019/utils"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	instructions := utils.LoadInstructions("example1")

	var wg sync.WaitGroup
	wg.Add(1)
	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 1)
	go intcode.RunAmplifier("", wg, instructions, inputChannel, outputChannel)

	output := intcode.WaitAndAccumulateOutput(outputChannel)
	assert.Equal(t, output, instructions, "FUCK")
}

func TestExample2(t *testing.T) {
	instructions := utils.LoadInstructions("example2")

	var wg sync.WaitGroup
	wg.Add(1)
	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 1)
	go intcode.RunAmplifier("", wg, instructions, inputChannel, outputChannel)

	output := intcode.WaitAndAccumulateOutput(outputChannel)
	assert.Equal(t, 1, len(output), "FUCK")
	assert.Equal(t, 16, len(utils.IntToSlice(output[0])), "FUCK")
}

func TestExample3(t *testing.T) {
	instructions := utils.LoadInstructions("example3")

	var wg sync.WaitGroup
	wg.Add(1)
	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 1)
	go intcode.RunAmplifier("", wg, instructions, inputChannel, outputChannel)

	output := intcode.WaitAndAccumulateOutput(outputChannel)
	assert.Equal(t, 1, len(output), "FUCK")
	assert.Equal(t, instructions[1], output[0], "FUCK")
}

func TestPart1(t *testing.T) {
	instructions := utils.LoadInstructions("input")

	var wg sync.WaitGroup
	wg.Add(1)
	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 1)
	go intcode.RunAmplifier("", wg, instructions, inputChannel, outputChannel)

	// Start test mode
	inputChannel <- 1

	output := intcode.WaitAndAccumulateOutput(outputChannel)
	assert.Equal(t, 3780860499, output[0], "FUCK")
}
