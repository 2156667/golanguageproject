package main

import (
	"fmt"
)

var (
	choice  int
	name    string
	age     string
	add     string
	occ     string
	email   string
	stat    string
	cont    string
	rel     string
	adpt    string
	dadpt   string
	alladpt string
	dint    string
	tint    string
)

func main() {

	fmt.Println("---")
	fmt.Println("PET ADOPTION")
	fmt.Println("---")
	fmt.Println("Please Choose from the Option")
	fmt.Println("[1] Application Form")
	fmt.Println("[2] Scheduling Form")
	fmt.Println("[3] Zoom Interview")
	fmt.Println("[4] Shelter Visits")
	fmt.Println("[5]	Back")
	fmt.Println("---")

	fmt.Print("Choice: ")
	fmt.Scanln(&choice)
	options()

}
func options() {
	switch choice {
	case 1:
		fmt.Println("[-Application Form-]")
		fmt.Println("Name: ")
		fmt.Scanln(&name)
		fmt.Println("Age: ")
		fmt.Scanln(&age)
		fmt.Println("Address: ")
		fmt.Scanln(&add)
		fmt.Println("Occupation: ")
		fmt.Scanln(&occ)
		fmt.Println("Email: ")
		fmt.Scanln(&email)
		fmt.Println("Status: ")
		fmt.Scanln(&stat)
		fmt.Println("Contact: ")
		fmt.Scanln(&cont)
		fmt.Println("Relationship: ")
		fmt.Scanln(&rel)
		fmt.Println("---")
		fmt.Println("What are you looking to adopt")
		fmt.Scanln(&adpt)
		fmt.Println("Do you know the name of the animal you want to adopt?")
		fmt.Scanln(&dadpt)
		fmt.Println("Describe your ideal pet,including its sex,age, appearance, temperament, etc.")
		fmt.Scanln(&alladpt)
		break
	case 2:
		fmt.Println("[-Scheduling Form-]")
		fmt.Println("[Full Name: ]")
		fmt.Scanln(&name)
		fmt.Println("[Phone: ]")
		fmt.Scanln(&cont)
		fmt.Println("Email: ")
		fmt.Scanln(&email)
		fmt.Println("[Requested Date for Zoom Interview: ]")
		fmt.Scanln(&dint)
		fmt.Println("[Requested Time for Zoom Interview: ]")
		fmt.Scanln(&tint)
		break
	case 3:
		fmt.Println("[---Zoom Interview---]")
		fmt.Println("Email: ")
		fmt.Scanln(&email)
		fmt.Println("[Full Name: ]")
		fmt.Scanln(&name)
		fmt.Println("[Phone: ]")
		fmt.Scanln(&cont)
		break
	case 4:
		fmt.Println("[-Shelter Visits-]")
		fmt.Println("Name: ")
		fmt.Scanln(&name)
		fmt.Println("Address: ")
		fmt.Scanln(&add)
		fmt.Println("[Phone: ]")
		fmt.Scanln(&cont)
		break
	default:
		fmt.Print("Exit")
		fmt.Scanln(&choice)
		options()
	}
}
