// Assignment 2:- Develop a system to manage employees and their departments. Each employee should have a name, age, and salary. Each department should have a name, a list of employees, and a method to calculate the average salary of its employees. Additionally, implement methods to add and remove employees from departments and to give a raise to an employee.

package main

import "fmt"

// Making the Structure of Employee having attibutes as Name, Age and Salary
type Employee struct {
	Name   string
	Age    int
	Salary float64
}

// Making the structure of Department having attributes as Dept Name and Employee struct
type Department struct {
	DeptName  string
	Employees map[string]*Employee
}

// Method to calculte the Vaerage Salary of the Employee
func (d Department) AverageSalary() float64 {
	total := 0.0
	for _, emp := range d.Employees {
		total += emp.Salary
	}
	// Returning the float value for the average salary
	return float64(total) / float64(len(d.Employees))
}

// Method to add a Employee
func (d *Department) addEmployee(e Employee) {
	d.Employees[e.Name] = &e
}

// Method to display all the Employees
func (d Department) ListAll() {
	for _, emp := range d.Employees { // Iterating on all employees to display the details
		fmt.Printf("%-10s | %2d |  ₹%.2f\n",
			emp.Name, emp.Age, emp.Salary)
	}
}

// Method to add a Employee by user
func (d *Department) addEmployeeByUser() {
	var newEmployee string
	fmt.Println("DO you want to add new Employee?")
	fmt.Scanln(&newEmployee)

	if newEmployee == "yes" || newEmployee == "Yes" || newEmployee == "Y" || newEmployee == "y" {
		var name string
		var age int
		var salary float64

		fmt.Print("Enter the Name of Employee:")
		fmt.Scanln(&name)

		fmt.Print("Enter the Age of Employee:")
		fmt.Scanln(&age)

		fmt.Print("Enter the salary of Employee: ")
		fmt.Scanln(&salary)

		// Adding the employee to the map
		d.addEmployee(Employee{Name: name, Age: age, Salary: salary})
	}
}

// Method to remove an Employee
func (d *Department) RemoveEmployee(name string) {
	delete(d.Employees, name)
}

// Method to give a raise to an Employee
func (d *Department) toGiveRaise(name string, amount float64) {
	fmt.Println("After Upraisal")
	emp := d.Employees[name]
	emp.Salary += amount
	fmt.Println("The Raise is given")
}

func main() {
	dept := Department{
		DeptName:  "IT",
		Employees: make(map[string]*Employee),
	}

	// Adding employees manually (Hardcoded)
	dept.addEmployee(Employee{"Aryan", 21, 11000})
	dept.addEmployee(Employee{"Atharv", 23, 11000})
	dept.addEmployee(Employee{"Purva", 25, 75000})
	dept.addEmployee(Employee{"Ram", 34, 90000})

	// Listing all the Employees
	dept.ListAll()

	// Printing average salary of Employees
	fmt.Printf("\nAverage salary in %s: ₹%.2f\n", dept.DeptName, dept.AverageSalary())

	// Adding the new Employee
	dept.addEmployeeByUser()

	// Giving Raise to the employee
	dept.toGiveRaise("Aryan", 1000)

	// Updated New Average
	fmt.Printf("\nAverage salary in %s: ₹%.2f\n", dept.DeptName, dept.AverageSalary())
	// Printing updated list of employees
	dept.ListAll()
}
