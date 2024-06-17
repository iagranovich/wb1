package main

import (
	"fmt"
)

func ChangeBit(num int64, pos int) (int64, error) {
	if pos > 64 || pos <= 0 {
		return 0, fmt.Errorf("ошибка: номер бита должен быть от 1 до 64")
	}

	// сдвигаем бит на нужную позицию и создаем маску
	mask := int64(1) << (pos - 1)

	// XOR между числом и маской чтобы заменить нужный бит
	return num ^ mask, nil
}

func main() {
	for {
		fmt.Print("Введите число: ")
		var number int64
		fmt.Scan(&number)

		fmt.Print("Введите номер бита который хотите изменить: ")
		var position int
		fmt.Scan(&position)

		newNumber, err := ChangeBit(number, position)
		if err != nil {
			fmt.Println(err.Error() + "\n")
			continue
		}

		fmt.Printf("Вы ввели: %064b : %d\n", uint64(number), number)
		fmt.Printf("Результат: %064b : %d\n\n", uint64(newNumber), newNumber)
	}

}
