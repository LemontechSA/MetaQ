package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"lemontech.com/metaq/domain"
	"lemontech.com/metaq/drivers/db"
	"lemontech.com/metaq/drivers/store"
	"lemontech.com/metaq/usecases/cli"
	"lemontech.com/metaq/usecases/ui"
)

func main() {
	source, err := cli.Run()
	if err != nil {
		return
	}

	store.ENV = domain.ENV{}
	if err = store.ENV.Load(source); err != nil {
		fmt.Println(err)
		return
	}

	if err = db.Connect(); err != nil {
		fmt.Println("\x1b[31m>> Couldn't reach DB\x1b[0m")
		return
	}

	ui.Render()
}
