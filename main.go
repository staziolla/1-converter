package main
import "fmt"

func main() {
	usum, ufrom, uto := getUserData()
	calcResult(usum, ufrom, uto)
}

func getUserData() (float64, string, string) {
	var usum float64
	var ufrom, uto string
	vald, vale, valr := "usd", "eur", "rub"
	for {
		fmt.Printf("Из валюты (%s, %s, %s): ", vald, vale, valr)
		fmt.Scan(&ufrom)
		if checkFromVal(ufrom) {
			break
		}
		fmt.Println("Неверно указана валюта")
	}
	for {
		fmt.Print("Сумма: ")
		fmt.Scan(&usum)
		if checkSum(usum) {
			break
		}
		fmt.Println("Неверно указана сумма")
	}
	var otherVals string
	switch ufrom {
	case "usd":
		otherVals = "eur, rub"
	case "eur":
		otherVals = "usd, rub"
	case "rub":
		otherVals = "usd, eur"
	}
	for {
		fmt.Printf("В валюту (%s): ", otherVals)
		fmt.Scan(&uto)
		if checkToVal(uto, ufrom) {
			break
		}
		fmt.Println("Неверно указана валюта")
	}
	return usum, ufrom, uto
}

func checkFromVal(ufrom string) bool {
	return ufrom == "usd" || ufrom == "eur" || ufrom == "rub" 
}

func checkSum(summa float64) bool {
	return summa > 0
}
func checkToVal(uto, ufrom string) bool {
	return (uto == "usd" || uto == "eur" || uto == "rub") && uto != ufrom 
}

func calcResult(summa float64, fromVal string, toVal string) float64 {
	const cUSDtoEUR = 0.85
	const cUSDtoRUB = 75.76
	var result float64
	type tcurval map[string]float64
	type tcur map[string]tcurval

	curvalmap := tcur{}
	curvalmap["usd"] = tcurval{"eur": cUSDtoEUR, "rub": cUSDtoRUB}
	curvalmap["eur"] = tcurval{"usd": 1 / cUSDtoEUR, "rub": cUSDtoRUB / cUSDtoEUR}
	curvalmap["rub"] = tcurval{"eur": 1 / (cUSDtoRUB * cUSDtoEUR), "usd": 1 / cUSDtoRUB}

	result = summa * curvalmap[fromVal][toVal]
	fmt.Printf("%0.2f %s = %.2f %s", summa, fromVal, result, toVal)
	return result
}