package main

import (
	"fmt"
	"math"
)

func main() {
	switchStatement()
}

func switchStatement() {
	name := "Luke"
	switch name {
	case "Matthew":
		fmt.Println("First Gospel")
	case "Mark":
		fmt.Println("Second Gospel")
	case "Luke":
		fmt.Println("Third Gospel")
	case "John":
		fmt.Println(("Fourth Gospel"))
	}
}

func ifAndElse() {
	if grade := 1111; grade < 60 {
		fmt.Println("Grade is F")
	} else if grade < 70 {
		fmt.Println("Grade is D")
	} else if grade < 80 {
		fmt.Println("Grade is C")
	} else if grade < 90 {
		fmt.Println("Grade is B")
	} else {
		fmt.Println("Excellent, grade is A")
	}

}

func forLoop() {
	var min, max int = 0, 5
	fmt.Println("Counting up to", max)

	for min <= max {
		fmt.Println(min)
		min = min + 1
	}

	fmt.Println("")

	for i := 0; i < 4; i++ {
		fmt.Println("i: ", i)
	}

	fmt.Println("")
	for j := range 4 {
		fmt.Println("Range: ", j)
	}
}

func variables() {
	const temp = -45
	var firstName = "James"
	lastName := "Leveille"

	var fullName = fmt.Sprintf("%s %s", firstName, lastName)

	var isAbsent = false
	if isAbsent {
		fmt.Println("Absent")
	}
	fmt.Println(fullName)
	fmt.Println(math.Abs(temp))
}

func helloWorld() {
	fmt.Println("Hello, World!")

	fmt.Println(4 + 5)

	fmt.Println("Concate" + "nate")
}
