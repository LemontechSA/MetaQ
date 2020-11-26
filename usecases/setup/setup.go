package setup

import (
	"fmt"
	"os"

	"lemontech.com/metaq/domain"
)

// Check will create the required folders if they don't exist
func Check() (err error) {
	if _, err = os.Stat(fmt.Sprintf("%s/.metaq/sources", domain.Home)); os.IsNotExist(err) {
		err = os.Mkdir(fmt.Sprintf("%s/.metaq/", domain.Home), 0755)
		if err != nil {
			return
		}
		err = os.Mkdir(fmt.Sprintf("%s/.metaq/sources", domain.Home), 0755)
		if err != nil {
			return
		}
	}
	return
}
