package main

import "fmt"

func main() {
	var num /* int64 uint64 float64 */ = 87e63
	fmt.Println(num)	// 8.7e+64
	fmt.Printf("%T", num)	// float64
}