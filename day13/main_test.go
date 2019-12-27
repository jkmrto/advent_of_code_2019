package main

import (
	"testing"

	"github.com/advent_of_code_2019/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	instructions := utils.LoadInstructions("./input")
	assert.Equal(t, Part1(instructions), 280, "FUCK")

}
