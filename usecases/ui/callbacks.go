package ui

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/olekukonko/tablewriter"
	"lemontech.com/metaq/drivers/db"
	"lemontech.com/metaq/drivers/store"
)

func refresh(g *gocui.Gui) (err error) {
	DatabasesV.Clear()
	time.Sleep(1 * time.Second)

	store.DBs, _ = db.ShowDBs()
	for i := range store.DBs {
		if strings.Contains(store.DBs[i].Name, store.ENV.FILTER) {
			store.DBs[i].Selected = true
			fmt.Fprintln(DatabasesV, store.DBs[i].Name)
		} else {
			store.DBs[i].Selected = false
		}
	}

	return

}

func filter(g *gocui.Gui, v *gocui.View) (err error) {
	DatabasesV.Clear()
	store.ENV.FILTER = strings.TrimSpace(FilterV.Buffer())

	for i := range store.DBs {
		if strings.Contains(store.DBs[i].Name, store.ENV.FILTER) {
			store.DBs[i].Selected = true
			fmt.Fprintln(DatabasesV, store.DBs[i].Name)
		} else {
			store.DBs[i].Selected = false
		}
	}

	return
}

func execute(g *gocui.Gui, v *gocui.View) (err error) {
	statement := strings.TrimSpace(QueryV.Buffer())

	OutputV.Clear()
	fmt.Fprintln(OutputV, "EXECUTING")
	fmt.Fprintln(OutputV, statement)

	store.Data, err = db.Query(store.DBs, statement)
	if err != nil {
		fmt.Fprintln(OutputV, err)
		return nil
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)

	table.SetHeader(store.Data.Headers)
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)
	table.AppendBulk(store.Data.Rows)
	table.Render()

	OutputV.Clear()
	fmt.Fprintln(OutputV, tableString.String())
	return
}

func export(g *gocui.Gui, v *gocui.View) (err error) {
	file, _ := os.Create("result.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(store.Data.Headers)
	for _, value := range store.Data.Rows {
		writer.Write(value)
	}
	return
}

func save(g *gocui.Gui, v *gocui.View) (err error) {
	statement := strings.TrimSpace(QueryV.Buffer())

	file, _ := os.Create("query.sql")
	defer file.Close()

	file.WriteString(statement)
	return
}

func dbsUp(g *gocui.Gui, v *gocui.View) (err error) {
	x, y := DatabasesV.Origin()
	DatabasesV.SetOrigin(x, y-1)
	return
}

func dbsDown(g *gocui.Gui, v *gocui.View) (err error) {
	x, y := DatabasesV.Origin()
	DatabasesV.SetOrigin(x, y+1)
	return
}

func outUp(g *gocui.Gui, v *gocui.View) (err error) {
	x, y := OutputV.Origin()
	OutputV.SetOrigin(x, y-1)
	return
}

func outDown(g *gocui.Gui, v *gocui.View) (err error) {
	x, y := OutputV.Origin()
	OutputV.SetOrigin(x, y+1)
	return
}
