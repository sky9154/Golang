package main
import (
	"fmt"

)

func main() {
	// 用 var 聲明變數
	var i = 0
	for i = 10; i >= 0; i -- {
		fmt.Print(i)
	}

	// 用 := 聲明變數
	for j := 10; j >= 0; j -- {
		fmt.Print(j)
	}
}