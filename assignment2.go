// Assignment 2:- Develop a system to manage employees and their departments. Each employee should have a name, age, and salary. Each department should have a name, a list of employees, and a method to calculate the average salary of its employees. Additionally, implement methods to add and remove employees from departments and to give a raise to an employee.

package main

import (
	"errors"
	"fmt"
)

// Making the Structure of Employee having attibutes as Name, Age and Salary
type employee struct {
	name   string
	age    uint
	salary float64
}

// Making the structure of Department having attributes as Dept Name and Employee struct
type department struct {
	deptName  string
	employees []employee
}

// Method to calculte the Vaerage Salary of the Employee
func (d department) averageSalary() (float64, error) {
	if len(d.employees) == 0 {
		return 0, errors.New("The department is empty so cannot calculate the salary")
	}

	total := 0.0
	for _, emp := range d.employees {
		total += emp.salary
	}
	// Returning the float value for the average salary
	return float64(total) / float64(len(d.employees)), nil
}

// Method to add a Employee
func (d department) addEmployee(e employee) department {
	d.employees = append(d.employees, e)
	return d
}

// Method to display all the Employees
func (d department) listAll() {
	for _, emp := range d.employees { // Iterating on all employees to display the details
		fmt.Printf("%-10s | %2d |  ₹%.2f\n",
			emp.name, emp.age, emp.salary)
	}
}

// Method to add a Employee by user
func (d department) addEmployeeByUser() (department, error) {
	var newEmployee string
	fmt.Println("DO you want to add new Employee?")
	_, err := fmt.Scanln(&newEmployee)
	if err != nil {
		return d, fmt.Errorf("Invalid Input", err)
	}

	if newEmployee == "yes" || newEmployee == "Yes" || newEmployee == "Y" || newEmployee == "y" {
		var name string
		var age uint
		var salary float64

		fmt.Print("Enter the Name of Employee:")
		fmt.Scanln(&name)

		fmt.Print("Enter the Age of Employee:")
		_, err := fmt.Scanln(&age)
		if err != nil {
			return department{}, fmt.Errorf("Invalid Age: %w", err)
		}

		fmt.Print("Enter the salary of Employee: ")
		_, err = fmt.Scanln(&salary)
		if err != nil {
			return department{}, fmt.Errorf("Invalid Salary: %w", err)
		}
		// Adding the employee to the map
		newEmployee := employee{name: name, age: age, salary: salary}
		return d.addEmployee(newEmployee), nil
	}
	return d, errors.New("no new employee to add")
}

// Method to remove an Employee
func (d department) removeEmployee(name string) (department, error) {
	for i, emp := range d.employees {
		if emp.name == name {
			d.employees = append(d.employees[:i], d.employees[i+1:]...)
			fmt.Printf("Employee '%s' has been removed.\n", name)
			return d, nil
		}
	}
	return d, fmt.Errorf("Employee with name %s does not exist", name)
}

// Method to give a raise to an Employee
func (d department) toGiveRaise(name string, amount float64) (department, error) {
	fmt.Println("After Upraisal")
	for i, emp := range d.employees {
		if emp.name == name {
			emp.salary += amount
			d.employees[i] = emp
			fmt.Println("The Raise is given to", emp.name)
			return d, nil
		}
	}
	return d, fmt.Errorf("The salary of %s has been raised", name)
}

func main() {
	dept := department{
		deptName:  "IT",
		employees: []employee{},
	}

	// Adding employees manually (Hardcoded)
	dept = dept.addEmployee(employee{"Aryan", 21, 11000})
	dept = dept.addEmployee(employee{"Atharv", 23, 11000})
	dept = dept.addEmployee(employee{"Purva", 25, 75000})
	dept = dept.addEmployee(employee{"Ram", 34, 90000})

	// Listing all the Employees
	fmt.Println("The first list of Employees:")
	dept.listAll()

	// Printing average salary of Employees
	avgSalary, err := dept.averageSalary()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nAverage salary in %s: ₹%.2f\n", dept.deptName, avgSalary)
	}

	// Create and add a new employee from user input
	dept, err = dept.addEmployeeByUser()
	if err != nil && err.Error() != "no new employee to add" {
		fmt.Println("Error adding employee:", err)
	}

	// Giving Raise to the employee
	dept, err = dept.toGiveRaise("Aryan", 1000)
	if err != nil {
		fmt.Println("\nThe raise cannot be given", err)
	}

	// To remove a employee
	dept, err = dept.removeEmployee("Ram")
	if err != nil {
		fmt.Println("Invalid Employee:", err)
	}

	// Updated New average
	fmt.Printf("\nAverage salary in %s: ₹%.2f\n", dept.deptName, avgSalary)
	// Printing updated list of employees
	dept.listAll()
}
