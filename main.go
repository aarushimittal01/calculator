package main

import (
	"fmt"

	"github.com/aarushimittal01/calculator/calculate"
	"github.com/aarushimittal01/calculator/menu"
)

func main() {
	var val int = 0

	for val != 5 && val < 6 {
		menu.Show_menu()
		fmt.Scanln(&val)
		fmt.Println("Input 2 Numbers\n")
		var num1, num2, ans int
		fmt.Scanln(&num1, &num2)

		switch val {
		case 1:
			ans = calculate.Add(num1, num2)
		case 2:
			ans = calculate.Subtract(num1, num2)
		case 3:
			ans = calculate.Multiply(num1, num2)
		case 4:
			ans = calculate.Divide(num1, num2)
			if ans == 0 {
				fmt.Println("Cannot Divide by 0")
				val = 5
			}
		}
		fmt.Println("Answer is : %v", ans)
	}
}
