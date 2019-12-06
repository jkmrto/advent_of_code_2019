package main

import "fmt"
import "math"

func divmod(numerator int, denominator int) (quotient int, remainder int) {
	quotient = numerator / denominator
	remainder = numerator % denominator
	return
}

func intToSlice(value int) []int {
	slice := []int{}
	digitPos := 5

	for digitPos >= 0 {
		quotient, _ := divmod(value, int(math.Pow10(digitPos)))
		value = value - quotient*int(math.Pow10(digitPos))
		slice = append(slice, quotient)
		digitPos--
	}
	return slice
}

func equalAdjuntDigits(slice []int, n_adjunts int) bool {
	return areAnyEquals(segmentByAdjunts(slice, n_adjunts))
}

func segmentByAdjunts(slice []int, n_adjunts int) [][]int {
	sliceAdjunts := [][]int{}
	for i := 0; i < len(slice)+1-n_adjunts; i++ {
		sliceAdjunts = append(sliceAdjunts, slice[i:(n_adjunts+i)])
		// fmt.Printf("%+v\n", sliceAdjunts)
	}

	return sliceAdjunts
}

func areAnyEquals(sliceAdjunts [][]int) bool {
	var expected int
	var equal bool
	for _, slice := range sliceAdjunts {
		expected = slice[0]
		equal = true
		for _, number := range slice[1:] {
			equal = equal && (expected == number)
		}
		if equal {
			return true
		}
	}
	return false
}

func isIncreasing(slice []int) bool {
	latestValue := slice[0]
	for _, value := range slice {
		if value >= latestValue {
			latestValue = value
		} else {
			return false
		}
	}
	return true
}

func IsAnyIsolatedCoupleEqualDigits(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == slice[i+1] {
			if i > 0 && slice[i-1] == slice[i] {
				continue
			}
			if i < len(slice)-2 && slice[i+2] == slice[i] {
				continue
			}

			return true
		}
	}
	return false
}

func main() {
	upRange := 171309
	downRange := 643603

	var slice []int
	part1Counter := 0
	part2Counter := 0

	for value := upRange; value < downRange; value++ {
		slice = intToSlice(value)

		if isIncreasing(slice) {
			fmt.Printf("\nIs increasing: %+v", slice)
			if equalAdjuntDigits(slice, 2) {
				part1Counter++
				fmt.Print("\t Is adjacents equal")

				if IsAnyIsolatedCoupleEqualDigits(slice) {
					fmt.Print("\t There is an isolated couple of digits")
					part2Counter++
				}
			}
		}
	}

	fmt.Printf("\n\n Part 1 Solution: %d\n", part1Counter)
	fmt.Printf("\n\n Part 2 Solution: %d\n", part2Counter)
}
