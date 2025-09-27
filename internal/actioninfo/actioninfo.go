package actioninfo

import (
 "errors"
 "fmt"
 "strings"
 "time"

 "github.com/Yandex-Practicum/go1fl-sprint5-final/internal/daysteps"
 "github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
)

// Action хранит данные о прогулке/тренировке.
type Action struct {
 Steps    int
 Duration time.Duration
 Personal personaldata.Personal
}

// Parse разбирает строку "1000,30m" и возвращает Action.
func Parse(data string, p personaldata.Personal) (*Action, error) {
 steps, dur, err := parseData(data)
 if err != nil {
  return nil, err
 }
 return &Action{
  Steps:    steps,
  Duration: dur,
  Personal: p,
 }, nil
}

// ActionInfo возвращает строку с описанием действия.
func (a Action) ActionInfo() (string, error) {
 if a.Steps <= 0 || a.Duration <= 0 || a.Personal.Height <= 0 || a.Personal.Weight <= 0 {
  return "", errors.New("invalid action data")
 }

 ds := daysteps.DaySteps{
  Steps:    a.Steps,
  Duration: a.Duration,
  Personal: a.Personal,
 }

 info, err := ds.ActionInfo()
 if err != nil {
  return "", err
 }
 return info, nil
}

// parseData — утилита для разбора "1000,30m".
func parseData(data string) (int, time.Duration, error) {
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
