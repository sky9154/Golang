package main
import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i ++ {	// i 的初始值設為 0 ，當 i 小於 10 的時候執行程式碼，執行完後 i + 1
		fmt.Print(i, " ")
	}
	fmt.Println("End !!!")
	// 0 1 2 3 4 5 6 7 8 9 End !!!
}