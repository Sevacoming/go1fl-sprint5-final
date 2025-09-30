package personaldata

import (
	"errors"
	"fmt"
)

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}

// Validate проверяет адекватность антропометрии.
func (p Personal) Validate() error {
	if p.Weight <= 0 || p.Height <= 0 {
		return errors.New("invalid personal data")
	}
	return nil
}
