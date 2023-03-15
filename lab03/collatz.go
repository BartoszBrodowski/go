package main

import "fmt"

func collatz(number int) []int {
	sequence := []int{number}
	for number != 1 {
		if number%2 == 0 {
			number = number / 2
		} else {
			number = 3*number + 1
		}
		sequence = append(sequence, number)
	}
	fmt.Println("Length of sequence: ", len(sequence))
	return sequence
}

func main() {
	fmt.Print("Enter a number: ")
	var number int
	fmt.Scanln(&number)
	collatz_sequence := collatz(number)
	fmt.Println(collatz_sequence)
}