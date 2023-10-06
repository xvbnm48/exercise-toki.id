package main

import (
	"fmt"
)

func main() {
	var N float64
	fmt.Scan(&N)

	floorN := 0
	ceilN := 0

	if N >= 0 {
		floorN = int(N)
		fmt.Println("kondisi awal floorN ", floorN)
		if float64(floorN) != N {
			ceilN = floorN + 1
			fmt.Println("kondisi baru ceilN : ", ceilN)
		} else {
			ceilN = floorN
			fmt.Println("kondisi ceiln kalo false:", ceilN)
		}
	} else {
		ceilN = int(N)
		fmt.Println("ceiln awal:", ceilN)
		fmt.Println("kondisi ketika ceilN tidak lebih dari 0:", ceilN)
		if float64(ceilN) != N {
			floorN = ceilN - 1
			fmt.Println("floor n kalo n tidak lebih 0 ", floorN)
		} else {
			floorN = ceilN
		}
	}

	fmt.Printf("%d %d\n", floorN, ceilN)
}
