package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	board := LoadBoard("./input")
	maxVisible, _, _ := part1(board)
	assert.Equal(t, 284, maxVisible, "FUCK")
}

func TestExample1(t *testing.T) {
	board := LoadBoard("./example1")
	maxVisible, _, _ := part1(board)
	assert.Equal(t, 8, maxVisible, "FUCK")
}
