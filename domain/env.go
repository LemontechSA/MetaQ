package domain

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jroimartin/gocui"
)

// ENV is a struct to transport environment related attributes
type ENV struct {
	NAME   string
	ENGINE string
	DBURL  string
	FILTER string
	COLOR  string
}

// Save will persist the ENV in the given env name
func (e *ENV) Save(name string) (err error) {
	env, err := godotenv.Unmarshal(fmt.Sprintf("ENGINE=%s\nDBURL=%s\nFILTER=%s\nCOLOR=%s", e.ENGINE, e.DBURL, e.FILTER, e.COLOR))
	if err != nil {
		return
	}

	err = os.Mkdir(fmt.Sprintf("%s/.metaq/sources/%s", Home, name), 0755)
	if err != nil {
		return
	}

	godotenv.Write(env, fmt.Sprintf("%s/.metaq/sources/%s/cfg", Home, name))
	return
}

// Load will read a persisted ENV in the given env name
func (e *ENV) Load(name string) (err error) {
	env, err := godotenv.Read(fmt.Sprintf("%s/.metaq/sources/%s/cfg", Home, name))
	if err != nil {
		fmt.Println(err)
		return
	}
	e.NAME = name
	e.ENGINE = env["ENGINE"]
	e.DBURL = env["DBURL"]
	e.FILTER = env["FILTER"]
	e.COLOR = env["COLOR"]
	return
}

// Color will return the corresponding gocui.Attribute value
func (e *ENV) Color() (color gocui.Attribute) {
	switch e.COLOR {
	case "Black":
		return gocui.ColorBlack
	case "Blue":
		return gocui.ColorBlue
	case "Cyan":
		return gocui.ColorCyan
	case "Green":
		return gocui.ColorGreen
	case "Magenta":
		return gocui.ColorMagenta
	case "Red":
		return gocui.ColorRed
	case "White":
		return gocui.ColorWhite
	case "Yellow":
		return gocui.ColorYellow

	}
	return
}
