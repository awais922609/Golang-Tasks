package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func AddEmployee(emp employee) {

}

func main() {

	var emp1 employee
	var emp2 employee
	var emp3 employee

	emp1.name = "Awais"
	emp1.position = "Full Stack Developer"
	emp1.salary = 80000

	emp2.name = "Rashid"
	emp2.position = "Coder"
	emp2.salary = 50000

	emp3.name = "Zakii"
	emp3.position = "MERN Stack Developer"
	emp3.salary = 80000

	emplys := []employee{emp1, emp2, emp3}

	var comp company

	comp.companyName = "Tesla"
	comp.employees = emplys

	fmt.Println("\n")
	fmt.Println("Company Name is == ", comp.companyName)
	fmt.Println("Employees are == ", comp.employees)

}
