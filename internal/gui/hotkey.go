package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

type KeyBinder struct {
	Key      gocui.Key
	Action   func(g *gocui.Gui, v *gocui.View) error
	ViewName string
}

// hotkey binder action
func (g *GUI) keyBinders() []KeyBinder {
	var binders []KeyBinder
	binders = append(binders, KeyBinder{
		Key:      gocui.MouseLeft,
		ViewName: "menu",
		Action:   g.setContent,
	})

	binders = append(binders, KeyBinder{
		Key:      gocui.KeyTab,
		ViewName: "",
		Action:   g.nextView,
	})

	return binders
}

// 關閉程式
func (g *GUI) quit(gui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// 設置content layout內容
func (g *GUI) setContent(gui *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	s, err := v.Line(cy)
	if err != nil {
		log.Printf("setContent(), v line error:%s", err.Error())
		s = ""
	}

	out, err := gui.View("content")
	if err != nil {
		return err
	}
	out.Clear()
	_, _ = fmt.Fprintln(out, "content: "+s)

	return nil
}

func (g *GUI) nextView(gui *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	if _, err := g.focus(name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 3 {
		g.Gui.Cursor = true
	} else {
		g.Gui.Cursor = false
	}

	active = nextIndex
	return nil
}
