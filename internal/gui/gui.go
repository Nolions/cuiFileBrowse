package gui

import (
	"github.com/jroimartin/gocui"
	"log"
)

type GUI struct {
	Gui gocui.Gui
}

// Create
// New GUI struct Object
func Create() *GUI {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	return &GUI{
		Gui: *g,
	}
}

// SetLayout
// Layout配置
func (g *GUI) SetLayout() *GUI {
	g.Gui.SetManagerFunc(layout)
	return g
}

// BindingKeys
// Hotkey binder
func (g *GUI) BindingKeys() *GUI {
	if err := g.Gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatalf("BindingKeys(), KeyCtrlC binding error:%s\n", err.Error())
	}

	for _, binder := range keyBinders() {
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
