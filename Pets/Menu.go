package main

import (
	"fmt"
)

var (
	choice int
)

func main() {

	fmt.Println("-------------------------")
	fmt.Println("Welcome to QiQian's Pet")
	fmt.Println("-------------------------")
	fmt.Println("Please choose what service would you to apply")
	fmt.Println("[1] Pet Services")
	fmt.Println("[2] Pet Adoption")
	fmt.Println("[3] Learning Center")
	fmt.Println("[4] Exit")
	fmt.Println("-------------------------")

	fmt.Print("Choice: ")
	fmt.Scanln(&choice)
	options()

}
func options() {
	switch choice {
	case 1:
		fmt.Print("[---- Pet Services -----]")
		break
	case 2:
		fmt.Print("[---- Pet Adoption -----]")
		break
	case 3:
		fmt.Print("[---- Learning Center-----]")
		break
	default:
		fmt.Print("Exit")

	}
}
