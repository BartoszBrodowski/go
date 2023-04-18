package main

import (
	"fmt"
	"flag"
)

func main() {
	var numbersStr string
	flag.StringVar(&numbersStr, "numbers", "", "numbers to be added")
	flag.Parse()

	numbers := strings.Split(numbersStr, ",")

	var numbersArr []int
	for 
}