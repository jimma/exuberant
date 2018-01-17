package object

import "fmt"

type Person struct {
	Name   string
	Gender string
	Id     int32
}

//The rule about pointers vs. values for receivers is that value methods can be
//invoked on pointers and values, but pointer methods can only be invoked on
//pointers
func (p *Person) GetName() string {
	return p.Name
}

//The first is so that the method can modify the value that its receiver points to.
//The second is to avoid copying the value on each method call. This can be more
//efficient if the receiver is a large struct
func (p *Person) ChangeName() string {
	p.Name = "Changed"
	return p.Name
}

//while methods with value receivers take either a value or a pointer as
//the receiver when they are called but if p is a large struct , don't do this
func (p Person) NotReallyChangeName() string {
	p.Name = "Changed"
	return p.Name
}

func (p *Person) PrintID() {
	fmt.Println("Person's ID")
}

type CommonPerson interface {
	PrintID()
}

//CallPersonID will print a person id
func PrintPersonID(commonPerson CommonPerson) {
	commonPerson.PrintID()
	//fmt.Println("Not really change name: " + commonPerson.PrintID())
}
