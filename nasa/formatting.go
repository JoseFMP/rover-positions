package nasa

// This is the most dangerous file.
// Parsing stuff strings that are coming from outside the system...
// In real life we should have some recovery policies upon wrong input.
// We should also try to interpret what the input meant, even if it was not correctly formatted.

import (
	"fmt"
	"strconv"
)

func formatPosition(roverPosition *roverPosition) string {
	return fmt.Sprintf("%v %v %c", roverPosition.position.x, roverPosition.position.y, roverPosition.facing)
}

func parseRoverPosition(positionAsString string, plateauSize *coordinates) (*roverPosition, error) {
	if len(positionAsString) != 3 {
		return nil, fmt.Errorf("Position should have three characters")
	}
	position, errorParsingPosition := parseCoordinates(positionAsString[0:2])
	if errorParsingPosition != nil {
		return nil, errorParsingPosition
	}

	if position.x < 0 || position.y < 0 || position.x > plateauSize.x || position.y > plateauSize.y {
		return nil, fmt.Errorf("It seems position of rover is incorrect. It is outside the plateau ! ")
	}

	facing := CardinalPoint(positionAsString[2])
	_, cardinalPointExists := cardinalPoints[facing]
	if !cardinalPointExists {
		return nil, fmt.Errorf("Cardinal point does not exist")
	}
	return &roverPosition{coordinates{position.x, position.y}, facing}, nil
}

func parseCoordinates(plateauSizeAsString string) (*coordinates, error) {
	if len(plateauSizeAsString) != 2 {
		return nil, fmt.Errorf("Position should have three characters")
	}
	plateauSizeRunes := []rune(plateauSizeAsString)

	xAsString := string(plateauSizeRunes[0])
	yAsString := string(plateauSizeRunes[1])

	x, errParsingCoordinateX := strconv.ParseUint(xAsString, 10, 64)
	y, errParsingCoordinateY := strconv.ParseUint(yAsString, 10, 64)

	if errParsingCoordinateX != nil || errParsingCoordinateY != nil {
		return nil, fmt.Errorf("Error parsing coordinates. [%v, %v]", xAsString, yAsString)
	}

	return &coordinates{x, y}, nil
}

func parseMovements(movementsAsString string) (*[]movement, error) {
	movements := make([]movement, 0)
	for _, movementRune := range movementsAsString {
		candidateMovement := movement(byte(movementRune))
		_, movementCharValid := possibleMovements[candidateMovement]
		if !movementCharValid {
			return nil, fmt.Errorf("Movement string seems malformed. What is \"\" supposed to mean?", movementCharValid)
		}
		movements = append(movements, candidateMovement)
	}
	return &movements, nil
}
