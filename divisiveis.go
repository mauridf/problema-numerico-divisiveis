package main

import "fmt"

func main() {
	fmt.Println("Divisíveis por 3:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("\nDivisíveis por 3 (PIN) e 5 (PAN):")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("PIN")
		}
		if i%5 == 0 {
			fmt.Print("PAN")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
