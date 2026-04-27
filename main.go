package main

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func createBin(id string, private bool, createdAt time.Time, name string) *Bin {
	b := Bin{id: id, private: private, createdAt: createdAt, name: name}
	return &b
}

func main() {
	BinList := []Bin{}
	ptr := createBin("qqq", true, time.Now(), "aaaa")
	BinList = append(BinList, *ptr)
	//BinList = append(BinList, *(createBin("222", true, time.Now(), "dfgdfg")))
	//fmt.Println(BinList)
}
