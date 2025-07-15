// Assignment 1 :- Implement a Person struct in Go to represent individuals with attributes like name and age. Include methods for introducing themselves, updating their age, and checking if they are eligible to vote.

package main

import "fmt"

// Structure of the Person
type Person struct{
	Name string
	Age int
}

// Method to introduce the person
func (p Person) Introduce(){
	fmt.Printf("I am %s and my age is %d years\n", p.Name, p.Age)
}

// making a slice to store all the individual's Value
var individuals = []Person{
	{Name: "Aryan", Age: 21},
	{Name: "Atharv",Age: 22},
	{Name: "Ram",Age: 16},
}

// Method to add a new Individual in the slice
func addNewPerson() {
	var name string
	var age int
	fmt.Println("Enter the name of new individual:")
	fmt.Scanln(&name)
	fmt.Println("Enter his/her age:")
	fmt.Scanln(&age)

// Append is used here to add the new individuals to the slice
	individuals = append(individuals,Person{Name:name,Age:age})
}

// Method to update the Age of Person
func UpdateAge(individuals []Person) {
	var ageChange string
	var newAge int
	fmt.Println("Enter the Name of Individual you want to change age:")
	fmt.Scanln(&ageChange)
	fmt.Println("Enter the new age:")
	fmt.Scanln(&newAge)

	for person := range individuals{                                    // Iterating on the slice individuals to find name
		if individuals[person].Name == ageChange{
			individuals[person].Age = newAge
		}
	}
	fmt.Println("Individual name Invalid!!")
}

// Method to check the eligibility to vote
func (p Person) EligibleToVote() bool{                                // Here it will return true or false 
	return p.Age >= 18
}

// Method to Display everything
func ShowInformation(){
	for _,p := range individuals{
	p.Introduce()
	if p.EligibleToVote(){
		fmt.Println("Eligible to vote!!")
	} else {
		fmt.Println("Not Eligible to vote!!")
	}
}
} 
func main(){

// Showing the details of all individuals
ShowInformation()

// Asking if anyone wants to add a new individual or not
var addAPerson string
fmt.Println("Do you want to add a new Individual?")
fmt.Scanln(&addAPerson)
if addAPerson == "yes" || addAPerson == "Yes" || addAPerson =="Yes"{
	addNewPerson()
}

// Upadting the Age by the name
var change string
fmt.Println("Do you want to Change the Age of any Individual?")
fmt.Scanln(&change)
if change == "yes" || change == "Yes" || change =="Yes"{
	UpdateAge(individuals)
}

// Showing the details of all individual after Age Updation
ShowInformation()
}