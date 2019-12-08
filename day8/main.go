package main

import (
	"fmt"
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
	fmt.Printf("Layers: %d", layers)

	image := [][][]int{}

	for layerIndex := 0; layerIndex < layers; layerIndex++ {
		fmt.Printf("From:%d to %d\n", (layerIndex)*(tall*wide), (layerIndex+1)*(tall*wide))
		layerPixels := pixels[(layerIndex)*(tall*wide) : (layerIndex+1)*(tall*wide)]
		image = append(image, BuildLayer(layerPixels, tall, wide))
	}
	return image
}

func CountOccurrences(numbers []int) map[int]int {
	tempDigitCount := map[int]int{}
	for _, pixel := range numbers {
		_, ok := tempDigitCount[pixel]
		if ok {
			tempDigitCount[pixel] = tempDigitCount[pixel] + 1
		} else {
			tempDigitCount[pixel] = 1
		}
	}
	return tempDigitCount
}

func part1(pixels []int, tall int, wide int) {
	digitsCountPerLayer := []map[int]int{}

	layers := len(pixels) / (tall * wide)

	for layerIndex := 0; layerIndex < layers; layerIndex++ {
		layerPixels := pixels[(layerIndex)*(tall*wide) : (layerIndex+1)*(tall*wide)]

		digitsCountPerLayer = append(digitsCountPerLayer, CountOccurrences(layerPixels))
	}

	minZerosLayerDigitsCount := digitsCountPerLayer[0]
	for _, digitsCount := range digitsCountPerLayer[1:] {
		if digitsCount[0] < minZerosLayerDigitsCount[0] {
			minZerosLayerDigitsCount = digitsCount
		}
	}

	fmt.Printf("Part 1 Solution :%d\n", minZerosLayerDigitsCount[1]*minZerosLayerDigitsCount[2])
}

func main() {
	pixels := loadImage()
	const wide = 25
	const tall = 6

	part1(pixels, tall, wide)
}
