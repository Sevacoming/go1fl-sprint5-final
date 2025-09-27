package personaldata

import (
 "errors"
 "fmt"
)

// Personal — данные пользователя.
type Personal struct {
 Name   string
 Weight float64 // кг
 Height float64 // м
}

var ErrInvalidPersonal = errors.New("invalid personal data")

// Validate — простая проверка корректности.
func (p Personal) Validate() error {
 if p.Weight <= 0 || p.Height <= 0 {
  return ErrInvalidPersonal
 }
 return nil
}

// Print — выводит данные пользователя на экран.
func (p Personal) Print() {
 fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
