package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("неверный формат данных: %q", datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("неверное значение шагов: %q: %w", parts[0], err)
	}

	if steps <= 0 {
		return fmt.Errorf("неверное значение шагов (отрицательное количество или ноль): %q", parts[0])
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("неверное значение продолжительности: %q: %w", parts[1], err)
	}

	if duration <= 0 {
		return fmt.Errorf("неверное значение продолжительности прогулки (отрицательное или ноль): %v", duration)
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	resultDistance := spentenergy.Distance(ds.Steps, ds.Height)

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, resultDistance, spentCalories)
	return result, nil
}
