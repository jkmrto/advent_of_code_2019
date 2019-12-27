package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGravityToMoon(t *testing.T) {

	moon111 := moon{
		coordinates{1, 1, 1},
		coordinates{0, 0, 0},
	}

	moon222 := moon{
		coordinates{2, 2, 2},
		coordinates{0, 0, 0},
	}

	moon000 := moon{
		coordinates{0, 0, 0},
		coordinates{0, 0, 0},
	}

	assert.Equal(t, coordinates{1, 1, 1}, GetGravityToMoon(moon111, moon222), "FUCK")
	assert.Equal(t, coordinates{-1, -1, -1}, GetGravityToMoon(moon111, moon000), "FUCK")
	assert.Equal(t, coordinates{0, 0, 0}, GetGravityToMoon(moon111, moon111), "FUCK")
}

func testInput() []moon {
	//	<x=-1, y=0, z=2>
	moon1 := moon{
		coordinates{-1, 0, 2},
		coordinates{0, 0, 0},
	}

	// <x=2, y=-10, z=-7>
	moon2 := moon{
		coordinates{2, -10, -7},
		coordinates{0, 0, 0},
	}

	// <x=4, y=-8, z=8>
	moon3 := moon{
		coordinates{4, -8, 8},
		coordinates{0, 0, 0},
	}

	moon4 := moon{
		coordinates{3, 5, -1},
		coordinates{0, 0, 0},
	}

	return []moon{moon1, moon2, moon3, moon4}

}

func TestGalaxyStep1(t *testing.T) {
	moons := testInput()
	moons = UpdateVelocities(moons)
	moons = UpdatePositions(moons)

	// pos=<x= 2, y=-1, z= 1>, vel=<x= 3, y=-1, z=-1>
	moon1Step1 := moon{
		coordinates{2, -1, 1},
		coordinates{3, -1, -1},
	}

	// pos=<x= 3, y=-7, z=-4>, vel=<x= 1, y= 3, z= 3>
	moon2Step1 := moon{
		coordinates{3, -7, -4},
		coordinates{1, 3, 3},
	}

	// pos=<x= 1, y=-7, z= 5>, vel=<x=-3, y= 1, z=-3>
	moon3Step1 := moon{
		coordinates{1, -7, 5},
		coordinates{-3, 1, -3},
	}

	//pos=<x= 2, y= 2, z= 0>, vel=<x=-1, y=-3, z= 1>
	moon4Step1 := moon{
		coordinates{2, 2, 0},
		coordinates{-1, -3, 1},
	}

	assert.Equal(t, moon1Step1, moons[0], "FUCK")
	assert.Equal(t, moon2Step1, moons[1], "FUCK")
	assert.Equal(t, moon3Step1, moons[2], "FUCK")
	assert.Equal(t, moon4Step1, moons[3], "FUCK")
}

func TestEnergyStep10(t *testing.T) {
	moons := testInput()
	moons10 := Simulate(moons, 10)

	calculateUniversEnergy(moons10)
	assert.Equal(t, calculateUniversEnergy(moons10), 179, "FUCK")

}

func TestPart1(t *testing.T) {
	moons := loadInput()
	moons1000 := Simulate(moons, 1000)
	assert.Equal(t, calculateUniversEnergy(moons1000), 12773, "FUCK")
}
