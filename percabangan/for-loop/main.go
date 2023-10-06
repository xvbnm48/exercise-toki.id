package main

import "fmt"


func main(){
	var n int 
	fmt.Scan(&n)
	totalBebek := 0
	for i := 0; i < n; i++ {
		var b int 
		fmt.Scan(&b)
		totalBebek += b
	}

	fmt.Println(totalBebek)
}