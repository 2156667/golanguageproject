package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var file2 = "teacherData.txt"

func TeacherLogin() {

	var (
		email string
		pass  string
	)
	fmt.Print("Email:")
	fmt.Scan(&email)
	fmt.Print("Password:")
	fmt.Scan(&pass)
	if email != "" && pass != "" {
		TeacherLoginInfo(email, pass)
	} else {
		fmt.Println("Invalid")
		main()

	}
}
func TeacherLoginInfo(email string, password string) (string, string) {
	data, err4 := ioutil.ReadFile(file2)
	if err4 != nil {
		log.Panicf("Invalid", err4)
	}
	token := strings.Split(string(data), ",")
	getfname, checkmail, checkpass := token[0], token[3], token[4]

	if strings.Contains(checkmail, email) && strings.Contains(checkpass, password) {
		fmt.Println("Welcome to the portal    " + getfname)
		TeacherPortal()

	} else {
		fmt.Println("\n" + "Invalid Input")
		TeacherLogin()
	}

	return email, password
}
func TeacherRegister() {

	var (
		fName string
		lName string
		dept  string
		email string
		pass  string
	)
	file2, err := os.Create(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	fmt.Print("First Name:")
	fmt.Scan(&fName)
	fmt.Print("Last Name:")
	fmt.Scan(&lName)
	fmt.Print("Department:")
	fmt.Scan(&dept)
	fmt.Print("Email:")
	fmt.Scan(&email)

	if !validateEmail(email) {
		fmt.Println("\n[!] Email address is Invalid" + "\n")
		TeacherRegister()
	} else {
		fmt.Print("Create Password:")
		fmt.Scan(&pass)
		if len(pass) < 5 {
			fmt.Println("\n[!] Password is short" + "\n")
			TeacherRegister()
		}
	}

	data := fName + "," + lName + "," + dept + "," + email + "," + pass
	_, err5 := file2.WriteString(data)
	if err5 != nil {
		log.Fatal(err5)
	}
	fmt.Println("\n You are now registered!")

}
func TeacherProfile() {
	data, err := ioutil.ReadFile(file2)
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(data), ",")
	cfname, clname, cdept, cmail := split[0], split[1], split[2], split[3]
	fmt.Println("Name: " + cfname + " " + clname + "\n" +
		"Course: " + cdept + "\n" +
		"Email: " + cmail + "\n")
}
