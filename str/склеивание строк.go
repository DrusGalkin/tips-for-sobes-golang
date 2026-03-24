package str

import "strings"

// Очень плохой вариант
func ConcatPlus(strs ...string) (str string) {
	for _, v := range strs {
		str += v // каждый раз создается новая строка, это очень плохо
	}
	return
}

func StrBuilder(strs ...string) string {
	var builder strings.Builder

	// у билдера есть свой буффер, нам слежует узнать точную длинну будущей строки
	// буфер такой же []byte

	var totalLength = 0
	for _, v := range strs {
		totalLength += len(v)
	}

	builder.Grow(totalLength) // устанавливаем длинну для оптимизации

	// склеиваем
	for _, v := range strs {
		builder.WriteString(v)
	}

	return builder.String()
}
