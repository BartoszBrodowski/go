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

func addTwoDimensionalArrayAndItsReverse(arr1 [][]int) {
	arr2 := make([][]int, len(arr1))
	result := make([][]int, len(arr1))
	for i := range arr2 {
		arr2[i] = make([]int, len(arr1[0]))
		result[i] = make([]int, len(arr1[0]))
	}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[0]); j++ {
			arr2[i][j] = arr1[len(arr1) - 1 - i][len(arr1[0]) - 1 - j]
		}
	}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[0]); j++ {
			result[i][j] = arr1[i][j] + arr2[i][j]
		}
	}

	fmt.Println(result)
}

func multiplyMatrixes(matrix1 [][]int, matrix2 [][]int) {
	if len(matrix1[0]) == len(matrix2) || len(matrix1) == len(matrix2[0]) {
		result := make([][]int, len(matrix1))
		for i := range result {
			result[i] = make([]int, len(matrix2[0]))
		}
		for i := 0; i < len(matrix1); i++ {
			for j := 0; j < len(matrix2[0]); j++ {
				for k := 0; k < len(matrix2); k++ {
					result[i][j] += matrix1[i][k] * matrix2[k][j]
				}
			}
		}
		fmt.Println(result)
	} else {
		fmt.Println("Matrixes can't be multiplied")
	}
}

func main() {
	// sequence := vectorSum()
	// fmt.Println(sequence)
	tablica1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(tablica1[2][2])
	tablica2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// result := goodHadamardsFormula(tablica1, tablica2)
	// fmt.Println(result)
	// matrixPrint([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	// addTwoDimensionalArrayAndItsReverse(tablica1)
	multiplyMatrixes(tablica1, tablica2)
}