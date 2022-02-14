package helpers

import (
	"errors"
	"fmt"
	"log"
)

func GetOptions(arguments []string) ([]string, error) {

	allApps := []string{"Console", "Feed", "Meat", "MEZ", "Plant", "Pork", "Poultry", "TradeCo", "Turkey"}

	if len(arguments) < 1 {
		return nil, errors.New("error: you must pass at least one argument")

	} else if len(arguments) > 2 {
		log.Fatalf("error: too many arguments: %s", arguments)

	} else if len(arguments) == 2 {
		app := arguments[0]
		element := arguments[1]

		for i := 0; i <= len(allApps)-1; i++ {
			if app == allApps[i] {
				break
			} else if i == len(allApps)-1 {
				log.Fatalf("error: wrong application name: %s", app)
			}
		}

		fmt.Println("\n\nStart delete shared descendants.")
		return []string{app, element}, nil
	}

	return nil, errors.New("error: unrecognized arguments")
}
