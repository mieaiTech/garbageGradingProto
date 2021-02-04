package main

import (
	"fmt"
)

// Sign In Function
func signIn() {
	fmt.Print("Email: ")
	var email string
	fmt.Scan(&email)
	switch email {
	default:
		fmt.Print("Password: ")
		var password string
		fmt.Scan(&password)
		fmt.Println("Welcome, USER#")
	}
}

// Student function: This asks the user which student they want to grade
func student() {
	fmt.Print("Which student would you like to grade: ")
	var stu string
	fmt.Scan(&stu)
	switch stu {
	case "mieai":
		fmt.Println("Student found: mieaiDiarrhea")
	default:
		fmt.Println("Student not found!")
	}
}

// Asking function: This asks the user if they want to grade an assignment
func asking() {
	fmt.Print("Would you like to grade an assignment: ")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "yes", "YES", "Yes":
		fmt.Println("Ok, this process will be ready in a few seconds.")

	case "no", "NO", "No":
		fmt.Println("Ok, you will be taken back to the sign in page.")
	default:
		fmt.Println("Invalid")
	}
}

// Sign-Up Function
func signUp() {
	fmt.Print("First Name: ")
	var firstName string
	fmt.Scan(&firstName)
	fmt.Print("Last Name: ")
	var lastName string
	fmt.Scan(&lastName)
	fmt.Print("Email: ")
	var sEmail string
	fmt.Scan(&sEmail)
	fmt.Print("Password: ")
	var sPassword string
	fmt.Scan(&sPassword)
}

// Calculations function
func calc() {
	// This asks the user how many assignments they would like to grade
	fmt.Print("How many assignments would you like to grade: ")
	var oneThroughFive string
	fmt.Scan(&oneThroughFive)
	switch oneThroughFive {
	case "1":
		fmt.Print("Numerator: ")
		var numerator float64
		fmt.Scan(&numerator)
		fmt.Print("Denominator: ")
		var denominator float64
		fmt.Scan(&denominator)
		total := numerator / denominator
		totals := fmt.Sprintf("%g", total*100)
		totalsf := "This student scored: " + totals + "%"
		fmt.Println(totalsf)
	}
}
func chosen() {
	fmt.Print("Would you like to grade an assignment: ")
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "yes", "YES", "Yes":
		fmt.Println("Ok, this process will be ready in a few seconds.")
		calc()

	case "no", "NO", "No":
		fmt.Println("Ok, you will be taken back to the sign in page.")
	default:
		fmt.Println("Invalid")
	}
}
func teach() {
	// This situation will lead the user through either a sign-up or sign-in situation
	fmt.Print("Would you like to sign up(u) or sign in(i): ")
	var decision string
	fmt.Scan(&decision)
	switch decision {
	case "i", "I":
		fmt.Println("You will be going through the sign-in process.")
		signIn()
	case "u", "U":
		fmt.Println("You will be going through the sign-up process.")
		signUp()
	}
	fmt.Print("Would you like to choose a student: ")
	var option string
	fmt.Scan(&option)
	switch option {
	case "yes", "YES", "Yes":
		student()
	case "no", "NO", "No":
		fmt.Println("Ok, you will be taken back to the sign in page.")
	default:
		fmt.Println("Invalid")
	}
	chosen()
}

// Student sign in function
func studt() {
	fmt.Print("Student Code: ")
	var semail string
	fmt.Scan(&semail)
	switch semail {
	case "1234":
		fmt.Println("Welcome, Student#1234")
		studtCheck()
	default:
		fmt.Println("Student not found!")
	}
}
func studtCheck() {
	fmt.Print("Would you like to check your grade: ")
	var stopt string
	fmt.Scan(&stopt)
	switch stopt {
	case "Yes", "YES", "yes":
		fmt.Println("Grades are pending. Check again later.")
	case "NO", "no", "No":
		fmt.Println("You will be taken back to the sign-in screen.")
	}
}

// Main function: Takes all results into account
func main() {
	fmt.Print("Are you a teacher or a student: ")
	var ts string
	fmt.Scan(&ts)
	switch ts {
	case "teacher", "TEACHER", "Teacher":
		fmt.Println("Welcome to mieaiGrade: Teacher Edition!")
		teach()
	case "student", "STUDENT", "Student":
		fmt.Println("Welcome to mieaiGrade: Student Edition!")
		studt()
	}
}
