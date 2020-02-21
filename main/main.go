package main

import (
	"fmt"
	"nasa/nasa"
)

func main() {
	plateauSize := "55"
	initialPosition := "12N"
	movements := "LMLMLMLMM"

	finalPosition1 := nasa.CalculatePosition(plateauSize, initialPosition, movements)
	fmt.Printf("Final position 1:\n%v\n\n", finalPosition1)

	plateauSize2 := "55"
	initialPosition2 := "33E"
	movements2 := "MMRMMRMRRM"

	finalPosition2 := nasa.CalculatePosition(plateauSize2, initialPosition2, movements2)
	fmt.Printf("Final position 2:\n%v\n", finalPosition2)
}
