package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

const (
	DirListViewName         = "dirListView"
	FileListViewName        = "fileListView"
	SearchBarInputViewName  = "pathInput"
	SearchBarButtonViewName = "btn"
)

type Point struct {
	X0 int
	Y0 int
	X1 int
	Y1 int
}

func (g *GUI) layout(gui *gocui.Gui) error {
	// dir list layout
	err := g.dirListViewLayout()
	if err != nil {
		return err
	}

	// file list layout
	err = g.fileListViewLayout()
	if err != nil {
		return err
	}

	// searchBtnOnClick input layout
	err = g.searchBarInputViewLayout()
	if err != nil {
		return err
	}

	// searchBtnOnClick button layout
	err = g.searchBarBtnViewLayout()
	if err != nil {
		return err
	}

	return nil
}

// 目錄選單列表layout
func (g *GUI) dirListViewLayout() error {
	if v, err := g.Gui.SetView(DirListViewName, g.DirListViewPoints.X0, g.DirListViewPoints.Y0, g.DirListViewPoints.X1, g.DirListViewPoints.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		err = g.setDirListViewContent(v)
		if err != gocui.ErrUnknownView {
			return err
		}

		if _, err = g.focus(DirListViewName); err != nil {
			return err
		}

	}

	return nil
}

// 設置目錄選單列表layout內容
func (g *GUI) setDirListViewContent(v *gocui.View) error {
	v.Clear()

	for _, entry := range dirs {
		_, _ = fmt.Fprintln(v, entry.Name())
	}

	return nil
}

// 檔案選單列表layout
func (g *GUI) fileListViewLayout() error {
	if v, err := g.Gui.SetView(FileListViewName, g.FileListViewPoints.X0, g.FileListViewPoints.Y0, g.FileListViewPoints.X1, g.FileListViewPoints.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		err = g.setFileListViewContent(v)
		if err != gocui.ErrUnknownView {
			return err
		}

		if _, err = g.focus(DirListViewName); err != nil {
			return err
		}
	}

	return nil
}

// 設置檔案選單列表layout內容
func (g *GUI) setFileListViewContent(v *gocui.View) error {
	v.Clear()

	for _, entry := range dirs {
		_, _ = fmt.Fprintln(v, entry.Name())
	}

	return nil
}

// EditText of file path
func (g *GUI) searchBarInputViewLayout() error {
	if v, err := g.Gui.SetView(SearchBarInputViewName, g.SearchBarInputViewPoints.X0, g.SearchBarInputViewPoints.Y0, g.SearchBarInputViewPoints.X1, g.SearchBarInputViewPoints.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.Wrap = true
	}
	return nil
}

// Button of Search
func (g *GUI) searchBarBtnViewLayout() error {
	if v, err := g.Gui.SetView(SearchBarButtonViewName, g.SearchBarBtnViewPoints.X0, g.SearchBarBtnViewPoints.Y0, g.SearchBarBtnViewPoints.X1, g.SearchBarBtnViewPoints.Y1); err != nil {
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
