package main

import (
	"fmt"
	"strings"
)

// воспользуемся пакетом strings
func WordsRevers(str string) string {
	// разбиваем строку на слова
	w := strings.Fields(str)

	// меняем слова местами
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}

	// соединяем слова назад через пробел
	return strings.Join(w, " ")
}

func main() {

	str := "snow dog sun"

	fmt.Println(WordsRevers(str))

}
