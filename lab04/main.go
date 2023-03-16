package main

import "fmt"

func vectorSum() [20]int64 {
	var wektor1 [20]int64
	var wektor2 [20]int64
	var result [20]int64
	for i := 0; i < 20; i++ {
		wektor1[i] = 2
	}
	for i := 0; i < 20; i++ {
		wektor2[i] = 3
	}
	for i := 0; i < 20; i++ {
		result[i] = wektor1[i] + wektor2[i]
	}
	return result
}

func hadamardsFormula(tab1 [][]int, tab2 [][]int) [][]int {
	var result [][]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = 1
		}
	}
	fmt.Println(result)
	return result
}

func goodHadamardsFormula(tab1 [][]int, tab2 [][]int) [][]int {
	if len(tab1[0]) != len(tab2[0]) || len(tab1) != len(tab2) {
		fmt.Println("Arrays must be the same size")
	}
	c := make([][]int, 3)
	for i := range c {
		c[i] = make([]int, len(tab1[0]))
		for j := range c[i] {
			c[i][j] = tab1[i][j] * tab2[i][j]
		}
	}
	return c
}

func matrixPrint(matrix [][]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j+1 % 3 == 3 {
				fmt.Println(matrix[i][j])
			} else {
				fmt.Print(matrix[i][j], " ")
			}
		}
	}
}

func main() {
	sequence := vectorSum()
	fmt.Println(sequence)
	tablica1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	tablica2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := goodHadamardsFormula(tablica1, tablica2)
	fmt.Println(result)
	matrixPrint([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}