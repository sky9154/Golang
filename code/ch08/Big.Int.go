package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigInt := big.NewInt(2147483647)
	fmt.Println(bigInt)		// 2147483647

	bigNumber := new(big.Int)
	bigNumber.SetString("2147483648", 10)	// int64 最大值為 2147483647，10 為 10 進制
	fmt.Print(bigNumber)	// 2147483648
}