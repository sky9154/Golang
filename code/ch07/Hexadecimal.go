package main
import "fmt"

func main() {
	var red, green, blue uint8 = 0, 141, 213
	fmt.Printf("%x %x %x\n", red, green, blue)    // 0 8d d5
	
	fmt.Printf("color：#%02x%02x%02x", red, green, blue)    // color：#008dd5
}