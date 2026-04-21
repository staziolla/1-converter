package main
import ("fmt"
		"strings"
		"slices"
		"os"
		"bufio"
		"strconv"
		)

func main() {
	oper, data := getUserInput()
	fmt.Println(oper)
	fmt.Println(data)
	res := calcOper(oper, data)
	fmt.Printf("Результат %s для %v = %f\n", oper, data, res)
}

func calcOper(oper string, data []float64) float64{
	var result float64
	ldata := len(data)
	if ldata == 0 {
		return result
	}
	switch {
	case oper == "SUM":
		result = getSum(data)
	case oper == "AVG":
		result = getSum(data) / float64(len(data))
	case oper == "MED":
		slices.Sort(data)
		if ldata%2 == 0 {
			result = (data[ldata / 2] + data[ldata / 2 - 1]) / 2
		} else {
			result = data[(ldata - 1) / 2]
		}
	}
	return result
}

func getSum(arr []float64) float64 {
	s := 0.0
	for _, value := range arr {
		s += value
	}
	return s
}

func getUserInput() (string, []float64) {
	var oper string
	var data []float64
	reader := bufio.NewReader(os.Stdin)
	allOpers := []string{"AVG", "SUM", "MED"}
	fmt.Printf("Укажите операцию %v или 'Q' для выхода\n", allOpers)
	for {
		oper, _ = reader.ReadString('\n')
		oper = strings.ToUpper(strings.TrimSpace(oper))
		idx := slices.Index(allOpers, oper)
		if oper == "Q" {
			return oper, nil
		} else if idx > -1 {
			break
		}
		fmt.Println("Неверно указана операция", oper)
	}
	fmt.Println("Укажите значения через запятую и пробел")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
 	sarr := strings.Split(s, ",")
	for _, value := range sarr {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		x, err := strconv.ParseFloat(value, 64)
		//fmt.Println(err)
		if err != nil {
			fmt.Printf("Значение %v не является числом\n", value)
			continue
		}
		data = append(data, x)
	}
	return oper, data
}	


