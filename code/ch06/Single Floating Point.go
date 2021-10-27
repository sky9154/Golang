package main
import (
	"fmt"
	"math"
)

func main() {
	pi64 := math.Pi
	pi32 := float32(math.Pi)
	fmt.Println(pi64)    // 3.141592653589793
	fmt.Println(pi32)    // 3.1415927
}