package main

import (
	"fmt"
)
func(main)
	number1 :=0
	number2 :=0
	number :=0

	fmt.Print("Enter value of number1: ")
	fmt.Scan(&number1)
	fmt.Print("Enter value of number2: ")
	fmt.Scan(&number2)
	fmt.Print("Enter number for operation 1=sum, 2=diff, 3=prod, 4=guot: ")
	fmt.Scan(&number)

	sum := number1 + number2
	diff := number1 - number2
	prod := number1 * number2
	quot := number1 / number2

	switch number {
	case 1:
		fmt.Print("Sum is", sum)
	case 2:
		fmt.Print("Difference is", diff)
	case 3:
		fmt.Print("Product is", prod)
	case 4:
		fmt.Print("Quotient is", quot)
	}
}
