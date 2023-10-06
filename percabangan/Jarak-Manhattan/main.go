package main

import "fmt"

func main() {
	var x1, x2, y1, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	hasil := abs(x1-x2) + abs(y1-y2)

	fmt.Println(hasil)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
