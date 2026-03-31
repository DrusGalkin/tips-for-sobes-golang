package defers

import "fmt"

type data struct {
	value int
}

// Аргументы функции вычесляются в момент объявления defer
// -> defer fmt.Println(d.value) == 5 - 5ка положиться на стек, и в конце выполнения функции defer будет вызываться с данным значением
func DeferEvalutionAsUsual() {
	d := data{5}
	defer fmt.Println(d.value)

	d.value++
}

// А тут значение берется уже после инкрементации
func DeferEvalutionOnDeamand() {
	d := data{5}

	defer func() {
		fmt.Println(d.value)
	}()

	d.value++
}
