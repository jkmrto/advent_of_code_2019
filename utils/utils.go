package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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

func IntToSlice(input int) []int {
	s := strconv.Itoa(input)
	slice := []int{}

	for i := 0; i < len(s); i++ {
		value, _ := strconv.Atoi(string(s[i]))
		slice = append(slice, value)
	}
	return slice
}

func AddPading(slice []int, digits int) []int {
	finalSlice := make([]int, digits)
	offset := digits - len(slice)

	for i := 0; i < len(slice); i++ {
		finalSlice[offset+i] = slice[i]
	}
	return finalSlice
}

func Max(values []int) int {
	max := 0

	for _, value := range values {
		if value > max {
			max = value
		}
	}

	return max
}

func LoadInstructions(path string) []int {
	var numbers []int
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		log.Panic(err)
	}
	for _, str := range strings.Split(string(bytes), ",") {
		integer, _ := strconv.Atoi(str)
		numbers = append(numbers, integer)
	}
	return numbers
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func DivMod(numerator int, denominator int) (quotient int, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
