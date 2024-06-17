package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	ss := []string{s1, s2, s3}

	for _, s := range ss {
		result := AllUnique(s)
		fmt.Println(result)

	}

}

func AllUnique(str string) bool {
	// для регистронезависимости
	lc := strings.ToLower(str)
	// для проверки уникальности
	h := make(map[rune]struct{})

	for _, s := range lc {

		// если ключ в мапе уже есть значит символ не уникальный
		if _, ok := h[s]; ok {
			return false
		}

		h[s] = struct{}{}
	}
	return true
}
