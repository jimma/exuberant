package main

import (
	"fmt"

	"github.com/jimma/exuberant/lang/object"
)

func PrintPersonIDToConsole(in interface{}) {
	pidPerson, ok := in.(object.PrintIDPerson)
	if ok {
		pidPerson.PrintID()
	} else {
		fmt.Printf("Input : %v isn't a PrintIDPerson\n", in)
	}
}

func PrintPersonID(pidPerson object.PrintIDPerson) {
	pidPerson.PrintID()
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	person := object.Person{"John", "male", 100010}
	employee := object.Employee{person, "MS", 10001, person}

	fmt.Printf("Employee's name is: %s \n", employee.GetName())
	//pointer receiver
	employee.PrintIDPerson.PrintID()

	employee.NotReallyChangeName()
	fmt.Printf("Value receive won't change name: %s \n", employee.GetName())
	employee.ChangeName()
	fmt.Printf("Pointer receive can actulaly change name: %s \n", employee.GetName())

	//a pointer type may call the methods of its associated value type,
	//but not vice verse
	//interface type needs a pointer receive, so we need to pass the pointer value

	PrintPersonID(&person)
	PrintPersonID(person)
	person.PrintID()
	employee.PrintID()
	(&employee).PrintID()
	//complie error PrintPersonID(employee)
	//object.PrintPersonID(employee)

	PrintPersonIDToConsole(person)
	PrintPersonIDToConsole(&person)
	PrintPersonIDToConsole(employee)
	PrintPersonIDToConsole(&employee)
}
