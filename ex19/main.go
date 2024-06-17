package main

import "fmt"

// для разделения на символы воспользуемся рунами
func StringRevers(s string) string {
	r := []rune(s)

	// разворачиваем руны в обратном порядке
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	// приводим к строке
	return string(r)
}

func main() {

	str := "главрыба!"

	fmt.Println(StringRevers(str))

}
