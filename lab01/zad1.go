package main

import (
	"fmt"
	"math/rand"
)

func delta(a float32, b float32, c float32) {
	delta := b*b - 4*a*c
	fmt.Println("Delta: ", delta)
	if delta < 0 {
		fmt.Println("No real solutions")
	} else if delta == 0 {
		fmt.Println("One solution: ", -b/(2*a))
	} else {
		fmt.Println("Two solutions: ", (-b+delta)/(2*a), " and ", (-b-delta)/(2*a))
	}
}

func randomGuessingThreeTries() {
	number := rand.Intn(10)
	var guessedNumber int
	fmt.Println("Guess the number: ")
	fmt.Scanln(&guessedNumber)
	for i := 1; i < 3; i++ {
		if guessedNumber != number {
			fmt.Printf("You have %v more tries \n", 3-i)
			fmt.Println("Wrong number, try again: ")
			fmt.Scanln(&guessedNumber)
		} else {
			fmt.Println("You guessed the number!")
			return
		}
	}
	fmt.Println("Unlucky, maybe next time!")
}

func randomGuessing() {
	number := rand.Intn(10)
	var guessedNumber int
	fmt.Println("Guess the number: ")
	fmt.Scanln(&guessedNumber)
	for guessedNumber != number {
		if guessedNumber != number && guessedNumber < number {
			fmt.Println("Too low!")
			fmt.Scanln(&guessedNumber)
		} else if guessedNumber != number && guessedNumber > number {
			fmt.Println("Too high!")
			fmt.Scanln(&guessedNumber)
		}
	}
	fmt.Println("You guessed the number!")
}

func main() {
	var age int
	var marsMultiplier float32 = 0.53
	var venusMultiplier float32 = 1.5
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("Months lived: ", age*12)
	fmt.Println("Days lived: ", age*365)
	fmt.Println("Age if you were to live on Mars: ", float32(age) * marsMultiplier)
	fmt.Println("Age if you were to live on Venus: ", float32(age) * venusMultiplier)
	delta(3.0, 10.0, 1.0)
	randomGuessing()
}