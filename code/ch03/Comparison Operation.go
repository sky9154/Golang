package main
import (
	"fmt"
)

func main() {
	fmt.Println("5 == 5 = ", 5 == 5)   // 5 等於 5 True
	fmt.Println("5 != 5 = ", 5 != 5) // 5 沒有不等於 5 False
	fmt.Println("5 < 10 = ", 5 < 10)   // 5 小於 10 True
	fmt.Println("5 <= 10 = ", 5 <= 10) // 5 小於等於 10 True
	fmt.Println("10 > 5 = ", 10 > 5)   // 10 大於 5 True
	fmt.Println("10 <= 5 = ", 10 >= 5) // 10 大於等於 5 True
}