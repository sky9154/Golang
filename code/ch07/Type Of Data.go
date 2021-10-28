package main
import (
	"fmt"
)

func main() {
	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a)
	// Type string for text

	b := 42
	fmt.Printf("Type %T for %[1]v\n", b)
	// Type int for 42

	c := 3.14
	fmt.Printf("Type %T for %[1]v\n", c)
	// Type float64 for 3.14

	d := true
	fmt.Printf("Type %T for %[1]v\n", d)
	// Type bool for true
}