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
