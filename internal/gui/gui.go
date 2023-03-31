package gui

import (
	"github.com/Nolions/cuiFileBrowser/internal/fileSystem"
	"github.com/jroimartin/gocui"
	"log"
	"os"
)

type GUI struct {
	Gui                      gocui.Gui
	DirListViewPoints        Point
	FileListViewPoints       Point
	SearchBarInputViewPoints Point
	SearchBarBtnViewPoints   Point
}

var (
	viewArr     = []string{DirListViewName, FileListViewName, SearchBarInputViewName}
	active      = 0
	dirs, files []os.DirEntry
)

// Create
// New GUI struct Object
func Create() *GUI {
	// TODO
	dirs, files = fileSystem.GetFiles(".")

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	return &GUI{
		Gui: *g,
	}
}

func (g *GUI) Size() (int, int) {
	return g.Gui.Size()
}

// SetLayout
// Layout配置
func (g *GUI) SetLayout(ml, cl, pil, bl Point) *GUI {
	g.DirListViewPoints = ml
	g.FileListViewPoints = cl
	g.SearchBarInputViewPoints = pil
	g.SearchBarBtnViewPoints = bl

	g.Gui.Highlight = true
	g.Gui.Cursor = true
	g.Gui.Highlight = true
	g.Gui.Cursor = true
	g.Gui.SelFgColor = gocui.ColorGreen

	g.Gui.SetManagerFunc(g.layout)

	return g
}

// BindingKeys
// Hotkey binder
func (g *GUI) BindingKeys() *GUI {
	if err := g.Gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, g.quit); err != nil {
		log.Fatalf("BindingKeys(), KeyCtrlC binding error:%s\n", err.Error())
	}

	for _, binder := range g.keyBinders() {
		if err := g.Gui.SetKeybinding(binder.ViewName, binder.Key, gocui.ModNone, binder.Action); err != nil {
			log.Fatalf("BindingKeys(), MouseLeft binding error:%s\n", err.Error())
		}
	}

	return g
}

// EnableMouseClick
// 啟用滑鼠點擊
func (g *GUI) EnableMouseClick(statue bool) *GUI {
	g.Gui.Mouse = statue
	return g
}

// EnableCursor
// 啟用光標
func (g *GUI) EnableCursor(statue bool) *GUI {
	g.Gui.Cursor = statue
	return g
}

// Show
// 顯示
func (g *GUI) Show() error {
	err := g.Gui.MainLoop()

	return err
}
