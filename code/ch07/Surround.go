package main
import (
	"fmt"
)

func main() {
	var red uint8 = 255
	red ++
	fmt.Println(red)    // 超過最大值 255 回到最小值 0

	var num int8 = 127
	num ++
	fmt.Println(num)    // 超過最大值 127 回到最小值 -128
}