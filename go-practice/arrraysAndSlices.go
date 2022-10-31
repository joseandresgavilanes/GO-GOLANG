package main

import "fmt"

func main() {

	//ARRAYS
	var array [4]int // arrays are immutable in Go, you cannot add or delete elements
	fmt.Println(array)
	array[0] = 1
	array[1] = 1

	//SLICE
	//Basic Declaration
	{
		var _ = []int{} //Its length is automatically 0
		
		//Creating a slice with a specific length and capacity
		var x = make([]float64, 5, 10)
		fmt.Println(len(x))
		fmt.Println(cap(x))

	}

	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	//METHODS OF SLICE
	fmt.Println(slice[0])
	fmt.Println(slice[:5])  //elements from the beginning to the element 5
	fmt.Println(slice[2:5]) //elements between the exclusive interval [2,5]
	fmt.Println(slice[2:])  //elements from 2 to the end

	//LEN FUNCTION
	fmt.Println(len(slice)) //Returns the length of the structure
	fmt.Println(len(array))

	//CAP FUNCTION
	fmt.Println(cap(slice)) //Returns the capacity of elements of the structure
	fmt.Println(cap(array))

	//APPEND FUNCTION
	slice = append(slice, 10) //It receives two arguments; first the slice that it is going to modify,
	slice = append(slice, 0)  //then the element that it will receive
	slice = append(slice, 7)
	fmt.Println(slice)

	//APPEND A NEW LIST
	var slice2 = []int{8, 9, 11}
	slice = append(slice, slice2...) //The spread operator references every one of the elements in the called slice
	fmt.Println(slice)

	//DIFFERENCES BETWEEN SLICE AND ARRAY
	//ARRAY
	//-	Immutable
	//-	More effience
	//- Length must be specified when declared
	//SLICE
	//-	Mutable
	//- Length can vary and it could not be specified when declared
}
