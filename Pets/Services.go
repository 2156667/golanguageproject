package main

import (
	"fmt"
)

func Services() {
	var (
		choice int
	)
	fmt.Println("-------------------------")
	fmt.Println("PET SERVICES")
	fmt.Println("-------------------------")
	fmt.Print("Please Choose from the Option")
	fmt.Print("[1] Grooming Salon")
	fmt.Print("[2] Vet Aid")
	fmt.Print("[3] Training Classes")
	fmt.Print("[4] Pet's Hotel")
	fmt.Print("[5] Pet Camps")
	fmt.Print("[6]	Back")
	fmt.Println("-------------------------")
	fmt.Scanln(choice)

	switch choice {
	case 1:
		fmt.Print("[----GROOMING SALON-----]")

		break
	case 2:
		fmt.Print("[---- VET AID -----]")
		break
	case 3:
		fmt.Print("[---- TRAINING CLASSES -----]")
		break
	case 4:
		fmt.Print("[---- PET'S HOTEL -----]")
		break
	case 5:
		fmt.Print("[---- PET CAMP ----]-")
		break
	default:

	}

}
