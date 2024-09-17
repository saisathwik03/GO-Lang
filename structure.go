package main
import "fmt"

type Person struct{
	name string
	age int
	job string
	salary int
}

func main(){
	var p1 Person
	var p2 Person

	p1.name = "Sathwik"
	p1.age = 22
	p1.job = "Student"
	p1.salary = 0

	p2.name = "Lalitha"
	p2.age = 21
	p2.job = "Teacher"
	p2.salary = 150000

	printPerson(p1)
	printPerson(p2)
}

func printPerson(p Person){
	fmt.Println("Name:",p.name)
	fmt.Println("Age:",p.age)
	fmt.Println("Job:",p.job)
	fmt.Println("Salary:",p.salary)
}