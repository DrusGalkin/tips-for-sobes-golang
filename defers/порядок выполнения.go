package defers

import "fmt"

// Порядок выполнения с самого последнего defer -> 3 2 1
func ПорядокВыполнения() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
