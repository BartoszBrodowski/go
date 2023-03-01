package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")

	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	// arr := []string{"a", "b", "c"}
	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}

	result, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}


}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
// * is a pointer to memory address
func inc(x *int) {
	*x++
}