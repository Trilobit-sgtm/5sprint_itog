package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		if strings.HasSuffix(info, "\n") {
			fmt.Print(info)
		} else {
			fmt.Println(info)
		}
	}
}
