package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"lemontech.com/metaq/drivers/store"
)

func layout(g *gocui.Gui) error {
	sidebar = 24
	maxX, maxY := g.Size()
	if int(0.2*float32(maxX)) > sidebar {
		sidebar = int(0.2 * float32(maxX))
	}
	if v, err := g.SetView("source", 1, 0, sidebar, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = false
		v.Autoscroll = false
		v.Frame = false
		v.FgColor = store.ENV.Color()
		fmt.Fprint(v, fmt.Sprintf("source: %s", store.ENV.NAME))
	}
	if v, err := g.SetView("filter", 0, 3, sidebar, 5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "filter"
		v.Editable = true
		v.Wrap = false
		FilterV = v

		fmt.Fprintln(FilterV, store.ENV.FILTER)
	}
	if v, err := g.SetView("databases", 0, 6, sidebar, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "databases"
		v.Editable = false
		v.Wrap = false
		v.Autoscroll = false
		DatabasesV = v
	}
	if v, err := g.SetView("query", sidebar+2, 0, maxX-1, maxY/2-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "query"
		v.Editable = true
		v.Wrap = true
		v.Autoscroll = false
		QueryV = v

		if _, err = setCurrentViewOnTop(g, "query"); err != nil {
			return err
		}
	}
	if v, err := g.SetView("output", sidebar+2, maxY/2, maxX-1, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "output"
		v.Editable = false
		v.Wrap = false
		v.Autoscroll = false
		OutputV = v
	}
	if v, err := g.SetView("help", 1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = false
		v.Autoscroll = false
		v.Frame = false
		v.FgColor = store.ENV.Color()
		fmt.Fprint(v, "LeftClick: change active panel, Tab: ▲ ▼, ^Space: ◄ ►, Enter: apply filter, ^Q: execute query, ^E: export to csv, ^S: save query")
	}
	return nil
}
