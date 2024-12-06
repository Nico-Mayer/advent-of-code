package day06

import (
	"fmt"

	"github.com/nico-mayer/aoc-2024/utils"
)

type Vector struct {
	X int
	Y int
}

type Guard struct {
	Position  Vector
	Direction Vector
}

type VisitedPos struct {
	Pos       Vector
	Direction Vector
}

var (
	UP    = Vector{X: 0, Y: -1}
	DOWN  = Vector{X: 0, Y: 1}
	LEFT  = Vector{X: -1, Y: 0}
	RIGHT = Vector{X: 1, Y: 0}
)

func Run(test bool, day int, part int) {
	lines := utils.LoadData(day, test)

	var guard Guard
	var startPos Vector
	for row, line := range lines {
		for col := range line {
			if lines[row][col] == byte('^') {
				startPos.X = col
				startPos.Y = row
				break
			}
		}
	}

	guard.Reset(startPos)

	// Part 01
	validPoints := utils.NewSet[Vector]()
	for {
		nextPos := guard.GetNextPos()
		if !inBounds(lines, nextPos) {
			break
		}
		objectInFront := lines[nextPos.Y][nextPos.X]
		if objectInFront == byte('#') {
			guard.Turn()
		} else {
			guard.SetPosition(nextPos)
			validPoints.Add(guard.Position)
		}
	}

	// Part 02
	loopCount := 0
	validPoints.ForEach(func(point Vector) {
		guard.Reset(startPos)
		isLoop := false
		visited := utils.NewSet[VisitedPos]()
		for {
			nextPos := guard.GetNextPos()
			if !inBounds(lines, nextPos) {
				break
			}

			objectInFront := lines[nextPos.Y][nextPos.X]
			if objectInFront == byte('#') || nextPos == point {
				guard.Turn()
			} else {
				guard.Position = nextPos
				vis := VisitedPos{
					Pos:       nextPos,
					Direction: guard.Direction,
				}
				if visited.Contains(vis) {
					isLoop = true
				} else {
					visited.Add(vis)
				}
			}
			if isLoop {
				loopCount++
				break
			}
		}
	})
	fmt.Println("Day06:")
	switch part {
	case 1:
		fmt.Printf("P1: %d\n", validPoints.Size()+1)
	case 2:
		fmt.Printf("P2: %d\n", loopCount)
	default:
		fmt.Printf("P1: %d\n", validPoints.Size()+1)
		fmt.Printf("P2: %d\n", loopCount)
	}

}

func inBounds(m []string, point Vector) bool {
	maxRows := len(m) - 1
	maxCols := len(m[0]) - 1
	if point.X > maxCols || point.X < 0 {
		return false
	}
	if point.Y > maxRows || point.Y < 0 {
		return false
	}
	return true
}

func (g *Guard) Turn() {
	switch g.Direction {
	case UP:
		g.Direction = RIGHT
	case DOWN:
		g.Direction = LEFT
	case LEFT:
		g.Direction = UP
	case RIGHT:
		g.Direction = DOWN
	}
}

func (g *Guard) SetPosition(pos Vector) {
	g.Position = pos
}

func (g *Guard) Reset(startPos Vector) {
	g.Position = startPos
	g.Direction = UP
}

func (g *Guard) GetNextPos() Vector {
	return Vector{
		X: g.Direction.X + g.Position.X,
		Y: g.Direction.Y + g.Position.Y,
	}
}
