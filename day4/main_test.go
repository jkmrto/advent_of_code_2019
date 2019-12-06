package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	assert.Equal(t, IsAnyIsolatedCoupleEqualDigits([]int{2,2,2,2}), false, "FUCK")
	assert.Equal(t, IsAnyIsolatedCoupleEqualDigits([]int{2,2,4,4}), true, "FUCK")
	assert.Equal(t, IsAnyIsolatedCoupleEqualDigits([]int{3,3,3,5}), false, "FUCK")
	assert.Equal(t, IsAnyIsolatedCoupleEqualDigits([]int{2,5,3,3}), true, "FUCK")
	assert.Equal(t, IsAnyIsolatedCoupleEqualDigits([]int{2,2}), true, "FUCK")
	assert.Equal(t, IsAnyIsolatedCoupleEqualDigits([]int{1,2,3,4}), false, "FUCK")
}
