package main

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			// i ve j'yi köşede olup olmamasına göre kontrol ediyoruz
			if (i == 1 && j == 1) || (i == 1 && j == x) || (i == y && j == 1) || (i == y && j == x) {
				fmt.Print("৹")
			} else if j == 1 || j == x {
				fmt.Print("|") // i ve j'yi sağ ve solda olup olmamasına göre kontrol ediyoruz
			} else if i == 1 || i == y {
				fmt.Print("-") // i ve j'yi üst ve altta olup olmamasına göre kontrol ediyoruz
			} else {
				fmt.Print(" ") // kalan yerlere boşluk koyuyoruz
			}
		}
		fmt.Print("\n") // her satırın sonunda alt satıra geçmesi için
	}
}

func main() {
	QuadA(5, 3)
}
