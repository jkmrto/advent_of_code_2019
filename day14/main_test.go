package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore1(t *testing.T) {
	input := make(map[string]int)
	input["1"] = 1

	r := Reaction{
		output_name:   "example",
		output_amount: 10,
		input:         input,
	}

	store := make(map[string]int)
	input, store = GetInputFromStore(r, 10, store)

	assert.Equal(t, input["1"], 1, "FUCK")
	assert.Equal(t, store["output"], 0, "FUCK")
}

func TestStore15(t *testing.T) {
	input := make(map[string]int)
	input["1"] = 1

	r := Reaction{
		output_name:   "output",
		output_amount: 15,
		input:         input,
	}

	store := make(map[string]int)
	input, store = GetInputFromStore(r, 10, store)

	assert.Equal(t, input["1"], 1, "FUCK")
	assert.Equal(t, store["output"], 5, "FUCK")
}

func TestStoreUsingReserves(t *testing.T) {
	input := make(map[string]int)
	input["1"] = 1

	r := Reaction{
		output_name:   "output",
		output_amount: 10,
		input:         input,
	}

	store := make(map[string]int)
	store["output"] = 5
	input, store = GetInputFromStore(r, 15, store)

	assert.Equal(t, input["1"], 1, "FUCK")
	assert.Equal(t, store["output"], 0, "FUCK")
}

func TestStoreSavingReserves(t *testing.T) {
	input := make(map[string]int)
	input["1"] = 2

	r := Reaction{
		output_name:   "output",
		output_amount: 3,
		input:         input,
	}

	store := make(map[string]int)
	store["output"] = 5
	amountRequired := 10
	input, store = GetInputFromStore(r, amountRequired, store)

	assert.Equal(t, input["1"], 4, "FUCK")
	assert.Equal(t, store["output"], 1, "FUCK")
}

func TestStoreSavingReserves1(t *testing.T) {
	input := make(map[string]int)
	input["1"] = 1

	r := Reaction{
		output_name:   "output",
		output_amount: 2,
		input:         input,
	}

	store := make(map[string]int)
	store["output"] = 3
	amountRequired := 5
	input, store = GetInputFromStore(r, amountRequired, store)

	assert.Equal(t, input["1"], 1, "FUCK")
	assert.Equal(t, store["output"], 0, "FUCK")
}
