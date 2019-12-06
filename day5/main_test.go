package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {

	assert.Equal(t, intToSlice(4567), []int{4,5,6,7}, "The two words values have to be the same.")
	assert.Equal(t, addPading([]int{4,5,6}, 5), []int{0,0, 4, 5 ,6}, "The two words values have to be the same.")
	
	var opCode int
	var parameterMode []int

	opCode, parameterMode = ProcessInstruction(3)
	assert.Equal(t, opCode, 3, "The two words values have to be the same.")
	assert.Equal(t, parameterMode, []int{0,0,0}, "The two words values have to be the same.")
	
	assert.Equal(t, ProcessInstruction(1110), []int{0, 1, 1, 1 , 0}, "The two words values have to be the same.")
	assert.Equal(t, ProcessInstruction(1), []int{0, 0, 0, 0 , 1}, "The two words values have to be the same.")
}
