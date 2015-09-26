package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

func TestRender(t *testing.T) {
	window, err := sdl.CreateWindow("Test123", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		t.Error("Окно не созданно: ", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		t.Error("Рендер не создан: ", err)
	}

	renderer.Clear()

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoint(150, 300)

	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLine(0, 0, 200, 200)

	renderer.Present()
}
