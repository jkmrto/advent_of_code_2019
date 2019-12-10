package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/advent_of_code_2019/utils"
)

func keys(mymap map[angle][]int) []angle {
	keys := []angle{}
	for k := range mymap {
		keys = append(keys, k)
	}
	return keys
}

func LoadBoard(path string) [][]string {
	var board [][]string
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		log.Panic(err)
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		tempRow := []string{}
		for _, element := range line {
			tempRow = append(tempRow, string(element))
		}

		board = append(board, tempRow)
	}
	return board
}

func reduceRational(num int, den int) (numFinal int, denFinal int, ratio int) {
	lcm := utils.GCD(utils.Abs(num), utils.Abs(den))
	// fmt.Printf("\t lcm: %d", lcm)

	if lcm == 0 {
		numFinal = num
		denFinal = den
		ratio = 1
	} else {
		numFinal = num / lcm
		denFinal = den / lcm
		ratio = lcm
	}

	return
}

type angle struct {
	X, Y int
}

func buildRadar(board [][]string, stationPosY int, stationPosX int) int {
	radar := map[angle][]int{}
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] == "#" && (y != stationPosY || x != stationPosX) {
				// fmt.Printf("Asteroid at: (%d, %d)", x, y)
				// fmt.Printf("\tRelative to Station(%d, %d)", x-stationPosX, y-stationPosY)
				ang, step := calculateAngleAndSteps(x-stationPosX, y-stationPosY)
				// fmt.Printf("\tReduced to Station %+v Ratio: %d\n", ang, step)
				addToRadar(radar, ang, step)

			}
		}
	}

	// fmt.Printf("%+v", radar)
	// fmt.Printf("\n%+v\n", len(keys(radar)))

	return len(keys(radar))

}

func calculateAngleAndSteps(distanceX int, distanceY int) (ang angle, steps int) {
	distanceX, distanceY, steps = reduceRational(distanceX, distanceY)
	ang = angle{distanceX, distanceY}
	return
}

func addToRadar(asteroidRef map[angle][]int, ang angle, ratio int) {
	value, exist := asteroidRef[ang]
	if exist {
		asteroidRef[ang] = append(value, ratio)
	} else {
		asteroidRef[ang] = []int{ratio}
	}
}

func part1(board [][]string) (maxVisibleAsteroids int, maxVisibleAsteroidsPosX int, maxVisibleAsteroidsPosY int) {
	maxVisibleAsteroids = 0

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] == "#" {
				visibleAsteroids := buildRadar(board, y, x)
				if visibleAsteroids > maxVisibleAsteroids {
					maxVisibleAsteroidsPosX = x
					maxVisibleAsteroidsPosX = y
					maxVisibleAsteroids = visibleAsteroids
				}
			}
		}
	}
	return
}

func main() {
	// board := LoadBoard("./day10/example1")
	board := LoadBoard("./day10/input")

	maxVisible, posX, posY := part1(board)
	fmt.Printf("Part2 Solution: (%d, %d): %d\n", posX, posY, maxVisible)

}
