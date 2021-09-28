package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Cell struct {
	armed 			bool
	clicked 		bool
	tile 			string
	x 				int
	y 				int
}

func NewCell(x, y int) Cell {
	c := Cell {
		x: x,
		y: y,
	}

	c.armed = false
	c.clicked = false

	return c
}


type Game struct {
	height 			int
	field 			[][]Cell
	id 				uuid.UUID
	width 			int
}

func NewGame(height, width int) *Game {
	g := Game{
		height: height,
		width: width,
	}

	g.field = make([][]Cell, width)
	g.id = uuid.New()

	return &g
}

func main() {
	fmt.Println("Starting minesweeper (go) ...")
	fmt.Println()

	game := NewGame(4, 6)

	for x := 0; x < game.width; x++ {
		game.field[x] = make([]Cell, game.height)
		for y := 0; y < game.height; y++ {
			game.field[x][y] = NewCell(x, y)
		}
	}
}
