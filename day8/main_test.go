package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildLayer(t *testing.T) {
	layer := BuildLayer([]int{1, 2, 3, 4}, 2, 2)
	assert.Equal(t, [][]int{{1, 2}, {3, 4}}, layer, "FUCK")
}

func TestBuildLayer1(t *testing.T) {
	layer := BuildLayer([]int{8, 2, 8, 4, 8, 3}, 3, 2)
	assert.Equal(t, [][]int{{8, 2}, {8, 4}, {8, 3}}, layer, "FUCK")
}

func TestBuildImage1(t *testing.T) {
	pixels := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	image := BuildImage(pixels, 2, 3)
	expected := [][][]int{
		{{1, 2, 3}, {4, 5, 6}},
		{{7, 8, 9}, {10, 11, 12}},
		{{13, 14, 15}, {16, 17, 18}},
	}
	assert.Equal(t, expected, image, "FUCK")
}

func TestBuildImage2(t *testing.T) {
	pixels := []int{0222112222120000}
	image := BuildImage(pixels, 2, 3)
	expected := [][][]int{
		{{1, 2, 3}, {4, 5, 6}},
		{{7, 8, 9}, {10, 11, 12}},
		{{13, 14, 15}, {16, 17, 18}},
	}
	assert.Equal(t, expected, image, "FUCK")
}

func TestCountOccurrences(t *testing.T) {
	pixels := []int{0, 1, 1, 2, 2, 2, 2}
	expected := map[int]int{
		0: 1,
		1: 2,
		2: 4,
	}
	assert.Equal(t, expected, CountOccurrences(pixels), "FUCK")
}
