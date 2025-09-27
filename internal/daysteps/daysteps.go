package daysteps

import (
 "errors"
 "fmt"
 "strings"
 "time"

 "github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
 "github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// DaySteps — данные о прогулке.
type DaySteps struct {
 Steps    int
 Duration time.Duration
 personaldata.Personal
}

var ErrInvalidDaySteps = errors.New("invalid day steps input")

// Parse разбирает строку вида "1000,30m" или "678,0h50m".
func (ds *DaySteps) Parse(datastring string) error {
 steps, dur, err := parsePackage(datastring)
 if err != nil {
  return err
 }
 ds.Steps = steps
 ds.Duration = dur
 return nil
}

// ActionInfo возвращает строку с данными о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
 if ds.Steps <= 0 || ds.Duration <= 0 || ds.Height <= 0 || ds.Weight <= 0 {
  return "", ErrInvalidDaySteps
 }

 dist := spentenergy.Distance(ds.Steps, ds.Height)
 cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
 if err != nil {
  return "", err
 }

 out := fmt.Sprintf(
  "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
  ds.Steps, dist, cal,
 )
 return out, nil
}

// parsePackage — утилита для теста: парсит "1000,30m".
func parsePackage(data string) (int, time.Duration, error) {
 parts := strings.Split(data, ",")
 if len(parts) != 2 {
  return 0, 0, errors.New("wrong data format")
 }

 var steps int
 if _, err := fmt.Sscanf(parts[0], "%d", &steps); err != nil {
  return 0, 0, err
 }
 dur, err := time.ParseDuration(parts[1])
 if err != nil {
  return 0, 0, err
 }
 return steps, dur, nil
}

// DayActionInfo — утилита для теста: создаёт DaySteps и возвращает ActionInfo.
func DayActionInfo(data string, weight, height float64) string {
 ds := DaySteps{
  Personal: personaldata.Personal{
   Weight: weight,
   Height: height,
  },
 }
 if err := ds.Parse(data); err != nil {
  return ""
 }
 s, err := ds.ActionInfo()
 if err != nil {
  return ""
 }
 return s
}
