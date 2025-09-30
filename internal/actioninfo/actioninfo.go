package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, p DataParser) {
	if p == nil || len(dataset) == 0 {
		return
	}
	for _, s := range dataset {
		if err := p.Parse(s); err != nil {
			log.Println("parse:", err)
			return
		}
	}
	out, err := p.ActionInfo()
	if err != nil {
		log.Println("action info:", err)
		return
	}
	fmt.Println(out)
}
