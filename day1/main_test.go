package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {

	assert.Equal(t, CalcFuel(100), 31, "The two words values have to be the same.")
	assert.Equal(t, CalcFuel(200), 64, "The two words values have to be the same.")
	assert.Equal(t, CalcFuel(120), 38, "The two words values have to be the same.")
}
