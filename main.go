package main

import (
	// "fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"os"
)

var (
	winTitle string = "Моделирование жидкости"

	particles [5]int
	speed     int = 30
	partNum   int = 1000
)

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	// var points []sdl.Point
	// var rect sdl.Rect
	// var rects []sdl.Rect

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1024, 720, 0)
	if err != nil {
		return 1
	}
	// Уничтожить окно при выходе из функции
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return 2
	}
	defer renderer.Destroy()
	// Иначе будет отображенно, что под окном
	renderer.Clear()

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoint(150, 300)

	for i := 0; i < 800; i++ {
		var y float64 = float64(i) * math.Sin(float64(i))
		renderer.DrawPoint(i, int(y))
	}

	// renderer.SetDrawColor(0, 0, 255, 255)
	// renderer.DrawLine(0, 0, 200, 200)

	// points = []sdl.Point{{0, 0}, {100, 300}, {100, 300}, {200, 0}}
	// renderer.SetDrawColor(255, 255, 0, 255)
	// renderer.DrawLines(points)

	// rect = sdl.Rect{300, 0, 200, 200}
	// renderer.SetDrawColor(255, 0, 0, 255)
	// renderer.DrawRect(&rect)

	// rects = []sdl.Rect{{400, 400, 100, 100}, {550, 350, 200, 200}}
	// renderer.SetDrawColor(0, 255, 255, 255)
	// renderer.DrawRects(rects)

	// rect = sdl.Rect{250, 250, 200, 200}
	// renderer.SetDrawColor(0, 255, 0, 255)
	// renderer.FillRect(&rect)

	// rects = []sdl.Rect{{500, 300, 100, 100}, {200, 300, 200, 200}}
	// renderer.SetDrawColor(255, 0, 255, 255)
	// renderer.FillRects(rects)

	renderer.Present()

	sdl.Delay(2000)

	return 0
}

func main() {
	os.Exit(run())
}
