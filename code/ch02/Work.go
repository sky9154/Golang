package main
import (
	"fmt"
)

func main() {
	const distance = 56000000
	var day = 28
	day *= 24 // day to hour
	var speed = distance / day // calculate the speed
	fmt.Print("時速 ", speed + 1, " km/h，即可於 28 天內到達 Malacandra")
}