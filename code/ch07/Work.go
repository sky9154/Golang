package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 隨機種子

	money := [3] int64 {	// 美分硬幣
		5, 
		10, 
		25,
	}
	var sum int64 = 0 // 儲蓄罐總量
	for(sum <= 20) {
		if (sum == 20) {
			break
		}
		sum += money[rand.Intn(2)]
		fmt.Printf("目前儲蓄罐有 %d 元\n", sum)
	}
}