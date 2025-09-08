package main

import (
	"fmt"
)

const usdToRub = 79.1
const usdToEur = 0.85
const eurToRub = 85.1
const eurToUsd = 1.17
const rubToUsd = 0.012
const rubToEur = 0.010

func main() {
	sourceСurrency := validateCurrency()
	amount := validateAmount()
	targetCurrency := validateCurrency()
	result := convertCurrency(sourceСurrency, amount, targetCurrency)
	fmt.Printf("%.2f %s = %.2f %s\n", amount, sourceСurrency, result, targetCurrency)
}

func validateCurrency() string {
	var currency string
	fmt.Println("Доступны валюты: USD, RUB, EUR")
	for {
		fmt.Scan(&currency)
		if currency != "USD" && currency != "RUB" && currency != "EUR" {
			fmt.Println("Ошибка! Введите USD, EUR или RUB:")
			continue
		}
		return currency
	}
}

func validateAmount() float64 {
	var amount float64
	for {
		fmt.Println("Введите число")
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Ошибка! Введен неверный формат")
			continue
		}
		if amount < 0 {
			fmt.Println("Число не может быть меньше нуля")
			continue
		}
		return amount
	}
}

func convertCurrency(sourceСurrency string, amount float64, targetCurrency string) float64 {
	var result float64
	switch {
	case sourceСurrency == "USD" && targetCurrency == "RUB":
		result = amount * usdToRub
	case sourceСurrency == "USD" && targetCurrency == "EUR":
		result = amount * usdToEur
	case sourceСurrency == "EUR" && targetCurrency == "USD":
		result = amount * eurToUsd
	case sourceСurrency == "EUR" && targetCurrency == "RUB":
		result = amount * eurToRub
	case sourceСurrency == "RUB" && targetCurrency == "EUR":
		result = amount * rubToEur
	case sourceСurrency == "RUB" && targetCurrency == "USD":
		result = amount * rubToUsd
	default:
		result = amount
	}
	return result
}
