package nasa

// CardinalPoint is a cardinal point, i.e. North, East, South or West. Let's have it typed, underlying data type is a bytes
type CardinalPoint byte

const ( // definition of the cardinal points
	north = CardinalPoint(byte('N'))
	south = CardinalPoint(byte('S'))
	east  = CardinalPoint(byte('E'))
	west  = CardinalPoint(byte('W'))
)

var cardinalPoints = map[CardinalPoint]struct{}{north: {}, south: {}, east: {}, west: {}}

// RoverPosition Is the position of a rover unit including its facing, i.e. to a cardinal point
type roverPosition struct {
	position coordinates
	facing   CardinalPoint
}

// movement Represents a movement instruction from nasa. Either move, or turn left/right
type movement byte

const (
	move      = movement(byte('M'))
	turnLeft  = movement(byte('L'))
	turnRight = movement(byte('R'))
)

var possibleMovements = map[movement]struct{}{move: {}, turnLeft: {}, turnRight: {}}

// coordinates Is just a generic two-point coordinate, used for the location of the rover but also to represent the plateau size, i.e. x=width, y=height
type coordinates struct {
	x uint64
	y uint64
}
