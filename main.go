package main

import "fmt"

const usdTOEur = 0.85
const usdTORub = 81.31
const EurTOrub = usdTORub / usdTOEur

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println("=========Калькулятор расчеты валюты=========")
	fmt.Println("1 Евро равен", EurTOrub, "рублей.")
	a := stdin()
	fmt.Println(a)
	nullfunc(num, 20.1, 1000)
}

func stdin() int {
	var data int
	fmt.Scan(&data)
	return data
}

func nullfunc(number int, cntuer float64, cntrub float64) {

}
