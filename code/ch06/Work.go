package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 隨機種子

	money := [3] float64 {	// 美分硬幣
		0.05, 
		0.10, 
		0.25,
	}
	var sum float64 = 0 // 儲蓄罐總量
	for(sum <= 20) {
		sum += money[rand.Intn(2)]
		fmt.Printf("目前儲蓄罐有 %.2f 元\n", sum)
	}
}