package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadNumbers() []int {
	var numbers []int
	bytes, _ := ioutil.ReadFile("./day2/input")
	for _, str := range strings.Split(string(bytes), ",") {
		integer, _ := strconv.Atoi(str)
		numbers = append(numbers, integer)
	}
	return numbers
}

func run(numbers []int) int {
	for i := 0; i < len(numbers); i = i + 4 {
		switch numbers[i] {
		case 1:
			numbers[numbers[i+3]] = numbers[numbers[i+1]] + numbers[numbers[i+2]]
		case 2:
			numbers[numbers[i+3]] = numbers[numbers[i+1]] * numbers[numbers[i+2]]
		case 99:
			break
		}
	}
	return numbers[0]
}

func initSequence(origin []int, firstInput int, secondInput int) []int {
	sequence := make([]int, len(origin))
	copy(sequence, origin)
	sequence[1] = firstInput
	sequence[2] = secondInput
	return sequence
}

func exercise2(numbers []int) int {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			// fmt.Printf("%+v:", run(initSequence(numbers, i, j)))
			if run(initSequence(numbers, i, j)) == 19690720 {
				return i*100 + j
			}
		}
	}
	return -1
}

func main() {
	numbers := loadNumbers()
	fmt.Printf("Part1 solution: %d\n", run(initSequence(numbers, 12, 2)))
	fmt.Printf("Part1 solution: %d\n", exercise2(numbers))
}
