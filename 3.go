package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

var global string

type Student struct {
	rollnumber int
	name       string
	address    string
	subject    []string
}

func NewStudent(rollno int, name string, address string, subject []string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.subject = subject
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, subjects []string) *Student {

	s2 := strconv.Itoa(rollno)
	s3 := strings.Join(subjects, " ")

	global = s2 + s3 + address + name

	fmt.Println((s2))
	st := NewStudent(rollno, name, address, subjects)
	ls.list = append(ls.list, st)

	return st
}

func (ls *StudentList) Print() {
	for i := range ls.list {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student Rollno: %d\n", ls.list[i].rollnumber)
		fmt.Printf("Student Name: %s\n", ls.list[i].name)
		fmt.Printf("Student Address: %s\n\n", ls.list[i].address)
		fmt.Println("Student Subjects", ls.list[i].subject)

		sum := sha256.Sum256([]byte(global))
		fmt.Printf("Hash is == %x\n", sum) //hexadecimal

	}
}

func main() {
	student := new(StudentList)

	arr1 := []string{"English", "Urdu", "Maths"}
	arr2 := []string{"Computer Science", "Applied Physics", "Calcculul-1"}
	student.CreateStudent(2192, "Zakaullah", "Abbotabad", arr1)
	student.CreateStudent(1746, "Tehammi Bajwa", "Chicha Watni", arr2)
	student.Print()
}
