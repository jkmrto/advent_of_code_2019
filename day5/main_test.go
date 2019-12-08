package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {

	var opCode int
	var parameterMode []int

	opCode, parameterMode = ProcessInstruction(3)
	assert.Equal(t, opCode, 3, "The two words values have to be the same.")
	assert.Equal(t, parameterMode, []int{0, 0, 0}, "The two words values have to be the same.")

	// This has to be fixed
	opCode, parameterMode = ProcessInstruction(1110)
	assert.Equal(t, parameterMode, []int{0, 1, 1}, "FUCK")
	assert.Equal(t, opCode, 10, "FUCK")

	opCode, parameterMode = ProcessInstruction(1)
	assert.Equal(t, parameterMode, []int{0, 0, 0}, "FUCK")
	assert.Equal(t, opCode, 1, "FUCK")

}
