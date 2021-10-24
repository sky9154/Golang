package main
import (
	"fmt"
	"strings"
)

func main() {
	var kirito = "starburst stream"
	var Boolean = strings.Contains(kirito, "starburst") // strings.Contains() 判斷變數中是否有包含指定字串
	fmt.Print("變數 kirito 是否含有 starburst：", Boolean)
}