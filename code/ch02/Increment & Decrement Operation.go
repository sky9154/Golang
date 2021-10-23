package main
import (
	"fmt"
)

func main() {
	var sum = 40

	// 自增運算符
	sum = sum + 1
	fmt.Println(sum)	// 41
	
	// 簡寫
	sum = 40
	sum += 1
	fmt.Println(sum)	// 41

	// 再簡寫
	sum = 40
	sum ++
	fmt.Println(sum)	// 41
	
	// 自減運算符
	sum = 40
	sum = sum - 1
	fmt.Println(sum)	// 39
	
	// 簡寫
	sum = 40
	sum -= 1
	fmt.Println(sum)	// 39
	
	// 再簡寫
	sum = 40
	sum --
	fmt.Println(sum)	// 39
}