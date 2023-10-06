package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	if angka > 0  {
		fmt.Println("positif")
	} else if angka < 0 {
		fmt.Println("negatif")
	} else {
		fmt.Println("nol")
	}
}
