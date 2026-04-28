package main

import (
	"time"
	"fmt"
	"test/project3/bins"
)


func main() {
	BinList := []bins.Bin{}
	ptr := bins.CreateBin("qqq", true, time.Now(), "aaaa")
	BinList = append(BinList, *ptr)
	BinList = append(BinList, *(bins.CreateBin("222", true, time.Now(), "dfgdfg")))
	fmt.Println(BinList)
}
