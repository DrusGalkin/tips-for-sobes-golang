package defers

import "fmt"

// обработка panic через recovery
func ProblemFunc() (num int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Все ок")
			num = 3
			return
		}
	}()

	num = 0

	if num <= 0 {
		panic("что-то не то")
	}
	return
}

func Protected(f func() error) bool {
	defer func() {

	}()
}
