package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, fmt.Errorf("неверное значение шагов (отрицательное количество или ноль): %d", steps)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("неверное значение веса пользователя (отрицательное или ноль): %.2f", weight)
	}

	if height <= 0 {
		return 0, fmt.Errorf("неверное значение роста пользователя (отрицательнок или ноль): %.2f", height)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("неверное значение продолжительности бега (отрицательное или ноль): %v", duration)
	}

	speed := MeanSpeed(steps, height, duration)

	spentCalories := (weight * speed * duration.Minutes()) / minInH

	return spentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, fmt.Errorf("неверное значение шагов (отрицательное количество или ноль): %d", steps)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("неверное значение веса пользователя (отрицательное или ноль): %.2f", weight)
	}

	if height <= 0 {
		return 0, fmt.Errorf("неверное значение роста пользователя (отрицательное или ноль): %.2f", height)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("неверное значение продолжительности бега (отрицательное или ноль): %v", duration)
	}

	speed := MeanSpeed(steps, height, duration)

	return (weight * speed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	if steps <= 0 {
		return 0
	}
	if height <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}

	distanceKm := Distance(steps, height)
	durationHours := duration.Hours()

	return distanceKm / durationHours
}

func Distance(steps int, height float64) float64 {

	stepLength := stepLengthCoefficient * height
	return (float64(steps) * stepLength) / mInKm
}
