package main

import "fmt"
import "strings"

func vowelCounter(phrase string)(int, int, int, int, int,){
	conta := 0
	conte := 0
	conti := 0
	conto := 0
	contu := 0
	for _ , value := range phrase{ 
		variable := strings.ToLower(string(value))
		switch variable {
		case "a" :
			conta++
		case "e" :
			conte++
		case "i" :
			conti++
		case "o" :
			conto++
		case "u" :
			contu++
		}
	}
	return conta, conte , conti, conto, contu
}

func main() {

	phrase := "Hi, this is my first GO program! Hope you like it"	
	a,e,i,o,u := vowelCounter(phrase)
	fmt.Printf("The phrase '%s' contains: \n",phrase)
	fmt.Printf("%d vowels a \n",a)
	fmt.Printf("%d vowels e \n",e)
	fmt.Printf("%d vowels i \n",i)
	fmt.Printf("%d vowels o \n",o)
	fmt.Printf("%d vowels u \n",u)	
}
