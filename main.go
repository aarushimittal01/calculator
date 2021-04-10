package main

import (
	"fmt"
	"os"

	"calculator/menu"
	"calculator/operation"
)

func main() {
	var val int

	fmt.Printf(menu.Show_menu())
	fmt.Scanln(&val)

	for val != 5 && val < 5 {
		fmt.Printf("Input 2 Numbers\n")
		var num1, num2, ans int
		fmt.Scanln(&num1, &num2)

		switch val {
		case 1:
			ans = operation.Add(num1, num2)
		case 2:
			ans = operation.Subtract(num1, num2)
		case 3:
			ans = operation.Multiply(num1, num2)
		case 4:
			ans = operation.Divide(num1, num2)
			if ans == 0 {
				fmt.Println("\nCannot Divide by 0")
				val = 5
			}
		default:
			os.Exit(1)
		}
		fmt.Printf("\nAnswer is : %v", ans)
		fmt.Printf(menu.Show_menu())
		fmt.Scanln(&val)
	}
}
