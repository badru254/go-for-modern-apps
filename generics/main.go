package main

import (
	"fmt"
)

//Generic interface
type Ordered interface {
	int | float64 | string
}

func min[T Ordered](values []T) (T, error) {

	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("min of empty slice")
	}

	m := values[0]

	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}

	return m, nil
}

func main() {

	strVal, _ := min([]string{"Pen", "Dog", "Cat", "Zebra", "Ant"})
	intVal, _ := min([]int{23, 55, 34, 2, 434, 6, 89, 56})
	floatVal, _ := min([]float64{23.5, 55.67, 34.45, 1.1, 434.3, 6, 89.6, 56.7, -9.56})

	fmt.Printf("Minimum of type %T is %v\n", strVal, strVal)
	fmt.Printf("Minimum of type %T is %v\n", intVal, intVal)
	fmt.Printf("Minimum of type %T is %v\n", floatVal, floatVal)

}
