package main

import "fmt"

func kenarmi(i, j, x, y int) bool {
	return i == 1 || i == y || j == 1 || j == x
}
func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 { // sol üst köşe
				fmt.Print("A")
			} else if i == y && j == 1 { // sol alt köşe
				fmt.Print("C")
			} else if i == 1 && j == x { // sağ üst köşe
				fmt.Print("C")
			} else if i == y && j == x { // sağ alt köşe
				fmt.Print("A")
			} else if kenarmi(i, j, x, y) {
				fmt.Print("B")
			} else {
				fmt.Print(" ") // kalan yerlere boşluk koyuyoruz
			}
		}
		fmt.Print("\n") // her satırın sonunda alt satıra geçmesi için
	}
}

func main() {
	QuadE(5, 3)
}
