package main

import (
	"fmt"
	"reflect"

	"github.com/advent_of_code_2019/utils"
)

type coordinates struct {
	x int
	y int
	z int
}

func (c *coordinates) getField(field string) int {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func getField(v *coordinates, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

type moon struct {
	position coordinates
	velocity coordinates
}

// Energy given by the movement
func (m moon) calculateKineticEnergy() int {
	return utils.Abs(m.velocity.x) + utils.Abs(m.velocity.y) + utils.Abs(m.velocity.z)
}

// Energy given by the position
func (m moon) calculatePotentialEnergy() int {
	return utils.Abs(m.position.x) + utils.Abs(m.position.y) + utils.Abs(m.position.z)
}

// Energy given by the position
func (m moon) calculateEnergy() int {
	return m.calculateKineticEnergy() * m.calculatePotentialEnergy()
}

func updateVelocity(m moon, gravity coordinates) moon {
	return moon{
		position: m.position,
		velocity: coordinates{
			m.velocity.x + gravity.x,
			m.velocity.y + gravity.y,
			m.velocity.z + gravity.z,
		},
	}
}

func updatePosition(m moon) moon {
	return moon{
		position: coordinates{
			m.position.x + m.velocity.x,
			m.position.y + m.velocity.y,
			m.position.z + m.velocity.z,
		},
		velocity: m.velocity,
	}
}

// <x=-7, y=-8, z=9>
// <x=-12, y=-3, z=-4>
// <x=6, y=-17, z=-9>
// <x=4, y=-10, z=-6>

func loadInput() []moon {

	//	<x=-7, y=-8, z=9>
	moon1 := moon{
		coordinates{-7, -8, 9},
		coordinates{0, 0, 0},
	}

	// <x=-12, y=-3, z=-4>
	moon2 := moon{
		coordinates{-12, -3, -4},
		coordinates{0, 0, 0},
	}

	// <x=6, y=-17, z=-9>
	moon3 := moon{
		coordinates{6, -17, -9},
		coordinates{0, 0, 0},
	}

	// <x=4, y=-10, z=-6>
	moon4 := moon{
		coordinates{4, -10, -6},
		coordinates{0, 0, 0},
	}

	return []moon{moon1, moon2, moon3, moon4}
}

func GetGravity(evaluatedMoon moon, galaxy []moon) coordinates {
	var gravity coordinates
	var gravityToMoon coordinates

	for _, moon := range galaxy {
		if moon == evaluatedMoon {
			continue
		}

		gravityToMoon = GetGravityToMoon(evaluatedMoon, moon)
		gravity.x += gravityToMoon.x
		gravity.y += gravityToMoon.y
		gravity.z += gravityToMoon.z
	}

	return gravity
}

func GetGravityToMoon(object_moon moon, mass_moon moon) coordinates {
	return coordinates{
		x: getGravityByCoordinate(object_moon.position.x, mass_moon.position.x),
		y: getGravityByCoordinate(object_moon.position.y, mass_moon.position.y),
		z: getGravityByCoordinate(object_moon.position.z, mass_moon.position.z),
	}
}

func getGravityByCoordinate(object_pos int, mass_pos int) int {
	if object_pos > mass_pos {
		return -1
	} else if object_pos < mass_pos {
		return 1
	} else {
		return 0
	}
}

func UpdateVelocities(moons []moon) []moon {

	var updatedMoons = make([]moon, len(moons))
	var gravity coordinates
	for i, moon := range moons {
		gravity = GetGravity(moon, moons)
		updatedMoons[i] = updateVelocity(moon, gravity)
	}
	return updatedMoons
}

func UpdatePositions(moons []moon) []moon {

	var updatedMoons = make([]moon, len(moons))
	for i, moon := range moons {
		updatedMoons[i] = updatePosition(moon)
	}
	return updatedMoons
}

func Simulate(moons []moon, steps int) []moon {

	for i := 1; i <= steps; i++ {
		moons = UpdateVelocities(moons)
		moons = UpdatePositions(moons)
	}

	return moons
}

func calculateUniversEnergy(moons []moon) int {
	var sum int = 0

	for _, moon := range moons {
		sum += moon.calculateEnergy()
	}

	return sum
}

func FindPeriodicity(moonsOrigin []moon) int {
	moons := make([]moon, len(moonsOrigin))
	copy(moons, moonsOrigin)
	periodX := findAxisPeriodicity(moonsOrigin, moons, "x")
	periodY := findAxisPeriodicity(moonsOrigin, moons, "y")
	periodZ := findAxisPeriodicity(moonsOrigin, moons, "z")
	return utils.LCM(periodX, periodY, periodZ)
}

func findAxisPeriodicity(moonsOrigin []moon, moons []moon, axis string) int {
	var step int = 0
	for {
		step++
		moons = UpdateVelocities(moons)
		moons = UpdatePositions(moons)

		if isSamePosition(moonsOrigin, moons, axis) {
			return step

		}
	}
}

func isSamePosition(moonsOrigin []moon, moons []moon, axis string) bool {
	for i := 0; i < len(moonsOrigin); i++ {
		if (&moons[i].position).getField(axis) != (&moonsOrigin[i].position).getField(axis) {
			return false
		}

		if (&moons[i].velocity).getField(axis) != (&moonsOrigin[i].velocity).getField(axis) {
			return false
		}

	}
	return true
}

func part1() {
	moons := loadInput()
	moons1000 := Simulate(moons, 1000)
	fmt.Printf("\n Part1 Solution: %d", calculateUniversEnergy(moons1000))
}

func main() {
	part1()

	moons := loadInput()
	period := FindPeriodicity(moons)
	fmt.Printf("\n Part2 solution: %+v\n", period)
}
