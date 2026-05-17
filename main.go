package main

import (
	"fmt"
	"test/project3/bins"
	"test/project3/file"
	"time"
)

func main() {
	stor := file.NewJsonDB("test1.json")
	BinList := bins.CreateBinsListWithDb(stor)
	BinList.Append("qqq", true, time.Now(), "aaaa")
	BinList.Append("222", true, time.Now(), "dfgdfg")
	//fmt.Println(BinList.db.filename)
	fmt.Println(len(BinList.Bins))
	for _, elem := range BinList.Bins {
		fmt.Println(elem)
	}
	BinList.Save()
	BinList.ShowData()

	//	vault := account.NewVault(files.NewJsonDB("data.json"))
	//BinList := []bins.Bin{}
	//	ptr := bins.CreateBin("qqq", true, time.Now(), "aaaa")
	//	BinList = append(BinList, *ptr)
	//	BinList = append(BinList, *(bins.CreateBin("222", true, time.Now(), "dfgdfg")))
	//	fmt.Println(BinList)

	// storage.SaveBinToFile("test1.json", BinList)

	// newdata, _ := storage.ReadBinFromFile("test1.json")
	// fmt.Println(newdata)

	// file.ReadAnyFile("test1.json")
	// fmt.Println(file.IsJsonFile("qqq.qqq"))
	// fmt.Println(file.IsJsonFile("test1.json"))
}
