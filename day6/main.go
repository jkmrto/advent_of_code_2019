package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func keyExist(hash map[string][]string, key string) bool {
	_, ok := hash[key]
	return ok
}

func buildObjectToOrbits(orbitsRelation []string) map[string][]string {
	// Object to list of orbitants
	objToOrbs := make(map[string][]string)
	for _, orbitRelation := range orbitsRelation {
		splitObitRelation := strings.Split(orbitRelation, ")")
		object := splitObitRelation[0]
		orbitant := splitObitRelation[1]

		if keyExist(objToOrbs, object) {
			objToOrbs[object] = append(objToOrbs[object], orbitant)
		} else {
			objToOrbs[object] = []string{orbitant}
		}

	}
	return objToOrbs
}

func calcOrbits(objToOrbits map[string][]string, key string, indirectOrbits int) int {
	nodeOrbits := indirectOrbits + 1
	acc := indirectOrbits
	if keyExist(objToOrbits, key) {
		for _, orbitObject := range objToOrbits[key] {
			acc += calcOrbits(objToOrbits, orbitObject, nodeOrbits)
		}
	}

	return acc
}

func main() {
	bytes, _ := ioutil.ReadFile("./day6/input")
	orbitsRelation := strings.Split(string(bytes), "\n")
	objToOrbs := buildObjectToOrbits(orbitsRelation)
	value := calcOrbits(objToOrbs, "COM", 0)
	fmt.Printf("The solution is: %d\n", value)
}
