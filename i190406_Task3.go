package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (Is *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	Is.list = append(Is.list, st)
	return st
}

/*
func (student *StudentList) printfunc(Student StudentList) *Student {
	fmt.Println(student.list[0].rollnumber)
	fmt.Println(student.list[0].name)
	fmt.Println(student.list[0].address)
}
*/

func main() {

	student := new(StudentList)

	student.CreateStudent(24, "Asim", "AAAAAA")
	// printfunc(student)
	fmt.Println("\n", strings.Repeat("=", 25), "List 0 ", strings.Repeat("=", 25))
	fmt.Println("student rollno    ", student.list[0].rollnumber)
	fmt.Println("student name      ", student.list[0].name)
	fmt.Println("student address   ", student.list[0].address)

	student.CreateStudent(25, "Naveed", "BBBBBB")
	fmt.Println("\n", strings.Repeat("=", 25), "List 1 ", strings.Repeat("=", 25))
	fmt.Println("student rollno    ", student.list[1].rollnumber)
	fmt.Println("student name      ", student.list[1].name)
	fmt.Println("student address   ", student.list[1].address)

}
