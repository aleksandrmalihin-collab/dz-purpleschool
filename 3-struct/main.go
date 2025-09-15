package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type Binlist struct {
	Bin
}

func (b *Bin) generatePrivate() bool {
	b.private = status[rand.IntN(2)]
	return b.private
}

func createBin(id string, private *bool, name string) (*Bin, error) {
	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}
	newBin := &Bin{
		id:        id,
		createdAt: time.Now(),
		name:      name,
	}
	if private == nil {
		newBin.generatePrivate()
	} else {
		newBin.private = *private
	}
	return newBin, nil

}

func createBinList(id string, private *bool, name string) (*Binlist, error) {
	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}
	newBin := &Binlist{
		Bin: Bin{id: id,
			createdAt: time.Now(),
			name:      name},
	}
	if private == nil {
		newBin.generatePrivate()
	} else {
		newBin.private = *private
	}
	return newBin, nil

}

var status = map[int]bool{
	0: true,
	1: false,
}

func main() {

	id := promtDataToBin("Введите id: ")
	name := promtDataToBin("Введите имя: ")
	private := promtStatus("Введите статус true/false: ")
	bin1, err := createBin(id, private, name)
	if err != nil {
		fmt.Println("Неверное имя или id!")
		return
	}

	binlist1, err := createBinList(id, private, name)
	if err != nil {
		fmt.Println("Неверное имя или id!")
		return
	}
	fmt.Println(*bin1)
	fmt.Println(*binlist1)
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
