package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

type GUI struct {
	Gui            gocui.Gui
	MenuPoints     Point
	ContentPoints  Point
	EditTextPoints Point
	BtnPoints      Point
}

var (
	viewArr = []string{"menu", "content", "pathInput"}
	active  = 0
)

// Create
// New GUI struct Object
func Create() *GUI {
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
	g.MenuPoints = ml
	g.ContentPoints = cl
	g.EditTextPoints = pil
	g.BtnPoints = bl

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
