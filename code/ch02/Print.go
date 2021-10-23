package main
import (
	"fmt"
)

func main() {
	fmt.Print("Kirito ")
	fmt.Print("Asuna")
	// Kirito Asuna

	fmt.Println("\n\nKirito")
	fmt.Println("Asuna")
	// Kirito
	// Asuna

	fmt.Printf("\nKirito %s %d\n", "Asuna", 78763)
	// Kirito Asuna 78763
	fmt.Printf("10 + 2 = %d", 10 + 2)
	// 10 + 2 = 12
	
	fmt.Printf("\n\nKirito%15v\n", "Asuna")
	// Kirito          Asuna -> Asuna 向左添加 15 格空格
	fmt.Printf("%-15vKirito", "Asuna")
	// Asuna          Kirito -> Asuna 向右添加 15 格空格
}