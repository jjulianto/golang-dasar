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

	//Array
	var languages [5]string
	languages[0] = "Go"
	languages[1] = "Phyton"
	languages[2] = "Java"
	languages[3] = "C++"
	languages[4] = "JavaScript"

	fmt.Println(languages)
	length := len(languages)
	fmt.Println(length)

	languagess := [...]string{
		"Ruby",
		"Go",
		"Java",
		"C#",
		"JavaScript",
		"C++",
	}

	for index, lang := range languagess {
		fmt.Println("Index : ", index, " Language :", lang)
	}

	//Slice
	var gamingConsoles []string
	gamingConsoles = append(gamingConsoles, "Playstation 4")
	gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	gamingConsoles = append(gamingConsoles, "Xbox One")

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}

	//Map
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["Ruby"] = 9
	myMap["Java"] = 7
	myMap["Kotlin"] = 8

	fmt.Println(myMap)
	fmt.Println(myMap["Ruby"])

	myMaps := map[string]string{
		"Go": "Niceeee",
		"C#": "Gooodd",
	}

	fmt.Println(myMaps)

	mapMap := map[string]string{
		"ruby":       "Very gooodd",
		"go":         "goooodddd",
		"javascript": "hype",
	}

	for key, value := range mapMap {
		fmt.Println("Key : ", key, " value : ", value)
	}

	fmt.Println("==============")

	delete(mapMap, "ruby")

	for key, value := range mapMap {
		fmt.Println("Key : ", key, " value : ", value)
	}

	value, isAvailable := mapMap["python"]
	fmt.Println(isAvailable)
	fmt.Println(value)

	//Slice of Map
	students := []map[string]string{
		{"name": "Agus", "score": "A"},
		{"name": "Nugroho", "score": "B"},
		{"name": "Rahmawan", "score": "C"},
	}

	for _, student := range students {
		fmt.Println(student["name"], " - ", student["score"])
	}

	//Quiz
	scores := [8]int{100, 88, 75, 92, 70, 93, 85, 67}

	var total int

	for _, score := range scores {
		total += score
	}

	lengthh := len(scores)
	average := float64(total) / float64(lengthh)

	fmt.Println(average)

	var goodScores []int

	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println(goodScores)
}
