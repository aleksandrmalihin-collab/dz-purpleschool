package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("Введите числа через запятую без пробелов!")
	fmt.Scan(&input)
	convertedInput := converterInputToInt(input)
	myResult := operation(convertedInput)
	fmt.Println("Мой итоговый результат:", myResult)
}

func converterInputToInt(input string) []float64 {
	var numbers []float64
	parts := strings.Split(input, ",")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		num, err := strconv.Atoi(p)
		if err != nil {
			fmt.Println("Ошибка конвертации строки!", err)
			continue
		}
		numbers = append(numbers, float64(num))
	}
	return numbers
}

func operation(numbers []float64) float64 {
	var userChoice string
	summ := 0.0
	result := 0.0
	for {
		fmt.Println("Выберите операцию из предложенных: AVG, SUM, MED")
		fmt.Scan(&userChoice)
		switch {
		case userChoice == "AVG":
			for i := 0; i < len(numbers); i++ {
				summ += float64(numbers[i])
			}
			result = summ / float64(len(numbers))

		case userChoice == "SUM":
			for i := 0; i < len(numbers); i++ {
				summ += float64(numbers[i])
			}
			result = summ

		case userChoice == "MED":
			sort.Float64s(numbers)
			if len(numbers)%2 == 1 {
				result = float64(numbers[len(numbers)/2])
			} else {
				result = (float64(numbers[len(numbers)/2]) + float64(numbers[len(numbers)/2-1])) / 2.0
			}

		default:
			fmt.Println("Неверная операция. Введите заново!")
			continue
		}
		return result
	}
}
