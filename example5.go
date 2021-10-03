package main

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Printf("State %s country %s", a.state, a.country)
}

func main() {
	var interfaceVariable Describer
	p1 := Person{"Sri", 26}
	interfaceVariable = p1
	interfaceVariable.Describe()

	p2 := Address{"Vic", "Australia"}
	interfaceVariable = &p2
	interfaceVariable.Describe()

	var pointerVar Describer
	p3 := Address{"NSW", "Australia"}
	pointerVar = &p3
	pointerVar.Describe()
}
