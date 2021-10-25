package main
import (
	"fmt"

)

func main() {
	var sum = 0
	for sum < 10 {
		var i = 1
		sum += i
		i ++
	}
	fmt.Print(sum)
}