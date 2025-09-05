package main

import "fmt"

const usdTOEur = 0.85
const usdTORub = 81.31
const EurTOrub = usdTORub / usdTOEur

func main() {
	fmt.Println("=========Калькулятор расчеты валюты=========")
	fmt.Println("1 Евро равен", EurTOrub, "рублей.")
}
