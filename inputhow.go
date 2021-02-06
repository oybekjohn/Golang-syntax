package main 

import "fmt"

func main() { 

	// Println function is used to 
	// display output in the next line 
	fmt.Println("Enter Your First Name: ") 

	// var then variable name then variable type 
	var Ism string 

	// Taking input from user 
	fmt.Scanln(&Ism) 
	fmt.Println("Enter Second Last Name: ") 
	var Family string 
	fmt.Scanln(&Family) 

	// Print function is used to 
	// display output in the same line 
	fmt.Print("Your Full Name is: ") 

	// Addition of two string 
	fmt.Print(Ism + " " + Family) 
} 