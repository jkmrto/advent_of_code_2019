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

func TestBuildImage3(t *testing.T) {
	pixels := []int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0}
	image := BuildImage(pixels, 2, 2)
	expected := [][][]int{
		{{0, 2}, {2, 2}},
		{{1, 1}, {2, 2}},
		{{2, 2}, {1, 2}},
		{{0, 0}, {0, 0}},
	}
	assert.Equal(t, expected, image, "FUCK")
}

func TestPart2Example(t *testing.T) {
	pixels := []int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0}
	image := Part2(pixels, 2, 2)
	expected := [][]int{{0, 1}, {1, 0}}
	assert.Equal(t, expected, image, "FUCK")
}
