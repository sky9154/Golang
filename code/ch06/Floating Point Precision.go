package main

import (
	"fmt"
)

func main() {
	num := .1
	num += .2
	fmt.Println(num)    // 0.30000000000000004

	celsius := 21.0
	fmt.Println((celsius / 5.0 * 9.0) + 32)
	fmt.Println((9.0 / 5.0 * celsius) + 32)
	
	fmt.Print((celsius * 9.0 / 5.0) + 32.0)

}