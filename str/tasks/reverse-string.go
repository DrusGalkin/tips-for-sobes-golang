package tasks

// Напишите функцию которая переворачивает строку на месте
// Пример: привет -> тевирп
func ReverseString(str string) string {
	// конвертируем строку в массив рун, для корректной работы с Unicode
	// runes := ['п', 'р', 'и', 'в', 'е', 'т']
	// индексы :  0    1    2    3    4    5

	runes := []rune(str)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
