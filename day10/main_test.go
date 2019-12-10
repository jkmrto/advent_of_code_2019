package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinatesToDegrees(t *testing.T) {

	// Typical case 1 to 1
	assert.Equal(t, 45.0, CoordinatesToDegrees(1, -1), "FUCK")
	assert.Equal(t, 135.0, CoordinatesToDegrees(1, 1), "FUCK")
	assert.Equal(t, 225.0, CoordinatesToDegrees(-1, 1), "FUCK")
	assert.Equal(t, 315.0, CoordinatesToDegrees(-1, -1), "FUCK")

	// // Typical case 1 to 3
	// assert.Equal(t, 18.43494882292201, CoordinatesToDegrees(1, 3), "FUCK")
	// assert.Equal(t, 161.56505117707798, CoordinatesToDegrees(1, -3), "FUCK")
	// assert.Equal(t, 198.43494882292202, CoordinatesToDegrees(-1, -3), "FUCK")
	// assert.Equal(t, 341.565051177078, CoordinatesToDegrees(-1, 3), "FUCK")

	// Special cases
	assert.Equal(t, 0.0, CoordinatesToDegrees(0, -1), "FUCK")
	assert.Equal(t, 90.0, CoordinatesToDegrees(1, 0), "FUCK")
	assert.Equal(t, 180.0, CoordinatesToDegrees(0, 1), "FUCK")
	assert.Equal(t, 270.0, CoordinatesToDegrees(-1, 0), "FUCK")

}
