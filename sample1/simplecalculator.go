package main

import (
	"fmt"
)

var num1 = 0
var num2 = 0
var choice = 0

func main() {
	fmt.Println("============Go Language Calculator===========")
	fmt.Println("Hello! Welcome to Basic Calculator! \nPlease enter your choice of operation.")
	fmt.Println("1. Addition \n2. Subtraction \n3. Multiplication \n4. Division")
	fmt.Println("=============================================")
	fmt.Print("Choice: ")
	fmt.Scanln(&choice)
	operations()

	fmt.Print("Try again? [y/n]: ")
	var ans = "n"
	fmt.Scanln(&ans)
	if ans == "y" || ans == "Y" {
		main()
	} else {
		fmt.Println("Thank you!")
	}
}

func operations() {
	switch choice {
	case 1:
		fmt.Println("Enter the first number: ")
		fmt.Scanln(&num1)
		fmt.Println("Enter the second number: ")
		fmt.Scanln(&num2)
		fmt.Println("The result is: ", add())
		break
	case 2:
		fmt.Println("Enter the first number: ")
		fmt.Scanln(&num1)
		fmt.Println("Enter the second number: ")
		fmt.Scanln(&num2)
		fmt.Println("The result is: ", subtract())
		break
	case 3:
		fmt.Println("Enter the first number: ")
		fmt.Scanln(&num1)
		fmt.Println("Enter the second number: ")
		fmt.Scanln(&num2)
		fmt.Println("The result is: ", multiply())
		break
	case 4:
		fmt.Println("Enter the first number: ")
		fmt.Scanln(&num1)
		fmt.Println("Enter the second number: ")
		fmt.Scanln(&num2)
		fmt.Println("The result is: ", divide())
		break
	default:
		fmt.Println("You have entered an invalid key. Please try again!")
	}
}

func add() int {
	return num1 + num2
}

func subtract() int {
	return num1 - num2
}

func multiply() int {
	return num1 * num2
}

func divide() int {
	return num2 / num1
}
