package main

import (
	"fmt"
)

func main() {
	kirito :=1.0 / 3
	fmt.Println(kirito)    // 0.3333333333333333
	fmt.Printf("%f\n", kirito)    // 0.333333
	fmt.Printf("%.2f\n", kirito)    // 0.33
	fmt.Printf("%4.2f", kirito)    // 0.33 不滿足 4 位則補 0
}