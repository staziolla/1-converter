package main
import "fmt"

func main() {
	const cUSDtoEUR = 0.85
	const cUSDtoRUB = 75.76
	EURtoRUB := 1 / cUSDtoEUR * cUSDtoRUB
	fmt.Println(EURtoRUB)
}