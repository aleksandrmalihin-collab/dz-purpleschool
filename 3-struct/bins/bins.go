package bins

import (
	"errors"
	"math/rand/v2"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type Binlist struct {
	Bin
}

func (b *Bin) GeneratePrivate() bool {
	b.Private = status[rand.IntN(2)]
	return b.Private
}

func CreateBin(id string, private *bool, name string) (*Bin, error) {
	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}
	newBin := &Bin{
		Id:        id,
		CreatedAt: time.Now(),
		Name:      name,
	}
	if private == nil {
		newBin.GeneratePrivate()
	} else {
		newBin.Private = *private
	}
	return newBin, nil

}

func CreateBinList(id string, private *bool, name string) (*Binlist, error) {
	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}
	newBin := &Binlist{
		Bin: Bin{Id: id,
			CreatedAt: time.Now(),
			Name:      name},
	}
	if private == nil {
		newBin.GeneratePrivate()
	} else {
		newBin.Private = *private
	}
	return newBin, nil

}

var status = map[int]bool{
	0: true,
	1: false,
}
