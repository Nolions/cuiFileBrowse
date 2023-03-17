package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	menuLayoutWidth := int(0.2 * float32(maxX))
	menuLayoutHeight := maxY - 5
	contentLayoutWidth := int(float32(maxX) * 0.8)

	// menu layout
	err := menuLayout(g, 0, 0, menuLayoutWidth, menuLayoutHeight)
	if err != nil {
		return err
	}

	err = contentLayout(g, menuLayoutWidth, 0, contentLayoutWidth, menuLayoutHeight)
	if err != nil {
		return err
	}

	return nil
}

// menu layout
func menuLayout(g *gocui.Gui, startXPoint, startYPoint, endXPoint, endYPoint int) error {
	if v, err := g.SetView("menu", 0, 0, endXPoint, endYPoint); err != nil {
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
func contentLayout(g *gocui.Gui, startXPoint, startYPoint, endXPoint, endYPoint int) error {
	if _, err := g.SetView("content", startXPoint, startYPoint, endXPoint, endYPoint); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}
