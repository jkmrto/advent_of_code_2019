package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/advent_of_code_2019/intcode"
	"github.com/advent_of_code_2019/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	instructions := utils.LoadInstructions("./input")

	var wg sync.WaitGroup
	wg.Add(1)

	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 1)
	go intcode.RunAmplifier("", wg, instructions, inputChannel, outputChannel)

	currentPos := pos{0, 0}
	facingPos := pos{0, 1}

	panel := map[pos]int{}

	for {
		inputColor := lookPositionColorInPanel(panel, currentPos)

		fmt.Printf("\nColor: %+v", inputColor)
		inputChannel <- inputColor
		color, _ := <-outputChannel
		panel[currentPos] = color
		facingChange, isActive := <-outputChannel

		if isActive {
			currentPos, facingPos = updatePostion(currentPos, facingPos, facingChange)
		} else {
			break
		}

	}

	assert.Equal(t, 1894, len(keys(panel)), "FUCK")

}
