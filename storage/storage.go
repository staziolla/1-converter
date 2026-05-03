package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"test/project3/bins"
)

//Сохранение bin в виде json в локальном файле
//Чтение списка bin в виде json из локального файла

func SaveBinToFile(filename string, bslice []bins.Bin) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка создания файла " + err.Error())
		return err
	}
	defer file.Close()
	data, err2 := json.Marshal(bslice)
	fmt.Println("------------")
	fmt.Println(bslice)
	fmt.Println(data)
	fmt.Println("------------")
	if err2 != nil {
		fmt.Println("Ошибка сериализации " + err2.Error())
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи файла " + err.Error())
		return err
	}
	fmt.Println("Запись успешна")
	return nil
}

func ReadBinFromFile(filename string) (*[]bins.Bin, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла " + err.Error())
		return nil, err
	}
	var b []bins.Bin
	json.Unmarshal(data, &b)
	return &b, nil
}
