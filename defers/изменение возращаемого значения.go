package defers

import "fmt"

// Вернется 9, а не 4
func ИзменениеВозращаемогоЗначения() (num int) {
	num = 4

	defer func() {
		num = 9
	}()

	num++
	return
}

// defer в любом случае выполняется
func PanicWhisDefer() {
	defer func() {
		fmt.Println("defer2")
	}()

	ff()
}

// panic с defer
func ff() {
	defer func() {
		fmt.Println("defer1")
	}()
	panic("kaappp")
}
