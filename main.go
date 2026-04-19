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
		if ufrom == "usd" || ufrom == "eur" || ufrom == "rub" {
			break
		}
		fmt.Println("Неверно указана валюта")
	}
	for {
		fmt.Print("Сумма: ")
		fmt.Scan(&usum)
		if usum > 0 {
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
		if (uto == "usd" || uto == "eur" || uto == "rub") && uto != ufrom {
			break
		}
		fmt.Println("Неверно указана валюта")
	}
	return usum, ufrom, uto

}

func calcResult(summa float64, fromVal string, toVal string) float64 {
	const cUSDtoEUR = 0.85
	const cUSDtoRUB = 75.76
	var result float64
	switch {
	case fromVal == "usd" && toVal == "eur":
		result = summa * cUSDtoEUR		
	case fromVal == "eur" && toVal == "usd":
		result = summa / cUSDtoEUR		
	case fromVal == "usd" && toVal == "rub":
		result = summa * cUSDtoRUB		
	case fromVal == "rub" && toVal == "usd":
		result = summa / cUSDtoEUR		
	case fromVal == "eur" && toVal == "rub":
		result = summa * cUSDtoRUB	/ cUSDtoEUR	
	case fromVal == "rub" && toVal == "eur":
		result = summa / cUSDtoRUB * cUSDtoEUR		
	}
	//EURtoRUB := 1 / cUSDtoEUR * cUSDtoRUB
	fmt.Printf("%0.2f %s = %.2f %s", summa, fromVal, result, toVal)
	return result
}