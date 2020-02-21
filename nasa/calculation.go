package nasa

import "fmt"

// CalculatePosition Computes the final position of the rover based on the encoded strings providing
// the initial position, plateau size and movements.
// The output is also encoded as string.
func CalculatePosition(plateauSize string, initialPosition string, movements string) string {
	parsedPlateauSize, errorParsingPlateauSize := parseCoordinates(plateauSize)
	if errorParsingPlateauSize != nil {
		panic("We could not parse the plateau size! Aborting Mission!")
	}

	parsedInitialPosition, errParsingRoverPosition := parseRoverPosition(initialPosition, parsedPlateauSize)
	if errParsingRoverPosition != nil {
		panic("We could not parse the rover initial position. Aborting Mission!")
	}
	parsedMovements, errParsingMovements := parseMovements(movements)
	if errParsingMovements != nil {
		panic("We could not decode the movements string... Aborting Mission!")
	}

	currentPosition := parsedInitialPosition

	for _, currentMovement := range *parsedMovements {
		switch currentMovement {
		case turnRight:
			currentPosition.facing = rightTurns[currentPosition.facing]
		case turnLeft:
			currentPosition.facing = leftTurns[currentPosition.facing]
		case move:
			errorMovingRover := moveRover(&currentPosition.position, currentPosition.facing, parsedPlateauSize)
			if errorMovingRover != nil {
				panic(fmt.Sprintf("We cannot move the rover because %v .... :(", errorMovingRover))
			}
		}
	}

	return formatPosition(currentPosition)
}

// moveRover Assumes the rover should not try to go out of the plateau. Otherwise it'd die.
func moveRover(currentPosition *coordinates, facing CardinalPoint, plateauSize *coordinates) error {
	switch facing {
	case north:
		if currentPosition.y+1 > plateauSize.y {
			return fmt.Errorf("Can't move. Rover would go out of the plateau")
		}
		currentPosition.y++
	case east:
		if currentPosition.x+1 > plateauSize.x {
			return fmt.Errorf("Can't move. Rover would go out of the plateau")
		}
		currentPosition.x++
	case south:
		if currentPosition.y-1 < 0 {
			return fmt.Errorf("Can't move. Rover would go out of the plateau")
		}
		currentPosition.y--
	case west:
		if currentPosition.x-1 < 0 {
			return fmt.Errorf("Can't move. Rover would go out of the plateau")
		}
		currentPosition.x--
	}
	return nil
}

var rightTurns = map[CardinalPoint]CardinalPoint{
	north: east,
	east:  south,
	south: west,
	west:  north,
}

var leftTurns = map[CardinalPoint]CardinalPoint{
	north: west,
	east:  north,
	south: east,
	west:  south,
}
