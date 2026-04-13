package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age:", age)
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)

	var agePointer *int = &age
	//agePointer = &age
	fmt.Println("AgePointer:", *agePointer)
	adultYearsPointer := getAdultYearsPointer(agePointer)
	fmt.Println(adultYearsPointer)
}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsPointer(agePointer *int) int {

	return *agePointer - 18
}
