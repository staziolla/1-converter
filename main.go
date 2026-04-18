package main
import "fmt"

func main() {
	const cUSDtoEUR = 0.85
	const cUSDtoRUB = 75.76
	EURtoRUB := 1 / cUSDtoEUR * cUSDtoRUB
	fmt.Println(EURtoRUB)

	usum, ufrom, uto := getUserData()
	calcResult(usum, ufrom, uto)
}

func getUserData() (float64, string, string) {
	var usum float64
	var ufrom, uto string
	fmt.Print("Сумма: ")
	fmt.Scan(&usum)
	fmt.Print("из валюты : ")
	fmt.Scan(&ufrom)
	fmt.Print("в валюту : ")
	fmt.Scan(&uto)
	return usum, ufrom, uto

}

func calcResult(summa float64, fromVal string, toVal string) float64 {
	fmt.Print(summa, fromVal, toVal)
	return 0
}