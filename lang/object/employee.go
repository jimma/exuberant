package object

import (
	"fmt"
)

type Employee struct {
	//A field declared with a type but no explicit field name is
	//an anonymous field, also called an embedded field or an embedding
	//of the type in the struct. An embedded type must be specified as a type name
	//T or as a pointer to a non-interface type name *T, and T itself may not be a
	// pointer type. The unqualified type name acts as the field name.
	//so person's default field name is Person
	Person
	Company       string
	EmployeeID    int32
	PrintIDPerson PrintIDPerson
}

func (e Employee) GetEmployeeID() int32 {
	return e.EmployeeID
}

func (e *Employee) ChangeEmployeeID() {
	e.EmployeeID = 0
}

func (e *Employee) PrintID() {
	fmt.Printf("Employee's ID : %d \n", e.EmployeeID)
}
