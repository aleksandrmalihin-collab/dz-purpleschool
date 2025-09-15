package main

import "fmt"

func main() {
	mapCurrency := map[string]float64{
		"usdToRub": 79.1,
		"usdToEur": 0.85,
		"eurToRub": 85.1,
		"eurToUsd": 1.17,
		"rubToUsd": 0.012,
		"rubToEur": 0.010,
	}
	sourceСurrency := validateCurrency()
	amount := validateAmount()
	targetCurrency := validateCurrency()
	result := convertCurrency(sourceСurrency, amount, targetCurrency, mapCurrency)
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
func convertCurrency(sourceСurrency string, amount float64, targetCurrency string, mapCurrency map[string]float64) float64 {
	var result float64
	switch {
	case sourceСurrency == "USD" && targetCurrency == "RUB":
		result = amount * mapCurrency["usdToRub"]
	case sourceСurrency == "USD" && targetCurrency == "EUR":
		result = amount * mapCurrency["usdToEur"]
	case sourceСurrency == "EUR" && targetCurrency == "USD":
		result = amount * mapCurrency["eurToUsd"]
	case sourceСurrency == "EUR" && targetCurrency == "RUB":
		result = amount * mapCurrency["eurToRub"]
	case sourceСurrency == "RUB" && targetCurrency == "EUR":
		result = amount * mapCurrency["rubToEur"]
	case sourceСurrency == "RUB" && targetCurrency == "USD":
		result = amount * mapCurrency["rubToUsd"]
	default:
		result = amount
	}
	return result
}
