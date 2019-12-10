package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
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

func buildRadar(board [][]string, stationPosY int, stationPosX int) map[angle][]int {
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

	return radar

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
				radar := buildRadar(board, y, x)
				visibleAsteroids := len(keys(radar))
				if visibleAsteroids > maxVisibleAsteroids {
					maxVisibleAsteroidsPosX = x
					maxVisibleAsteroidsPosY = y
					maxVisibleAsteroids = visibleAsteroids
				}
			}
		}
	}
	return
}

type asteroidsRelativeToStation struct {
	degrees          float64
	ang              angle
	stepsToAsteroids []int
}

func buildRadarWithDegrees(board [][]string, posX int, posY int) []asteroidsRelativeToStation {
	radar := buildRadar(board, posY, posX)
	radarWithDegrees := []asteroidsRelativeToStation{}
	for ang, stepsToAsteroids := range radar {
		sort.Ints(stepsToAsteroids)
		asteroidsRelativePosition := asteroidsRelativeToStation{
			degrees:          CoordinatesToDegrees(ang.X, ang.Y),
			ang:              ang,
			stepsToAsteroids: stepsToAsteroids,
		}
		radarWithDegrees = append(radarWithDegrees, asteroidsRelativePosition)

	}

	sort.Slice(radarWithDegrees, func(i, j int) bool {
		return radarWithDegrees[i].degrees < radarWithDegrees[j].degrees
	})

	return radarWithDegrees
}

func buildPositionWithRadar(posToStation asteroidsRelativeToStation, index int, stationPosX int, stationPosY int) (posX int, posY int) {
	// fmt.Printf("%d, %d", posToStation.ang.X*posToStation.stepsToAsteroids[index], posToStation.ang.Y*posToStation.stepsToAsteroids[index])
	posX = stationPosX + posToStation.ang.X*posToStation.stepsToAsteroids[index]
	posY = stationPosY + posToStation.ang.Y*posToStation.stepsToAsteroids[index]
	return
}

func main() {
	// board := LoadBoard("./day10/example1")
	// board := LoadBoard("./day10/example2")
	// board := LoadBoard("./day10/example3")

	board := LoadBoard("./day10/input")

	maxVisible, posX, posY := part1(board)
	fmt.Printf("Part2 Solution: (%d, %d): %d\n", posX, posY, maxVisible)

	radarWithDegrees := buildRadarWithDegrees(board, posX, posY)

	i := 0
	counter := 1
	for {
		asteroidX, asteroidY := buildPositionWithRadar(radarWithDegrees[i], 0, posX, posY)
		fmt.Printf("\nCounter: %d, Position: (%d, %d) Removing element: %+v", counter, asteroidX, asteroidY, radarWithDegrees[i])
		if len(radarWithDegrees[i].stepsToAsteroids) > 1 {
			radarWithDegrees[i].stepsToAsteroids = radarWithDegrees[i].stepsToAsteroids[1:]
			i++
		} else {
			radarWithDegrees = remove(radarWithDegrees, i)
		}

		if i == len(radarWithDegrees) {
			fmt.Printf("\n>>>>>> Starting Other Round <<<<<<")
			i = 0
		}

		if len(radarWithDegrees) == 0 {
			break
		}
		counter++
	}

}

func remove(slice []asteroidsRelativeToStation, s int) []asteroidsRelativeToStation {
	return append(slice[:s], slice[s+1:]...)
}

func CoordinatesToDegrees(x int, y int) float64 {
	if y < 0 && x > 0 {
		return math.Atan(float64(utils.Abs(x))/float64(utils.Abs(y))) * 180 / math.Pi
	} else if y > 0 && x > 0 {
		return 180 - math.Atan(float64(utils.Abs(x))/float64(utils.Abs(y)))*180/math.Pi
	} else if y > 0 && x < 0 { // 3 cuadrante
		return 180 + math.Atan(float64(utils.Abs(x))/float64(utils.Abs(y)))*180/math.Pi
	} else if y < 0 && x < 0 { // 4 cuadrante
		return 360 - math.Atan(float64(utils.Abs(x))/float64(utils.Abs(y)))*180/math.Pi
	} else { // Special cases
		if y == -1 && x == 0 {
			return float64(0)
		} else if y == 0 && x == 1 {
			return float64(90)
		} else if y == 1 && x == 0 {
			return float64(180)
		} else if y == 0 && x == -1 {
			return float64(270)
		} else {
			log.Panic("Something weird has happened")
		}
	}

	return float64(-1)
}
