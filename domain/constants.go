package domain

import (
	"fmt"
	"os"
)

// Home is the path to the home folder
var Home string

// Sources is the path to the sources folder
var Sources string

// Source is the path to a single source folder
var Source string

// CFG is the path to a single cfg file
var CFG string

// AddSrcStr is the string for asking to create a new source
var AddSrcStr = "Add a new source"

func init() {
	Home, _ = os.UserHomeDir()
	Sources = fmt.Sprintf("%s/.metaq/sources", Home)
	Source = fmt.Sprintf("%s/.metaq/sources/%s", Home, "%s")
	CFG = fmt.Sprintf("%s/.metaq/sources/%s/cfg", Home, "%s")
}
