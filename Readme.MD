# Quick info
 * `calculation.go` contains the logic to find out the new position of the rover. No parsing/formatting here
 * `formatting.go` contains the logic to parse/format strings from/to NASA
 * `structures.go` defines some types for the position of the rover, the movement possibilities, etc.
 * `main.go` checkouts the logic with a basic example
 * Issue `go run ./main` to try it!

# Assumptions
* The rover should not go out of the plateau
* The strings from nasa contain the expected characters without spaces between each char.

# Others
* In real life this would be most likely in hardware. I'd love to try it in an FPGA!
* Upon wrong string input, we panic. In real life we should have some back-up ideas...
