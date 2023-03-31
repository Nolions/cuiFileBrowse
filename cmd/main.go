package main

import (
	"github.com/Nolions/cuiFileBrowser/internal/gui"
	"log"
)

func main() {
	g := gui.Create()
	defer g.Gui.Close()

	maxX, maxY := g.Size()
	menu := gui.Point{
		X0: 0,
		Y0: 0,
		X1: int(0.2 * float32(maxX)),
		Y1: maxY - 5,
	}

	content := gui.Point{
		X0: int(0.2 * float32(maxX)),
		Y0: 0,
		X1: int(float32(maxX) * 0.8),
		Y1: maxY - 5,
	}

	editText := gui.Point{
		X0: 0,
		Y0: maxY - 5,
		X1: int(0.7 * float32(maxX)),
		Y1: maxY - 3,
	}

	button := gui.Point{
		X0: int(0.7 * float32(maxX)),
		Y0: maxY - 5,
		X1: int(float32(maxX) * 0.8),
		Y1: maxY - 3,
	}

	err := g.SetLayout(menu, content, editText, button).
		BindingKeys().
		EnableCursor(true).
		EnableMouseClick(true).
		Show()
	if err != nil {
		log.Fatalf("error:%s\n", err.Error())
	}
}
