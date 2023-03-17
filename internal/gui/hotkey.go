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
func keyBinders() []KeyBinder {
	var binders []KeyBinder
	binders = append(binders, KeyBinder{
		Key:      gocui.MouseLeft,
		ViewName: "menu",
		Action:   setContent,
	})
	return binders
}

// 關閉程式
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// 設置content layout內容
func setContent(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	s, err := v.Line(cy)
	if err != nil {
		log.Printf("setContent(), v line error:%s", err.Error())
		s = ""
	}

	out, err := g.View("content")
	if err != nil {
		return err
	}
	out.Clear()
	_, _ = fmt.Fprintln(out, "content: "+s)

	return nil
}
