package main

import (
	"github.com/Nolions/cuiFileBrowser/internal/gui"
	"log"
)

func main() {
	g := gui.Create()
	defer g.Gui.Close()

	err := g.SetLayout().BindingKeys().EnableCursor(true).EnableMouseClick(true).Show()
	if err != nil {
		log.Fatalf("error:%s\n", err.Error())
	}
}
