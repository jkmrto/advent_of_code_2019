package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

type crossPoint struct {
	point   pos
	seq0Pos int
	seq1Pos int
}

var instructionToMove = map[string]pos{
	"U": pos{0, 1},
	"D": pos{0, -1},
	"R": pos{1, 0},
	"L": pos{-1, 0},
}

func loadInstructions() ([]string, []string) {
	bytes, _ := ioutil.ReadFile("./day3/input")
	seq := strings.Split(string(bytes), "\n")
	return strings.Split(seq[0], ","), strings.Split(seq[1], ",")
}

func findCrossPoints(seq0 []pos, seq1 []pos) []crossPoint {
	crossPoints := []crossPoint{}
	for i, point1 := range seq0 {
		for j, point2 := range seq1 {
			if point1.x == 0 && point1.y == 0 {
				continue
			}
			if (point1.x == point2.x) && (point1.y == point2.y) {
				crossPoints = append(crossPoints, crossPoint{point1, i, j})
				break
			}
		}
	}

	return crossPoints
}

func appendMovements(positions []pos, move pos, n_moves int) []pos {
	newPositions := buildPositions(positions[len(positions)-1], move, n_moves)
	return append(positions, newPositions...)
}

func buildPositions(origin pos, move pos, n_moves int) []pos {
	newPositions := []pos{}
	previousPos := origin
	for i := 0; i < n_moves; i++ {
		newPositions = append(newPositions, pos{previousPos.x + move.x, previousPos.y + move.y})
		previousPos = pos{previousPos.x + move.x, previousPos.y + move.y}
	}
	return newPositions
}

func buildPath(instructions []string) []pos {
	var move pos
	positions := []pos{pos{0, 0}}

	for _, instruction := range instructions {
		n_moves, _ := strconv.Atoi(string(instruction[1:]))
		move = instructionToMove[string(instruction[0])]
		positions = appendMovements(positions, move, n_moves)
	}

	return positions
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findNearestCrossPoint(crossPoints []crossPoint) int {
	minDistance := Abs(crossPoints[0].point.x) + Abs(crossPoints[0].point.y)
	var tempDistance int
	for _, crossPoint := range crossPoints[1:] {
		tempDistance = Abs(crossPoint.point.x) + Abs(crossPoint.point.y)
		if minDistance > tempDistance {
			minDistance = tempDistance
		}
	}
	return minDistance
}

func findSoonerCrossPoint(crossPoints []crossPoint) int {
	minCrossPointSteps := crossPoints[0].seq0Pos + crossPoints[0].seq1Pos
	var tempSteps int
	for _, crossPoint := range crossPoints[1:] {
		tempSteps = crossPoint.seq0Pos + crossPoint.seq1Pos
		if minCrossPointSteps > tempSteps {
			minCrossPointSteps = tempSteps
		}
	}
	return minCrossPointSteps
}

func main() {
	seq0, seq1 := loadInstructions()

	path0 := buildPath(seq0)
	path1 := buildPath(seq1)

	crossPoints := findCrossPoints(path0, path1)
	// fmt.Printf("%+v\n", crossPoints)

	distance := findNearestCrossPoint(crossPoints)
	fmt.Printf("Part1: %+v\n", distance)

	minSteps := findSoonerCrossPoint(crossPoints)
	fmt.Printf("Part2: %+v\n", minSteps)
}

// func test() {
// 	seq0 := []pos{pos{1, 0}, pos{2, 0}, pos{1, 0}, pos{2, 0}}
// 	seq1 := []pos{pos{1, 1}, pos{2, 2}, pos{1, 1}, pos{2, 2}}

// 	crossPoints := findCrossPoints(seq0, seq1)
// 	fmt.Printf("%+v", crossPoints)

// 	distance := findNearestCrossPoint(crossPoints)
// 	fmt.Printf("%+v", distance)
// }

// func test1() {
// 	seq0 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
// 	seq1 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}

// 	path0 := buildPath(seq0)
// 	path1 := buildPath(seq1)

// 	crossPoints := findCrossPoints(path0, path1)
// 	fmt.Printf("%+v\n", crossPoints)

// 	distance := findNearestCrossPoint(crossPoints)
// 	fmt.Printf("%+v\n", distance)

// 	minSteps := findSoonerCrossPoint(crossPoints)
// 	fmt.Printf("%+v\n", minSteps)
// }

// func test2() {
// 	seq0 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
// 	seq1 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}

// 	path0 := buildPath(seq0)
// 	path1 := buildPath(seq1)

// 	crossPoints := findCrossPoints(path0, path1)
// 	fmt.Printf("%+v\n", crossPoints)

// 	distance := findNearestCrossPoint(crossPoints)
// 	fmt.Printf("%+v\n", distance)

// 	minSteps := findSoonerCrossPoint(crossPoints)
// 	fmt.Printf("%+v\n", minSteps)
// }
