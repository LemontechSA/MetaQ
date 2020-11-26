package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"lemontech.com/metaq/domain"
	"lemontech.com/metaq/drivers/db"
	"lemontech.com/metaq/usecases/setup"
)

// Run will start de cli prompts
func Run() (source string, err error) {
	err = setup.Check()
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := listSources()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(f) == 0 {
		addSource()
	}

	for source, err = chooseSource(); source == domain.AddSrcStr; {
		if err != nil {
			return
		}
		addSource()
		source, err = chooseSource()
	}

	return
}

func chooseSource() (source string, err error) {
	f, err := listSources()
	items := append(f, domain.AddSrcStr)
	prompt := promptui.Select{
		Label:    "Source",
		Items:    items,
		HideHelp: true,
		Stdout:   &nosound{},
	}

	_, source, err = prompt.Run()
	return
}

func listSources() ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(domain.Sources)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

func addSource() {
	if !newSource() {
		return
	}

	name, err := newName()
	if err != nil {
		return
	}

	engine, err := newEngine()
	if err != nil {
		return
	}

	url, err := newURL()
	if err != nil {
		return
	}

	filter, err := newFilter()
	if err != nil {
		return
	}

	color, err := newColor()
	if err != nil {
		return
	}

	env := domain.ENV{
		ENGINE: engine,
		DBURL:  url,
		FILTER: filter,
		COLOR:  color,
	}
	env.Save(name)
}

func newSource() bool {
	prompt := promptui.Prompt{
		Label:     domain.AddSrcStr,
		IsConfirm: true,
		Stdout:    &nosound{},
	}

	_, err := prompt.Run()

	if err == nil {
		return true
	}
	return false
}

func newName() (name string, err error) {
	prompt := promptui.Prompt{
		Label:    "Name",
		Validate: unusedFolder,
		Stdout:   &nosound{},
	}

	name, err = prompt.Run()
	return
}

func newEngine() (engine string, err error) {
	prompt := promptui.Select{
		Label:    "DB Engine",
		Items:    []string{"MySQL", "PgSQL"},
		HideHelp: true,
		Stdout:   &nosound{},
	}

	_, engine, err = prompt.Run()
	return
}

func newURL() (url string, err error) {
	for url == "" {
		prompt := promptui.Prompt{
			Label:    "DB URL",
			Validate: urlValidator,
			Default:  "root:password@tcp(localhost)/",
			Stdout:   &nosound{},
		}

		url, err = prompt.Run()

		if err != db.CheckURL(url) {
			fmt.Println("\x1b[31m>> Couldn't reach DB\x1b[0m")
			url = ""
		}
	}

	return
}

func newFilter() (filter string, err error) {
	prompt := promptui.Prompt{
		Label:  "Default Filter",
		Stdout: &nosound{},
	}

	filter, err = prompt.Run()
	return
}

func newColor() (color string, err error) {
	prompt := promptui.Select{
		Label:    "Accent Color",
		Items:    []string{"Black", "Blue", "Cyan", "Green", "Magenta", "Red", "White", "Yellow"},
		HideHelp: true,
		Stdout:   &nosound{},
	}

	_, color, err = prompt.Run()
	return
}

func emptyValidator(input string) error {
	length := len(input)
	if length == 0 {
		return errors.New("Can't be empty")
	}
	return nil
}

func unusedFolder(input string) error {
	length := len(input)
	if length == 0 {
		return errors.New("Can't be empty")
	}
	if strings.Contains(input, " ") {
		return errors.New("Can't have spaces")
	}
	if _, err := os.Stat(fmt.Sprintf(domain.Source, input)); !os.IsNotExist(err) {
		return errors.New("Already used")
	}
	return nil
}

func urlValidator(input string) error {
	length := len(input)
	if length == 0 {
		return errors.New("Can't be empty")
	}
	return nil
}

type nosound struct{}

func (ns *nosound) Write(b []byte) (int, error) {
	const charBell = 7 // c.f. readline.CharBell
	if len(b) == 1 && b[0] == charBell {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

// Close implements an io.WriterCloser over os.Stderr.
func (ns *nosound) Close() error {
	return os.Stderr.Close()
}
