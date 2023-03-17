package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type LocationPoint struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
}

func (g *GUI) layout(gui *gocui.Gui) error {
	// menu layout
	err := g.menuLayout()
	if err != nil {
		return err
	}

	err = g.contentLayout()
	if err != nil {
		return err
	}

	return nil
}

// menu layout
func (g *GUI) menuLayout() error {
	if v, err := g.Gui.SetView("menu", g.MenuLocation.X0, g.MenuLocation.Y0, g.MenuLocation.X1, g.MenuLocation.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		for _, item := range []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"} {
			_, _ = fmt.Fprintln(v, item)
		}
	}

	return nil
}

// content Layout
func (g *GUI) contentLayout() error {
	if _, err := g.Gui.SetView("content", g.ContentLocation.X0, g.ContentLocation.Y0, g.ContentLocation.X1, g.ContentLocation.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}
