package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию -- OK
	array := strings.Split(data, ",")
	fmt.Println(array)
	if len(array) != 2 {
		return 0, 0, errors.New("Длинна массивов не 2")
	}

	stepsCount, err := strconv.Atoi(array[0])
	if err != nil || stepsCount <= 0 {
		return 0, 0, errors.New("Ошибка приведения шагов ")
	}

	timeCount, err := time.ParseDuration(array[1])
	if err != nil {
		return 0, 0, err
	}
	if timeCount <= 0 {
		return 0, 0, errors.New("продолжительность должна быть больше 0")
	}
	return stepsCount, timeCount, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию -- OK
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		log.Println("количество шагов должно быть больше 0")
		return ""
	}

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}

	distance := float64(steps) * stepLength / mInKm

	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps,
		distance,
		calories,
	)
}
