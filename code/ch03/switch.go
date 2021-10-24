package main
import (
	"fmt"
)

func main() {
	var score = 70
	switch {
	case score < 0, score > 100:    // 如果 score 小於 0 或是大於 100 輸出 Error
		fmt.Print("Error")
	case score < 60:    // 如果 score 小於 60 輸出不及格
		fmt.Print("不及格")
	case score >= 60:    // 如果 score 大於等於 60 輸出及格
		fmt.Print("及格")
	default:
		fmt.Print("Error")
	}
}