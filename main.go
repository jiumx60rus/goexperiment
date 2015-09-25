package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"os"
)

const x = math.Pi / 180

func Rad(d float64) float64 { return d * x }
func Deg(r float64) float64 { return r / x }

var (
	winTitle      string  = "SIN"
	width, height int     = 1024, 720
	w             float64 = float64(width) / 2
	h             float64 = float64(height) / 2

	fps           uint32 = 60         //Кол-во кадров
	fpsMill       uint32 = 1000 / fps //В один кадр в мс
	currentSpeed  uint32 = 0          //Время ожиданиия после выполнения цикла примера
	currentTime   uint32 = 0          //Вреся перед началом запуска игрового цикла
	countedFrames int    = 0          //Кадров в сек

	particles [5]int
	speed     int = 30
	partNum   int = 1000

	window   *sdl.Window
	renderer *sdl.Renderer
	event    sdl.Event
)

func main() {
	os.Exit(run())
}

func run() int {
	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, width, height, sdl.WINDOW_SHOWN)
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

	//Главный цикл
	for {
		currentTime = sdl.GetTicks()

		//Step
		step()
		window.SetTitle(winTitle + " Кадр: " + fmt.Sprint(countedFrames))
		countedFrames++

		currentSpeed = sdl.GetTicks() - currentTime
		if fpsMill > currentSpeed {
			// sdl.Delay(1000)
			sdl.Delay(fpsMill - currentSpeed)
		}
	}

	return 0
}

func step() {

	// Иначе будет отображенно, что под окном
	renderer.Clear()

	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.MouseButtonEvent:
			os.Exit(1)
		}
	}

	renderer.SetDrawColor(90, 140, 70, 255)

	for t := 0; t <= width; t++ {
		var x int = t
		var y int = int(100 * math.Sin(float64(x*100)+float64(countedFrames/4)))
		// fmt.Fprintf(os.Stderr, fmt.Sprint(float64(countedFrames/3))+"\n")
		renderer.DrawPoint(x, y+360)
	}

	// Цвет фона
	renderer.SetDrawColor(51, 51, 51, 255)
	renderer.Present()
}
