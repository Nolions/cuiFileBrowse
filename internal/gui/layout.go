package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

const (
	MenuView            = "menu"
	ContentView         = "content"
	SearchBarInputView  = "pathInput"
	SearchBarButtonView = "btn"
)

type Point struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
}

func (g *GUI) layout(gui *gocui.Gui) error {
	// dir list layout
	err := g.menuLayout()
	if err != nil {
		return err
	}

	// file list layout
	err = g.contentLayout()
	if err != nil {
		return err
	}

	// settingPath input layout
	err = g.editTextLayout()
	if err != nil {
		return err
	}

	// settingPath button layout
	err = g.btnLayout()
	if err != nil {
		return err
	}

	return nil
}

// menu layout
func (g *GUI) menuLayout() error {
	if v, err := g.Gui.SetView(MenuView, g.MenuPoints.X0, g.MenuPoints.Y0, g.MenuPoints.X1, g.MenuPoints.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		for _, item := range []string{
			"Item 1", "Item 2", "Item 3", "Item 4", "Item 5", "Item 6", "Item 7", "Item 8", "Item 9", "Item 10"} {
			_, _ = fmt.Fprintln(v, item)
		}

		if _, err = g.focus("menu"); err != nil {
			return err
		}
	}

	return nil
}

// content Layout
func (g *GUI) contentLayout() error {
	if v, err := g.Gui.SetView(ContentView, g.ContentPoints.X0, g.ContentPoints.Y0, g.ContentPoints.X1, g.ContentPoints.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Wrap = true
		v.Autoscroll = true
	}

	return nil
}

// EditText of file path
func (g *GUI) editTextLayout() error {
	if v, err := g.Gui.SetView(SearchBarInputView, g.EditTextPoints.X0, g.EditTextPoints.Y0, g.EditTextPoints.X1, g.EditTextPoints.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.Wrap = true
	}
	return nil
}

// Button of Search
func (g *GUI) btnLayout() error {
	if v, err := g.Gui.SetView(SearchBarButtonView, g.BtnPoints.X0, g.BtnPoints.Y0, g.BtnPoints.X1, g.BtnPoints.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		_, _ = fmt.Fprintln(v, "Search")
	}
	return nil
}

func (g *GUI) focus(name string) (*gocui.View, error) {
	if _, err := g.Gui.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.Gui.SetViewOnTop(name)
}
