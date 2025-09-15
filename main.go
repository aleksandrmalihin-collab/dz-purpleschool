package main

import "fmt"

func main() {
	mapCurrency := map[string]float64{
		"USDTORUB": 79.1,
		"USDTOEUR": 0.85,
		"EURTORUB": 85.1,
		"EURTOUSD": 1.17,
		"RUBTOUSD": 0.012,
		"RUBTOEUR": 0.010,
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

func convertCurrency(sourceCurrency string, amount float64, targetCurrency string, mapCurrency map[string]float64) float64 {
	if sourceCurrency == targetCurrency {
		return amount
	}

	key := sourceCurrency + "TO" + targetCurrency
	if rate, exists := mapCurrency[key]; exists {
		return amount * rate
	}

	return amount // fallback для неподдерживаемых пар
}
