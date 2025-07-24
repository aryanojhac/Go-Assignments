// Assignment 1 :- Implement a Person struct in Go to represent individuals with attributes like name and age. Include methods for introducing themselves, updating their age, and checking if they are eligible to vote.

package main

import "fmt"

// Structure of the Person
type person struct {
	name string
	age  uint
}

// Method to introduce the person
func (p person) introduce() {
	fmt.Printf("I am %s and my age is %d years\n", p.name, p.age)
}

// making a slice to store all the individual's Value
var individuals = []person{
	{name: "Aryan", age: 21},
	{name: "Atharv", age: 22},
	{name: "Ram", age: 16},
}

// Method to add a new Individual in the slice
func addNewPerson() {
	var name string
	var age uint
	fmt.Println("Enter the name of new individual:")
	fmt.Scanln(&name)
	fmt.Println("Enter his/her age:")
	fmt.Scanln(&age)

	// Append is used here to add the new individuals to the slice
	individuals = append(individuals, person{name: name, age: age})
	return
}

// Method to update the age of Person by using his name
func updateAge(individuals []person) {
	var ageChange string
	var newage uint
	fmt.Println("Enter the name of Individual you want to change age:")
	fmt.Scanln(&ageChange)
	fmt.Println("Enter the new age:")
	fmt.Scanln(&newage)

	for person := range individuals { // Iterating on the slice individuals to find name
		if individuals[person].name == ageChange {
			individuals[person].age = newage
		}
	}
	fmt.Println("Individual name Invalid!!")
}

// Method to check the eligibility to vote
func (p person) eligibleToVote() bool { // Here it will return true or false
	return p.age >= 18
}

// Method to Display everything and to tell if one is eligible to vote or not
func showInformation() {
	for _, p := range individuals {
		p.introduce()
		if p.eligibleToVote() {
			fmt.Println("eligible to vote!!")
		} else {
			fmt.Println("Not eligible to vote!!")
		}
	}
}

func main() {

	// showing that the code executed successfully
	fmt.Println("Information gathered Successfully")

	// Asking if anyone wants to add a new individual or not
	var addAPerson string
	fmt.Println("Do you want to add a new Individual?")
	fmt.Scanln(&addAPerson)
	if addAPerson == "yes" || addAPerson == "Yes" || addAPerson == "Yes" {
		addNewPerson()
	}

	// Upadting the age by the name
	var change string
	fmt.Println("Do you want to Change the age of any Individual?")
	fmt.Scanln(&change)
	if change == "yes" || change == "Yes" || change == "Yes" {
		updateAge(individuals)
	}

	// showing the details of all individual after age Updation
	showInformation()
}
