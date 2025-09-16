package main

import (
	"3-struct/api"
	"3-struct/bins"
	"3-struct/file"
	"3-struct/storage"
	"fmt"
	"strings"
)

func main() {

	id := promtDataToBin("Введите id: ")
	name := promtDataToBin("Введите имя: ")
	private := promtStatus("Введите статус true/false: ")
	bin1, err := bins.CreateBin(id, private, name)
	if err != nil {
		fmt.Println("Неверное имя или id!")
		return
	}

	binlist1, err := bins.CreateBinList(id, private, name)
	if err != nil {
		fmt.Println("Неверное имя или id!")
		return
	}
	fmt.Println(*bin1)
	fmt.Println(*binlist1)

	api.JsonForBin()
	file.ReadFile()
	file.WriteFile()
	storage.Store()
}

func promtDataToBin(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}

func promtStatus(prompt string) *bool {
	fmt.Println(prompt)
	var s string
	fmt.Scanln(&s)

	s = strings.TrimSpace(strings.ToLower(s))
	if s == "" {
		return nil
	}
	v := s == "true" || s == "t" || s == "1" || s == "yes" || s == "y"
	return &v
}
