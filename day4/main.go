package main

import "fmt"

import "math"

var upRange = 171309
var downRange = 643603

func adjustToIncreaseDigits(input []int) int {
	return 0
}

func intToArray(value int) []int {
	slice := []int{}
	digit := len()

	for value > 0 {
		rest := int(value % int(math.Pow10(digit)))
		value = value - rest
		slice = append(slice, (rest / int(math.Pow10(digit-1))))
		fmt.Printf("value: %d, Residuo; %d, array; %d\n, ", value, rest, slice)
		digit++
	}

	return slice
}

func main() {
	fmt.Printf("%d\n", upRange)
	initial := intToArray(upRange)
	fmt.Printf("%v", initial)

}
