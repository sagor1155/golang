package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

// animate bouncing ball
// algorithm:
// create the board [][]bool
// screen.clear()
// inside draw loop:
// calculate and update the next ball position
// draw the board into a []rune buffer
// screen.MoveTopLeft()
// print the []rune buffer: string(buffer)
// time.Sleep(..)

func main() {
	// create and draw the board
	const (
		width     = 50
		height    = 10
		cellEmpty = ' '
		cellBall  = '⚾'
		maxFrames = 400
	)

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)

	board := make([][]bool, width)
	for row := range board {
		board[row] = make([]bool, height)
	}

	for i := 0; i < maxFrames; i++ {
		screen.Clear()
		// increment balls position
		px += vx
		py += vy

		// bounce the ball
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// remove the previous ball
		for y := range board[0] { // vertical position
			for x := range board { // horizontal position
				board[x][y] = false
			}
		}

		// set the ball position
		board[px][py] = true

		buf := make([]rune, 0, width*height)
		buf = buf[:0]

		// draw board
		for y := range board[0] { // vertical position
			for x := range board { // horizontal position
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		fmt.Println(string(buf))
		time.Sleep(time.Second / 20)
	}
}
