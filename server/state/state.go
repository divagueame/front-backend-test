package state

import (
	"fmt"
	// "strconv"
)

type Position struct {
	x int
	y int
}

type Direction int

const (
	DirectionUp = iota
	DirectionUpperRight
	DirectionRight
	DirectionLowerRight
	DirectionDown
	DirectionLowerLeft
	DirectionLeft
	DirectionUpperLeft
)

var directionName = map[Direction]string{
	DirectionUp:         ":arrow_up:",
	DirectionUpperRight: ":arrow_upper_right:",
	DirectionRight:      ":arrow_right:",
	DirectionLowerRight: ":arrow_lower_right:",
	DirectionDown:       ":arrow_down:",
	DirectionLowerLeft:  ":arrow_lower_left:",
	DirectionLeft:       ":arrow_left:",
	DirectionUpperLeft:  ":arrow_upper_left:",
}

type GridState struct {
	position  Position
	grid      [30][30]int
	direction Direction
	brush     string
}

var state *GridState

func Initialize() {
	state = &GridState{
		position: Position{
			x: 15,
			y: 15,
		},
		direction: DirectionUp,
		brush:     "hover",
	}
}
func parseSquare(num int) string {
	if num == 0 {
		return " "
	} else if num == 1 {
		return "X"
	} else if num == 2 {
		return "P"
	} else {
		fmt.Print("ERror parsing Square", num)
		return " "
	}
}
func GetCanvas() string {
	first_line := "╔══════════════════════════════╗"
	last_line := "╚══════════════════════════════╝"

	line := ""
	line += first_line
	line += "\n"

	for _, row := range state.grid {
		line += "║"
		for _, square := range row {
			line += parseSquare(square)
		}
		line += "║"
		line += "\n"

	}

	line += last_line
	line += "\n\n"

	// fmt.Print("", line)
	return line
}

func ChangeBrush(newBrush string) {
	if newBrush == "hover" {
		fmt.Println("Changing to hoverj")
	} else if newBrush == "right" {
		fmt.Println("Changing to draw")
	} else if newBrush == "eraser" {
		fmt.Println("Changing to draw")
	} else {
		fmt.Println("---> Wrong brush", newBrush)
	}
}

func QuitConnection() {
	fmt.Println("Quitting connection")
}

func ClearCanvas() {
	fmt.Println("Clearing canvas")
}

func ChangeDirection(newDirection string, steps int) {
	if newDirection == "left" {
		fmt.Println("Changing to left")
	} else if newDirection == "right" {
		fmt.Println("Changing to right")
	} else {
		fmt.Println("---> Wrong Direction", newDirection)
	}
}

func MoveCursor(steps int) {
	fmt.Println("mvoing cursor", steps)
}

func PrintState() {
	PrintCoord()
	fmt.Println("Grid", state.grid)
	fmt.Println("Direction", state.direction)
}

func PrintCoord() {
	fmt.Println("Priting pos...", state.position.x)
	fmt.Println("Priting pos...", state.position.y)
}
func updateCurrentPosition(x int, y int) {
	state.position.x = x
	state.position.y = y
}

func GetPosition() Position {
	return state.position
}
func GetPositionStr() string {
	return fmt.Sprintf("(%d,%d)", state.position.x, state.position.y)
}

// func GetPositionStr(pos Position) string {
// parsedX := strconv.Itoa(pos.x)
// parsedY := strconv.Itoa(pos.y)
// return fmt.Sprintf("(%s,%s)", parsedX, parsedY)
// return fmt.Sprintf("(%d,%d)", pos.x, pos.y)
// }

func clearGrid() {
	state.grid = [30][30]int{}
}
