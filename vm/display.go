/* Very simple, cursor based display.

The screen is rectangular.
*/

package vm

import (
	"fmt"
	"runtime"
	"time"
)

const (
	RESET_CURSOR      = 0x00
	MOVE_CURSOR_RIGHT = 0x10
	MOVE_CURSOR_LEFT  = 0x11
	MOVE_CURSOR_DOWN  = 0x12
	MOVE_CURSOR_UP    = 0x13
	FRAMES_PER_SECOND = 5
)

type Surface struct {
	cols int
	rows int
}

type Coordinate struct {
	x int
	y int
}

type Display struct {
	name       string
	size       Surface
	resolution Surface
	cursor     Coordinate
	pixels     [][]*pixel
}

func (display Display) Read() byte {
	return 0
}

func (display Display) Write(value byte) {
	switch value {
	case RESET_CURSOR:
		fmt.Println("Resetting cursor")
		display.cursor.x = 0
		display.cursor.y = 0
	case MOVE_CURSOR_RIGHT:
		fmt.Println("Move cursor right")
		if display.cursor.x < display.size.cols-1 {
			display.cursor.x++
		}
	case MOVE_CURSOR_LEFT:
		fmt.Println("Move cursor left")
		if display.cursor.x > 0 {
			display.cursor.x--
		}
	case MOVE_CURSOR_DOWN:
		fmt.Println("Move cursor down")
		if display.cursor.y < display.size.rows-1 {
			display.cursor.y++
		}
	case MOVE_CURSOR_UP:
		fmt.Println("Move cursor up")
		if display.cursor.y > 0 {
			display.cursor.y--
		}
	default:
		fmt.Printf("Invalid operation: %X\n", value)
	}
}

func NewDisplay(cols int, rows int) *Display {
	runtime.LockOSThread()

	display := Display{
		name:       "GrogVM Display",
		resolution: Surface{cols: 320, rows: 200},
		size:       Surface{cols: cols, rows: rows},
		cursor:     Coordinate{x: 0, y: 0},
	}
	go display.Init()
	return &display
}

func (d *Display) Init() {

	window := initGlfw(d.resolution, d.name)
	//defer glfw.Terminate()

	program := initOpenGL()

	d.pixels = makePixels(d.size)

	for !window.ShouldClose() {
		renderStartTime := time.Now()
		draw(d.pixels, window, program)
		time.Sleep(delay(renderStartTime))
	}
}

func makePixels(surface Surface) [][]*pixel {
	pixels := make([][]*pixel, surface.rows, surface.cols)
	for row := 0; row < surface.rows; row++ {
		for col := 0; col < surface.cols; col++ {
			pixels[row] = append(pixels[row], newPixel(Coordinate{x: col, y: row}, surface))
		}
	}
	return pixels
}

func delay(renderStartTime time.Time) time.Duration {
	return time.Second/time.Duration(FRAMES_PER_SECOND) - time.Since(renderStartTime)
}
