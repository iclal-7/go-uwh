package main

import "fmt"

func SatirYazdir(ilkKar, ortaKar, sonKar string, uzunluk int) {
	fmt.Print(ilkKar)
	for i := 0; i < uzunluk-2; i++ {
		fmt.Print(ortaKar)
	}
	if uzunluk > 1 {
		fmt.Print(sonKar)
	}

	fmt.Println() // \n ile aynı şeyi yapar, bir alt satıra geçer
}

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		if i == 1 {
			SatirYazdir("A", "B", "A", x)
		} else if i > 1 && i < y {
			SatirYazdir("B", " ", "B", x)
		} else {
			SatirYazdir("C", "B", "C", x)
		}
	}
}

func main() {
	QuadC(5, 3)
}
