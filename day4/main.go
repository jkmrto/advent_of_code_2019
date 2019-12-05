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
		value = value - quotient * int(math.Pow10(digitPos)) 
		slice = append(slice, quotient)
		digitPos--
	}
	return slice
}


func equalAdjuntDigits(slice []int, n_adjunts int) bool{
	return areAnyEquals(segmentByAdjunts(slice, n_adjunts))
}

func segmentByAdjunts(slice []int, n_adjunts int) [][]int{
	sliceAdjunts := [][]int{}
	for i := 0; i < len(slice) + 1 - n_adjunts; i++{
		sliceAdjunts = append(sliceAdjunts, slice[i:(n_adjunts + i)])
		// fmt.Printf("%+v\n", sliceAdjunts)
	}

	return sliceAdjunts
}

func areAnyEquals(sliceAdjunts [][]int) bool {
	var expected int
	var equal bool
	for _, slice := range sliceAdjunts{
		expected = slice[0]
		equal = true
		for _, number := range slice[1:]{
			equal = equal && (expected == number)
		}
		if equal {
			return true
		}
	}
	return false
}

func isIncreasing(slice []int) bool{
	latestValue := slice[0]
	for _, value := range(slice) {
		if value >= latestValue{
			latestValue = value
		}else{
			return false
		}
	}
	return true 
}

func main() {	
	upRange := 171309
	downRange := 643603

	var slice []int
	counter := 0

	for value:= upRange; value < downRange; value++ {
		slice = intToSlice(value)

		if isIncreasing(slice) {
			// fmt.Printf("\nIs increasing: %+v", slice)
			if equalAdjuntDigits(slice, 2) {
				counter++
				// fmt.Printf("\t Is adjacents equal: %+v", slice)
			}
		}
	}

	fmt.Printf("\n\n Numbers found: %d", counter)
}
