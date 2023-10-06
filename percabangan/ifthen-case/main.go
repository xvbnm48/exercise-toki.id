package main

import "fmt"

func main() {
	var angka int

	fmt.Scan(&angka)

	if angka >= 0 && angka <= 9 {
		fmt.Println("satuan")
	} else if angka >= 10 && angka <= 99 {
		fmt.Println("puluhan")
	} else if angka >= 100 && angka <= 999 {
		fmt.Println("ratusan")
	} else if angka >= 1000 && angka <= 9999{
		fmt.Println("ribuan")
	} else if angka >= 10000 && angka <= 99999 {
		fmt.Println("puluhribuan")
	} else {
		fmt.Println("error")
	}
}
