package file

import (
	"fmt"
	"os"
	"strings"
)

type JsonDB struct {
	filename string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{filename: name,
	}
}

func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(data))
	return data, nil
}

func (db *JsonDB) Write(content []byte) error {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Запись успешна")
	return  nil
}





//Чтение любого файла
//Проверка что это json расширение файла


func ReadAnyFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return nil, err
	}
	return file, nil
}

func IsJsonFile(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".json")
}




