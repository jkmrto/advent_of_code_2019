package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func CalcFuel(weight int) int {
	return weight/3 - 2
}

func CalcFuelRecursive(weight int) int {
	result := weight/3 - 2
	if result > 0 {
		return result + CalcFuelRecursive(result)
	} else {
		return 0
	}
}

func loadNumbers() []int {
	var numbers []int
	bytes, _ := ioutil.ReadFile("./day1/input")

	for _, str := range strings.Split(string(bytes), "\n") {
		integer, _ := strconv.Atoi(str)
		numbers = append(numbers, integer)
	}
	return numbers
}

func main() {

	var result int
	var resultRecursive int

	for _, module := range loadNumbers() {
		result += CalcFuel(module)
		resultRecursive += CalcFuelRecursive(module)
	}

	fmt.Printf("Part 1 result: %d\n", result)
	fmt.Printf("Part 2 result: %d\n", resultRecursive)
}
