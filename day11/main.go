package main

import (
	"fmt"
	"sync"

	"github.com/advent_of_code_2019/intcode"
	"github.com/advent_of_code_2019/utils"
)

// All the panels starts as black
// Input
// black-panel -> 0
// white-panel -> 1
// Outputs
// 1 -> paint (O || 1)
// 2 0 -> left, 1 -> right

type pos struct {
	x int
	y int
}

func keys(mymap map[pos]int) []pos {
	keys := []pos{}
	for k := range mymap {
		keys = append(keys, k)
	}
	return keys
}

func lookPositionColorInPanel(panel map[pos]int, position pos) int {
	value, exist := panel[position]
	if exist {
		return value
	} else {
		return 0
	}
}

func factingTurn(currentFacing pos, turnInstruction int) pos {
	switch currentFacing {
	case pos{0, 1}:
		switch turnInstruction {
		case 0:
			return pos{-1, 0}
		case 1:
			return pos{1, 0}
		}
	case pos{1, 0}:
		switch turnInstruction {
		case 0:
			return pos{0, 1}
		case 1:
			return pos{0, -1}
		}
	case pos{0, -1}:
		switch turnInstruction {
		case 0:
			return pos{1, 0}
		case 1:
			return pos{-1, 0}
		}
	case pos{-1, 0}:
		switch turnInstruction {
		case 0:
			return pos{0, -1}
		case 1:
			return pos{0, 1}
		}
	}

	return pos{-9999999, -99999}
}

func updatePostion(currentPostion pos, currentFacing pos, instruction int) (newPosition pos, newFacing pos) {
	newFacing = factingTurn(currentFacing, instruction)

	newPosition = pos{
		x: newFacing.x + currentPostion.x,
		y: newFacing.y + currentPostion.y,
	}
	return
}

func paintPanel(panel map[pos]int, instructions []int) map[pos]int {

	var wg sync.WaitGroup
	wg.Add(1)

	inputChannel := make(chan int, 1)
	outputChannel := make(chan int, 1)
	go intcode.RunAmplifier("", wg, instructions, inputChannel, outputChannel)

	currentPos := pos{0, 0}
	facingPos := pos{0, 1}

	for {
		inputColor := lookPositionColorInPanel(panel, currentPos)

		inputChannel <- inputColor

		color, _ := <-outputChannel

		panel[currentPos] = color

		facingChange, isActive := <-outputChannel

		if isActive {
			currentPos, facingPos = updatePostion(currentPos, facingPos, facingChange)
		} else {
			return panel
		}

	}

}

func getPanelBounds(panel map[pos]int) (xUpBound, xDownBound, yUpBound, yDownBound int) {
	for key := range panel {
		if key.x > xUpBound {
			xUpBound = key.x
		} else if key.x < xDownBound {
			xDownBound = key.x
		}

		if key.y > yUpBound {
			yUpBound = key.y
		} else if key.y < yDownBound {
			yDownBound = key.y
		}
	}
	return
}

func part1(instructions []int) {

	panel := map[pos]int{}

	panel = paintPanel(panel, instructions)

	fmt.Printf("\n\nPart1 solution: %+v \n", len(keys(panel)))
}

func main() {
	instructions := utils.LoadInstructions("./day11/input")

	// part1(instructions)

	panel := map[pos]int{}
	panel[pos{0, 0}] = 1

	panel = paintPanel(panel, instructions)

	xUpBound, xDownBound, yUpBound, yDownBound := getPanelBounds(panel)

	for y := yUpBound; y >= yDownBound; y-- {
		fmt.Printf("\n")
		for x := xDownBound; x <= xUpBound; x++ {
			if lookPositionColorInPanel(panel, pos{x, y}) == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
	}
}
