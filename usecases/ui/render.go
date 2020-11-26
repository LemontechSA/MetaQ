package ui

import (
	"log"

	"github.com/jroimartin/gocui"
	"lemontech.com/metaq/drivers/store"
)

var (
	viewArr = []string{"filter", "databases", "query", "output"}
	active  = "query"
	sidebar = 0
)

// FilterV is the filter view
var FilterV *gocui.View

// DatabasesV is the databases view
var DatabasesV *gocui.View

// QueryV is the query view
var QueryV *gocui.View

// OutputV is the output view
var OutputV *gocui.View

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func switchActive(g *gocui.Gui, nextActive string) {
	if _, err := setCurrentViewOnTop(g, nextActive); err != nil {
		return
	}

	if nextActive == "filter" || nextActive == "query" {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextActive
}

func switchV(g *gocui.Gui, v *gocui.View) error {
	nextActive := ""
	switch active {
	case "filter":
		nextActive = "databases"
	case "databases":
		nextActive = "filter"
	case "query":
		nextActive = "output"
	case "output":
		nextActive = "query"
	}
	switchActive(g, nextActive)
	return nil
}

func switchH(g *gocui.Gui, v *gocui.View) error {
	nextActive := ""
	switch active {
	case "filter":
		nextActive = "query"
	case "databases":
		nextActive = "output"
	case "query":
		nextActive = "filter"
	case "output":
		nextActive = "databases"
	}
	switchActive(g, nextActive)
	return nil
}

func switchMouse(g *gocui.Gui, v *gocui.View) error {
	switchActive(g, v.Name())
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Render will start the UI
func Render() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.Mouse = true
	g.SelFgColor = store.ENV.Color()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("filter", gocui.KeyEnter, gocui.ModNone, filter); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("databases", gocui.KeyArrowDown, gocui.ModNone, dbsDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("databases", gocui.KeyArrowUp, gocui.ModNone, dbsUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, execute); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlE, gocui.ModNone, export); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlS, gocui.ModNone, save); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("output", gocui.KeyArrowDown, gocui.ModNone, outDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("output", gocui.KeyArrowUp, gocui.ModNone, outUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlSpace, gocui.ModNone, switchH); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, switchMouse); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, switchV); err != nil {
		log.Panicln(err)
	}

	go g.Update(refresh)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
