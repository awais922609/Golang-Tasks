package main

import "fmt"

type person struct {
	name   string
	age    int
	job    string
	salary int
}

func PrintingValues(pers1 person) {
	fmt.Println("Name is == ", pers1.name, "\n")
	fmt.Println("Age is ==  ", pers1.age, "\n")
	fmt.Println("Job is ==  ", pers1.job, "\n")
	fmt.Println("Salary is ==  ", pers1.salary, "\n")

}

func main() {

	var pers1 person
	var pers2 person

	pers1.name = "Hege"
	pers1.job = "Teacher"
	pers1.age = 45
	pers1.salary = 6000

	pers2.name = "Cecile"
	pers2.job = "Marketing"
	pers2.salary = 4500
	pers2.age = 24

	PrintingValues(pers1)

	PrintingValues(pers2)

}
