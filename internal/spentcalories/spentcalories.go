package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	array := strings.Split(data, ",")
	if len(array) != 3 {
		return 0, "", 0, errors.New("Длинна массивов в parseTraining не 3")
	}

	stepsCount, err := strconv.Atoi(array[0])
	if err != nil || stepsCount <= 0 {
		return 0, "", 0, errors.New("Ошибка приведения шагов в пакете spentcalories - ParseTraining")
	}

	activityInfo := array[1]

	activityTime, err := time.ParseDuration(array[2])
	if err != nil {
		return 0, "", 0, err
	}
	if activityTime <= 0 {
		return 0, "", 0, errors.New("продолжительность должна быть больше 0")
	}

	return stepsCount, activityInfo, activityTime, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceInfo := stepLength * float64(steps)
	distanceInfo = distanceInfo / mInKm
	return distanceInfo
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distanceInfo := distance(steps, height)
	hours := duration.Hours()

	if hours <= 0 {
		return 0
	}

	return distanceInfo / hours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию -- OK
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64

	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:

		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		log.Println(err)
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		activity,
		duration.Hours(),
		dist,
		speed,
		calories,
	), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию -- OK
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Ошибка входящего значения в пакете spentcalories - RunningSpentCalories")
	}

	speedInfo := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	spentCaloriesCount := (weight * speedInfo * durationInMinutes) / minInH
	return spentCaloriesCount, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию -- OK
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("Ошибка входящего значения в пакете spentcalories - WalkingSpentCalories")
	}

	speedInfo := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCaloriesCount := (weight * speedInfo * durationInMinutes) / minInH
	spentCaloriesCount = spentCaloriesCount * walkingCaloriesCoefficient
	return spentCaloriesCount, nil
}
