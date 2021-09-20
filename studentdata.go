package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var file = "studentData.txt"

func StudentLogin() {

	var (
		email string
		pass  string
	)
	fmt.Print("Email:")
	fmt.Scan(&email)
	fmt.Print("Password:")
	fmt.Scan(&pass)
	if email != "" && pass != "" {
		LoginInfo(email, pass)
	} else {
		fmt.Println("Invalid")
		main()

	}
}
func LoginInfo(email string, password string) (string, string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panicf("Invalid", err)
	}
	token := strings.Split(string(data), ",")
	getfname, checkmail, checkpass := token[0], token[3], token[4]

	if strings.Contains(checkmail, email) && strings.Contains(checkpass, password) {

		fmt.Println("\n\nWelcome to the portal    " + getfname)
		StudentPortal()

	} else {
		fmt.Println("\n" + "Invalid Input")
		StudentLogin()
	}

	return email, password
}
func StudentRegister() {
	var (
		fName  string
		lName  string
		course string
		email  string
		pass   string
	)
	file, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Print("First Name:")
	fmt.Scan(&fName)
	fmt.Print("Last Name:")
	fmt.Scan(&lName)
	fmt.Print("Course:")
	fmt.Scan(&course)
	fmt.Print("Email:")
	fmt.Scan(&email)
	if !validateEmail(email) {
		fmt.Println("\n[!] Email address is Invalid" + "\n")
		StudentRegister()
	} else {
		fmt.Print("Create Password:")
		fmt.Scan(&pass)
		if len(pass) < 5 {
			fmt.Println("\n[!] Password is short" + "\n")
			StudentRegister()
		}
	}
	data := fName + "," + lName + "," + course + "," + email + "," + pass
	_, err2 := file.WriteString(data)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("\n You are now registered!")
}
func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)

}
func StudentProfile() {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(data), ",")
	cfname, clname, crce, cmail := split[0], split[1], split[2], split[3]
	fmt.Println("Name: " + cfname + " " + clname + "\n" +
		"Course: " + crce + "\n" +
		"Email: " + cmail + "\n")

}
