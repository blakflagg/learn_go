package main

import (
	"fmt"
)

func main() {
	age := 32
	fmt.Println("Age:", age)
	agePtr := &age

	fmt.Println("age pointer:", agePtr)
	fmt.Println("age pointer value:", *agePtr)
	getAdultYears(&age)
	adultYears := age

	fmt.Println("Adult Years:", adultYears)

}

func getAdultYears(age *int) {
	//Direct mutation of the pointer variable
	*age = *age - 18
}
