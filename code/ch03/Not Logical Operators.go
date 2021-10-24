package main
import (
	"fmt"
)

func main() {
	var Kirito = true
	var Asuna = false
	if Kirito == !Asuna {
		fmt.Print("LINK START!!!")
	}
}