package main

import (
	"fmt"
	"math"
)

func tubBoluvchilarniTop(son int) []int {
	tublar := []int{}

	for i := 2; i <= son; i++ {
		for son%i == 0 && tub(i) {
			tublar = append(tublar, i)
			son /= i
		}
	}

	return tublar
}

func tub(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// Test qilish
	son1 := 24
	fmt.Printf("Kirish: %d\n", son1)
	fmt.Println("Chiqish:", tubBoluvchilarniTop(son1))
	son2 := 25
	fmt.Printf("Kirish: %d\n", son2)
	fmt.Println("Chiqish:", tubBoluvchilarniTop(son2))
}
