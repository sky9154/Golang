package main
import (
	"fmt"
)

func main() {
	var num = 2.5
	num = num * 3.5
	fmt.Println(num)	// 8.75
	num = 2.5
	num *= 3.5
	fmt.Println(num)	// 8.75
}