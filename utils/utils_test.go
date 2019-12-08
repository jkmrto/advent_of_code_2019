package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOccurrences(t *testing.T) {
	pixels := []int{0, 1, 1, 2, 2, 2, 2}
	expected := map[int]int{
		0: 1,
		1: 2,
		2: 4,
	}
	assert.Equal(t, expected, CountOccurrences(pixels), "FUCK")
}

func TestIntToSLice(t *testing.T) {
	assert.Equal(t, IntToSlice(4567), []int{4, 5, 6, 7}, "The two words values have to be the same.")

}

func TestAddPadding(t *testing.T) {
	assert.Equal(t, AddPading([]int{4, 5, 6}, 5), []int{0, 0, 4, 5, 6}, "The two words values have to be the same.")

}
