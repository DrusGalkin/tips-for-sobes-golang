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

// Так называемый middleware для проверки функции на ошибку
func Protected(f func() error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ошибка выполнения функции:", err)
		}
	}()

	fmt.Println("старт функции")
	if err := f(); err != nil {
		panic(err)
	}
}
