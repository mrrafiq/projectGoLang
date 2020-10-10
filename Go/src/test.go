package main

import "fmt"

func main() {
	// var firstName string = "john"

	// var lastName string
	// lastName = "wick"
	// motor := "Honda"

	// fmt.Printf("halo %s %s!\n", firstName, lastName)
	// fmt.Print(motor)

	// var positivenumber = 70
	// var negativenumber = -90

	// fmt.Printf("Positif = %d \n", positivenumber)
	// fmt.Printf("Negative = %d \n", negativenumber)
	// fmt.Print(positivenumber)

	var mobil [4]string
	mobil[0]="Toyota"
	mobil[1]="Honda"
	mobil[2]="Suzuki"
	mobil[3]="Daihatsu"
	
	var a=0
	for i:=0; i<=3; i++ {
		fmt.Println(mobil[a])
		a++
	}
	
}
