package main

import (
	"fmt"
	"github.com/advent_of_code_2019/utils"
	"io/ioutil"
	"strconv"
)

func loadImage() []int {
	var numbers []int
	bytes, _ := ioutil.ReadFile("./day8/input")

	for _, str := range bytes {
		integer, _ := strconv.Atoi(string(str))
		numbers = append(numbers, integer)

	}
	return numbers
}

func BuildLayer(pixels []int, tall int, wide int) [][]int {
	layer := [][]int{}

	for rowIndex := 0; rowIndex < tall; rowIndex++ {
		// fmt.Printf("Row from:%d to %d\n", (rowIndex)*(wide), (rowIndex+1)*(wide))
		row := pixels[rowIndex*wide : (rowIndex+1)*wide]
		layer = append(layer, row)
	}
	return layer
}

func BuildImage(pixels []int, tall int, wide int) [][][]int {
	layers := len(pixels) / (tall * wide)

	image := [][][]int{}

	for layerIndex := 0; layerIndex < layers; layerIndex++ {
		// fmt.Printf("From:%d to %d\n", (layerIndex)*(tall*wide), (layerIndex+1)*(tall*wide))
		layerPixels := pixels[(layerIndex)*(tall*wide) : (layerIndex+1)*(tall*wide)]
		image = append(image, BuildLayer(layerPixels, tall, wide))
	}
	return image
}

func part1(pixels []int, tall int, wide int) int {
	digitsCountPerLayer := []map[int]int{}

	layers := len(pixels) / (tall * wide)

	for layerIndex := 0; layerIndex < layers; layerIndex++ {
		layerPixels := pixels[(layerIndex)*(tall*wide) : (layerIndex+1)*(tall*wide)]

		digitsCountPerLayer = append(digitsCountPerLayer, utils.CountOccurrences(layerPixels))
	}

	minZerosLayerDigitsCount := digitsCountPerLayer[0]
	for _, digitsCount := range digitsCountPerLayer[1:] {
		if digitsCount[0] < minZerosLayerDigitsCount[0] {
			minZerosLayerDigitsCount = digitsCount
		}
	}

	return minZerosLayerDigitsCount[1] * minZerosLayerDigitsCount[2]
}

func selectFirstVisiblePixel(image [][][]int, row int, column int) int {
	for _, layer := range image {
		if layer[row][column] != 2 {
			return layer[row][column]
		}
	}
	return -1
}

func Part2(pixels []int, tall int, wide int) [][]int {
	image := BuildImage(pixels, tall, wide)

	renderedImage := [][]int{}
	for row := 0; row < tall; row++ {
		renderedRow := make([]int, wide)
		for column := 0; column < wide; column++ {
			renderedRow[column] = selectFirstVisiblePixel(image, row, column)
		}
		renderedImage = append(renderedImage, renderedRow)
	}
	return renderedImage
}

func main() {
	pixels := loadImage()
	const wide = 25
	const tall = 6

	fmt.Printf("Part 1 Solution :%d\n", part1(pixels, 6, 25))
	part2Solution := Part2(pixels, 6, 25)

	//Loop to print part2 Solution
	for _, row := range part2Solution {
		fmt.Printf("%+v\n", row)
	}
}
