package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func keyExist(hash map[string]string, key string) bool {
	_, ok := hash[key]
	return ok
}

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}

func buildOrbitansToOrbits(orbitsRelation []string) map[string]string {
	orbitantToOrbit := make(map[string]string)
	for _, orbitRelation := range orbitsRelation {
		splitObitRelation := strings.Split(orbitRelation, ")")
		object := splitObitRelation[0]
		orbitant := splitObitRelation[1]
		orbitantToOrbit[orbitant] = object

	}
	return orbitantToOrbit
}

func doBuildOrbitPath(orbitantsToOrbits map[string]string, orbitant string, orbitPath []string) []string {
	if keyExist(orbitantsToOrbits, orbitant) {
		orbit, _ := orbitantsToOrbits[orbitant]
		updatedOrbitPath := append(orbitPath, orbit)
		return doBuildOrbitPath(orbitantsToOrbits, orbit, updatedOrbitPath)
	} else {
		return orbitPath
	}
}

func buildOrbitPath(orbitantsToOrbits map[string]string, orbitant string) []string {
	orbitPath := []string{orbitant}
	return reverse(doBuildOrbitPath(orbitantsToOrbits, orbitant, orbitPath))
}

func main() {
	// bytes, _ := ioutil.ReadFile("./day6/part2/example_input")
	bytes, _ := ioutil.ReadFile("./day6/part2/input")

	orbitsRelation := strings.Split(string(bytes), "\n")
	orbitantsToOrbits := buildOrbitansToOrbits(orbitsRelation)

	sanOrbitPath := buildOrbitPath(orbitantsToOrbits, "YOU")
	youOrbitPath := buildOrbitPath(orbitantsToOrbits, "SAN")

	var youCommonOrbitPath []string
	var sanCommonOrbitPath []string
	for pos, _ := range sanOrbitPath {
		if youOrbitPath[pos] != sanOrbitPath[pos] {
			youCommonOrbitPath = youOrbitPath[pos:(len(youOrbitPath) - 1)]
			sanCommonOrbitPath = sanOrbitPath[pos:(len(sanOrbitPath) - 1)]
			break
		}
	}

	fmt.Printf("Part2 Solution: %d\n", len(youCommonOrbitPath)+len(sanCommonOrbitPath))
}
