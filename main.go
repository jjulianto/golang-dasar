package main

import (
	"bwagolang/calculation"
	"fmt"
)

func main() {
	var name string
	var grade string
	name = "Golang"
	score := 70
	age := 20
	age = 30
	number := 5
	fmt.Println("Hello, World!")

	result := calculation.Add(10, 10)
	resultTwo := calculation.Multiply(10, 10)

	fmt.Println(result, resultTwo, name, age)

	// If conditional
	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else {
		grade = "Kamu hebat"
	}

	fmt.Println(grade)

	// Switch conditional
	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("Default")
	}

	//Perulangan for
	for i := 1; i <= 5; i++ {
		fmt.Println("Saya sedang belajar Go:", i)
	}

	// While in Go
	i := 1
	for i <= 5 {
		fmt.Println("Saya sedang belajar Go ibu:", i)
		i++
	}

	//Foreach in Go
	title := "Golang dan Python"
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index :", index, " letter :", string(letter))
		}
	}

	for index, letter := range title {
		letterString := string(letter)

		// Using if conditional
		if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
			fmt.Println("index :", index, " letter :", string(letter))
		}

		// Using switch conditional
		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index :", index, " letter :", string(letter))
		}
	}
}
