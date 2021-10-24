package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var random = rand.Intn(99) + 1 // 隨機產生 1 ~ 100 的整數
	var number = 0 // 輸入的變數
	for {
		fmt.Print("輸入一個 1 ~ 100 的整數：")
		fmt.Scanln(&number) // 將輸入的整數存入 number 這個變數
		if number < 1 || number > 100 {
			fmt.Println("輸入錯誤")
			continue
		}
		fmt.Println("答案是：", random)
		if number > random {
			fmt.Println("太大")
		} else if number < random {
			fmt.Println("太小")
		} else if number == random {
			fmt.Println("猜對了")
			break
		}
	}
}