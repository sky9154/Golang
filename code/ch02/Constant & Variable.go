package main
import (
	"fmt"
)

func main() {
	const BigNumber = 10000000
	var MiniNumber = 1000
	fmt.Printf("BigNumber + MiniNumber = %d\n", BigNumber + MiniNumber)
	// BigNumber + MiniNumber = 10001000
	
	MiniNumber = 10
	fmt.Printf("BigNumber + MiniNumber = %d", BigNumber + MiniNumber)
	// BigNumber + MiniNumber = 10000010

	const (
		const1 = 10
		const2 = 20
	)
	const const3, const4 = 30, 40
	fmt.Printf("\n\nconst1, const2, const3, const4 = %d, %d, %d, %d", const1, const2, const3, const4)
	// const1, const2, const3, const4 = 10, 20, 30, 40

	var (
		var1 = 10
		var2 = 20
	)
	var var3, var4 = 30, 40
	fmt.Printf("\nvar1, var2, var3, var4 = %d, %d, %d, %d", var1, var2, var3, var4)
	// var1, var2, var3, var4 = 10, 20, 30, 40
}
