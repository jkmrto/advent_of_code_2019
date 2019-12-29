package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Reaction struct {
	output_name   string
	output_amount int
	input         map[string]int
}

func keys(mymap map[string]int) []string {
	keys := []string{}
	for k := range mymap {
		keys = append(keys, k)
	}
	return keys
}

func getComponentAndAmount(componentAmountStr string) (component string, amount int) {
	componentAmountList := strings.Split(componentAmountStr, " ")
	amount, _ = strconv.Atoi(componentAmountList[0])
	component = componentAmountList[1]
	return
}

func getInputAndOutputStrings(reactionString string) (input string, output string) {
	reactionSplitted := strings.Split(reactionString, " => ")
	input = reactionSplitted[0]
	output = reactionSplitted[1]
	return
}

func parseInput(inputStr string) map[string]int {
	input := make(map[string]int)
	for _, inputComponentStr := range strings.Split(inputStr, ", ") {
		component, amount := getComponentAndAmount(inputComponentStr)
		input[component] = amount
	}
	return input

}

func loadReactions(filepath string) map[string]Reaction {
	bytes, _ := ioutil.ReadFile(filepath)
	reactionsList := strings.Split(string(bytes), "\n")
	reactions := make(map[string]Reaction)

	// var reactions map[string]reaction
	for _, reactionString := range reactionsList {
		inputStr, outputStr := getInputAndOutputStrings(reactionString)

		component, amount := getComponentAndAmount(outputStr)

		r := Reaction{
			output_name:   component,
			output_amount: amount,
			input:         parseInput(inputStr),
		}
		reactions[component] = r
	}

	return reactions
}

func GetReactionsRequired(amountRequired int, r Reaction, stored int) int {
	reactionsCount := 0
	for {
		if stored+reactionsCount*r.output_amount >= amountRequired {
			return reactionsCount
		}
		reactionsCount += 1
	}
}

func GetInputFromStore(r Reaction, amountRequired int, store map[string]int) (map[string]int, map[string]int) {
	var stored int
	if value, exist := store[r.output_name]; exist {
		stored = value
	} else {
		stored = 0
	}

	reactionsRequired := GetReactionsRequired(amountRequired, r, stored)

	scaledInput := make(map[string]int)
	for k := range r.input {
		scaledInput[k] = reactionsRequired * r.input[k]
	}

	store[r.output_name] = stored + reactionsRequired*r.output_amount - amountRequired

	return scaledInput, store
}

func main() {
	defer fmt.Printf("\n")
	reactions := loadReactions("./day14/input")

	missedComponents := reactions["FUEL"].input

	missedComponentsIndex := keys(missedComponents)

	store := make(map[string]int)
	inputRequired := make(map[string]int)

	ORE := 0

	for {
		if len(missedComponentsIndex) == 0 {
			break
		}

		componentEvaluated := missedComponentsIndex[0]
		r := reactions[componentEvaluated]
		requireAmount := missedComponents[componentEvaluated]

		inputRequired, store = GetInputFromStore(r, requireAmount, store)

		for component := range inputRequired {
			inputAmount := inputRequired[component]

			if "ORE" == component {
				ORE += inputAmount
				continue
			} else {
				_, exist := missedComponents[component]
				if exist {
					missedComponents[component] += inputAmount
				} else {
					missedComponents[component] = inputAmount
					missedComponentsIndex = append(missedComponentsIndex, component)
				}
			}
		}

		delete(missedComponents, componentEvaluated)
		missedComponentsIndex = missedComponentsIndex[1:len(missedComponentsIndex)]

	}

	fmt.Printf("\nSolution: %+v", ORE)
}
