package main

import "fmt"


func multiply (a, b int) int {
	return a * b
}

//Find if a number is odd



//Find if a password belongs to a user


func main(){

	//Triangle
	multiplicationResultTriangle := multiply(10, 10)
	finalResultTriangle := multiplicationResultTriangle / 2

	fmt.Println("The traingle area is:", finalResultTriangle)

	//Rectagle
	multiplicationResultRectangle := multiply(10, 20)
	fmt.Println("The rectangle area is:", multiplicationResultRectangle)

	//Trapeze
	multiplicationResultTrapeze := multiply(20, 30)
	finalResultTrapeze := multiply(multiplicationResultTrapeze, 10)
	fmt.Println("The trapeze area is:", finalResultTrapeze)

	name := "I'm Go" 


	fmt.Println("Hello, World!", name)
}