package main

import (
	"fmt"
	"test/project3/bins"
	"test/project3/file"
	"test/project3/storage"
	"time"
)

func main() {
	BinList := []bins.Bin{}
	ptr := bins.CreateBin("qqq", true, time.Now(), "aaaa")
	BinList = append(BinList, *ptr)
	BinList = append(BinList, *(bins.CreateBin("222", true, time.Now(), "dfgdfg")))
	fmt.Println(BinList)

	storage.SaveBinToFile("test1.json", BinList)

	newdata, _ := storage.ReadBinFromFile("test1.json")
	fmt.Println(newdata)

	file.ReadAnyFile("test1.json")
	fmt.Println(file.IsJsonFile("qqq.qqq"))
	fmt.Println(file.IsJsonFile("test1.json"))
}
