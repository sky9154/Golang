package main
import (
	"fmt"
)

func main() {
	fmt.Println("請問 2021 年是閏年嗎？")
	var year = 2021
	var leap = year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) // 如果符合 year % 400 == 0 則不繼續判斷的條件後面
	if leap {
		fmt.Print(year, " 年是閏年")
	} else {
		fmt.Print(year, " 年不是閏年")
	}
}