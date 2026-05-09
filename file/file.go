package file

import (
	"fmt"
	"os"
	"strings"
)

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


// type JsonDB struct {
// 	filename string
// }

// func NewJsonDB(name string) *JsonDB {
// 	return &JsonDB{filename: name,
// 	}
// }

// func (db *JsonDB) Read() ([]byte, error) {
// 	data, err := os.ReadFile(db.filename)
// 	if err != nil {
// 		outputs.PrintError(err)
// 		return nil, err
// 	}
// 	fmt.Println(string(data))
// 	return data, nil
// }

// func (db *JsonDB) Write(content []byte) {
// 	file, err := os.Create(db.filename)
// 	if err != nil {
// 		outputs.PrintError(err)
// 		return
// 	}
// 	defer file.Close()
// 	_, err = file.Write(content)
// 	if err != nil {
// 		outputs.PrintError(err)
// 		return
// 	}
// 	fmt.Println("Запись успешна")
// }

