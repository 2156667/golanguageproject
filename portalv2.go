package main

import (
	"fmt"
	"os"
)

var choices int

func main() {
	fmt.Println("_______________________")
	fmt.Println("       Welcome          ")
	fmt.Println("_______________________")
	fmt.Println("Choose Action    ")
	fmt.Println("[1] Login   ")
	fmt.Println("[2] Register     ")
	fmt.Println("_______________________")
	fmt.Scanln(&choices)
	switch choices {
	case 1:
		Login()
		break
	case 2:
		Register()
		break
	default:
	}
}
func Login() {
	fmt.Println("Log in as       ")
	fmt.Println("[1] Student     ")
	fmt.Println("[2] Teacher     ")
	fmt.Println("[3] Back     ")
	fmt.Println("_______________________")
	fmt.Scanln(&choices)
	switch choices {
	case 1:
		StudentLogin()
		break
	case 2:
		TeacherLogin()
		break
	default:
		main()
	}
}
func Register() {
	fmt.Println("Register as       ")
	fmt.Println("[1] Student     ")
	fmt.Println("[2] Teacher     ")
	fmt.Println("[3] Back")
	fmt.Println("_______________________")
	fmt.Scanln(&choices)
	switch choices {
	case 1:
		StudentRegister()
		break

	case 2:
		TeacherRegister()
		break

	default:
		main()
	}
}
func StudentPortal() {
	var choice1 int
	fmt.Println("_______________________")
	fmt.Println("what do you want to do?")
	fmt.Println("[1] Profile")
	fmt.Println("[2] CheckGrade")
	fmt.Println("[3] Back")
	fmt.Println("[4] Logout")
	fmt.Println("_______________________")

	fmt.Scan(&choice1)
	switch choice1 {
	case 1:
		StudentProfile()
		break
	case 2:
		fmt.Println("This is the grades ")
		break
	case 3:
		main()
		break
	case 4:
		var choice2 string
		fmt.Println("Do you want to exit?")
		fmt.Println("[y] yes    [n] no")
		fmt.Scanln(&choice2)
		if choice2 == "y" {
			fmt.Println(".....Bye")
			os.Exit(0)
		} else {
			StudentPortal()
		}
		break
	default:
	}
}
func TeacherPortal() {
	var choice int

	fmt.Println("_______________________")
	fmt.Println("what do you want to do?")
	fmt.Println("[1] Profile")
	fmt.Println("[2] PostGrades")
	fmt.Println("[3] Back")
	fmt.Println("[4] Logout")
	fmt.Println("_______________________")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		TeacherProfile()
		break
	case 2:
		fmt.Println("This is the grades ")
		break
	case 3:
		main()
		break
	case 4:
		var choice2 string
		fmt.Println("Do you want to exit?")
		fmt.Println("[y] yes    [n] no")
		fmt.Scanln(&choice2)
		if choice2 == "y" {
			fmt.Println(".....Bye")
			os.Exit(0)
		} else {
			StudentPortal()
		}
		break
	default:
	}
}
