package main
import (
	"fmt"
)

func main() {
	var score = 70
	if score < 60 && score > 0 {    // 如果 score 小於 60 且大於 0 輸出不及格
		fmt.Print("不及格")
	} else if score >= 60 && score <= 100 {    // 如果 score 大於等於 60 且小於 100 輸出及格
		fmt.Print("及格")
	} else if score < 0 || score > 100 {    // 如果 score 小於 0 或者大於 100 就輸出 Error
		fmt.Print("Error")
	} else {    // 如果上述條件都不成立就輸出 Error
		fmt.Print("Error")
	}
}