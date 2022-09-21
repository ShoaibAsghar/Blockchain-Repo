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

func main() {
	//part 1
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"umar", 60000, "Data analyst"}
	emp3 := employee{"usman", 70000, "Software Engineer"}
	//part 2
	emplys := []employee{emp1, emp2, emp3}
	fmt.Println("The New empoyees Added ->: ", emplys)
	//part 3

	compani1 := company{"Tetra ->", emplys}

	fmt.Println("The Company details is ->", compani1)

}
