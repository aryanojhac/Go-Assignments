// Assignment 1 :- Implement a Person struct in Go to represent individuals with attributes like name and age. Include methods for introducing themselves, updating their age, and checking if they are eligible to vote.

package main
import "fmt"

// Structure of the Person
type Person struct{
	name string
	age int
}

// Method to introduce the person
func (p Person) Introduce(){
	fmt.Printf("I am %s and I am %d years old\n", p.name, p.age)
}

// Method to update the Age of Person
func (p *Person) UpdateAge(newAge int){
	p.age = newAge
}

// Method to check the eligibility to vote
func (p Person) EligibleToVote() bool{                                // Here it will return true or false 
	return p.age >= 18
}

// Main Function to pass all the values 
func main(){

// creating a new Person
	person1 := Person{"Aryan",21}

// Introducing the Person
	person1.Introduce()

// Checking the eligibity of the Person to vote
	fmt.Println("Is he eligible to Vote?")

// Condition to check the eligibility to vote 
	if person1.EligibleToVote(){
		fmt.Println("He is Eligible to vote")
	} else {
		fmt.Println("He is not Eligible to Vote")
	}

// Now updating the age to check the method 
	person1.UpdateAge(17)

// After the above call I want to again check the age so I will again use Introduce method
	person1.Introduce()

// Again checking the eligibility to vote
	fmt.Println("Is he eligible to Vote?")

// We can use this way as well to check the eligibility
	if person1.EligibleToVote(){
		fmt.Println("He is Eligible to vote")
	} else {
		fmt.Println("He is not Eligible to Vote")
	}
}