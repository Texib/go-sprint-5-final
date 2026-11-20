package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {

	for _, character := range dataset {
		err := dp.Parse(character)
		if err != nil {
			log.Println(err)
		}
		message, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
		} else {
			fmt.Print(message)
		}
	}
}
