package main

import "fmt"

const usdTOEur = 0.85
const usdTORub = 81.31
const EurTOrub = usdTORub / usdTOEur

func main() {
	var num int
	var valuta1, valuta2 string
	fmt.Scan(&valuta1, &valuta2)
	fmt.Scan(&num)
	fmt.Println("=========Калькулятор расчеты валюты=========")
	fmt.Println("1 Евро равен", EurTOrub, "рублей.")
	a := stdin()
	fmt.Println(a)
	nullfunc(a, valuta1, valuta2)
}

func stdin() int {
	var data int
	fmt.Scan(&data)
	return data
}

func nullfunc(amount int, fromCurrency string, toCurrency string) {

}
