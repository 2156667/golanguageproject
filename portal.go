package main

//
//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"os"
//	"strings"
//)
//
//var path = "users.txt"
//
//func main() {
//	fmt.Println("_______________________")
//	fmt.Println("       Welcome          ")
//	fmt.Println("_______________________")
//	fmt.Println("		Log in          ")
//
//	var (
//		name string
//		pass string
//	)
//	fmt.Print("Name:")
//	fmt.Scan(&name)
//	fmt.Print("Password:")
//	fmt.Scan(&pass)
//	if name != "" && pass != "" {
//		logIn(name, pass)
//	} else {
//		fmt.Println("Invalid")
//		main()
//
//	}
//}
//func register() {
//	var (
//		name string
//		pass string
//	)
//
//	var file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//	fmt.Print("Enter name:")
//	fmt.Scanln(&name)
//	fmt.Print("Enter pass:")
//	fmt.Scanln(&pass)
//	if len(pass) < 5 {
//		fmt.Println("too short")
//		register()
//
//		//if name == "" && pass == "" {
//		//	fmt.Println("No input!")
//		//	os.Exit(0)
//	} else {
//		fmt.Println("\n" + "You are now Registered!")
//		_, err2 := file.WriteString("\n" + name + "," + pass + "\n")
//		main()
//		if err2 != nil {
//			log.Fatal(err2)
//
//		}
//	}
//}
//func logIn(name string, pass string) (string, string) {
//	data, err := ioutil.ReadFile("users.txt")
//	if err != nil {
//		log.Panicf("Invalid", err)
//	}
//	token := strings.Split(string(data), ",")
//	checkname, checkpass := token[0], token[1]
//	if strings.Contains(checkname, name) && strings.Contains(checkpass, pass) {
//		fmt.Println("Welcome to the portal    " + name)
//		portal()
//
//	} else {
//		var (
//			choice int
//		)
//		fmt.Println("\n" + "Invalid Input")
//		fmt.Println("_______________________")
//		fmt.Println("[0] Register")
//		fmt.Println("[1] Back")
//		fmt.Println("_______________________")
//
//		fmt.Scan(&choice)
//		switch choice {
//		case 0:
//			register()
//			break
//		case 1:
//			main()
//			break
//		default:
//		}
//	}
//	return name, pass
//}
//func portal() {
//	var choice1 int
//	fmt.Println("_______________________")
//	fmt.Println("what do you want to do?")
//	fmt.Println("[1] Back")
//	fmt.Println("[2] Logout")
//	fmt.Println("_______________________")
//
//	fmt.Scan(&choice1)
//	switch choice1 {
//	case 1:
//		main()
//		break
//	case 2:
//		var choice2 string
//		fmt.Println("Do you want to exit?")
//		fmt.Println("[y] yes    [n] no")
//		fmt.Scanln(&choice2)
//		if choice2 == "y" {
//			fmt.Println(".....Bye")
//			os.Exit(0)
//		} else {
//			portal()
//		}
//		break
//	default:
//	}
//}
