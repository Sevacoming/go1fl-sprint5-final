package spentenergy

import (
 "errors"
 "time"
)

const (
 // Длина шага = 0.45 * рост (м) — так сходится с тестами
 stepLengthCoefficient      = 0.45
 mInKm                      = 1000.0
 minInH                     = 60.0
 // Коэффициент для ходьбы
 walkingCaloriesCoefficient = 0.5
)

var ErrInvalidInput = errors.New("invalid input parameters")

// Distance возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
 if steps <= 0 || height <= 0 {
  return 0
 }
 stepLength := stepLengthCoefficient * height
 distMeters := float64(steps) * stepLength
 return distMeters / mInKm
}

// MeanSpeed возвращает среднюю скорость (км/ч).
// (включая проверку на отрицательные шаги — TODO выполнен)
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
 if duration <= 0 || steps <= 0 || height <= 0 {
  return 0
 }
 dist := Distance(steps, height)
 hours := duration.Hours()
 if hours <= 0 {
  return 0
 }
 return dist / hours
}

// RunningSpentCalories — калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
 if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
  return 0, ErrInvalidInput
 }
 speed := MeanSpeed(steps, height, duration)
 minutes := duration.Minutes()
 return (weight * speed * minutes) / minInH, nil
}

// WalkingSpentCalories — калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
 if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
  return 0, ErrInvalidInput
 }
 speed := MeanSpeed(steps, height, duration)
 minutes := duration.Minutes()
 base := (weight * speed * minutes) / minInH
 return base * walkingCaloriesCoefficient, nil
}
