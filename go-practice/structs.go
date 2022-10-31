package main

import "fmt"

//STRUCTS
type person struct {
	name     string
	lastname string
	age      int
	cristian bool
}

func main() {
	//Instancing
	me := person{name: "David", lastname: "Olivos", age: 15, cristian: true}
	fmt.Println(me)

	//Another way of instancing
	var otherPerson person
	otherPerson.name = "Jonathan"
	otherPerson.lastname = "Ramirez"
	fmt.Println(otherPerson)

}