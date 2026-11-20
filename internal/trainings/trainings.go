package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("неверный формат данных: %q", datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("неверное значение шагов: %q: %w", parts[0], err)
	}

	if steps <= 0 {
		return fmt.Errorf("неверное значение шагов (отрицательнок количество или ноль): %q", parts[0])
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("неверное значение продолжительности: %q: %w", parts[2], err)
	}

	if duration <= 0 {
		return fmt.Errorf("неверное значение продолжительности прогулки (отрицательное или 0):%q: %v", parts[2], duration)
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {

	var spentCalories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	resultDistance := spentenergy.Distance(t.Steps, t.Height)
	resultMeanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), resultDistance, resultMeanSpeed, spentCalories)

	return result, nil
}
