package main

import "fmt"

type Person struct {
	Name   string
	age    int
	job    string
	salary int
}

func main() {

	person1 := Person{Name: "Hege", age: 45, job: "Teacher", salary: 6000}

	argu_StructFunc(person1)

	person2 := Person{Name: "Cecilie", age: 24, job: "Marketing", salary: 4500}

	argu_StructFunc(person2)

}

//function that takes struct type as argument

func argu_StructFunc(obj Person) {
	fmt.Println("Name:   ", obj.Name)
	fmt.Println("Age:    ", obj.age)
	fmt.Println("Job:    ", obj.job)
	fmt.Println("Salary: ", obj.salary)

}
