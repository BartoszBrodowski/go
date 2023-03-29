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
	return sequence
}

func biggestCollatzSequence(num1 int, num2 int) (int, int) {
	max := len(collatz(10))
	var max_number int 
	for i := num1; i <= num2; i++ {
		collatz_sequence := collatz(i)
		if len(collatz_sequence) > max {
			max = len(collatz_sequence)
			max_number = i
		}
	}
	return max_number, max
}

func main() {
	// fmt.Print("Enter a number: ")
	// var number int
	// fmt.Scanln(&number)
	// collatz_sequence := collatz(number)
	// fmt.Println(collatz_sequence)
	max_number, max_length := biggestCollatzSequence(2000, 3000)
	fmt.Println("Longest sequence comes out of number: ", max_number, " with length: ", max_length)
}