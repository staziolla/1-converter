package file

import ("os"
		"fmt"
		"strings"
)

//Чтение любого файла
//Проверка что это json расширение файла

func ReadAnyFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err!=nil {
		fmt.Println("Ошибка чтения файла")
		return nil, err
	}
	return file, nil
}

func IsJsonFile(filename string) bool {
	return strings.ToLower(filename)[len(filename) - 5:] == ".json" 
}