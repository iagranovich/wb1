package main

import (
	"fmt"
	"math/big"
)

func Summa(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func Raznost(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func Umnozhenie(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

func Delenie(x, y *big.Int) *big.Int {
	return new(big.Int).Quo(x, y)
}

func main() {

	a, _ := new(big.Int).SetString("-3354483313135354313135", 10)
	b, _ := new(big.Int).SetString("558461313431351354333", 10)

	fmt.Println(Summa(a, b))
	fmt.Println(Raznost(a, b))
	fmt.Println(Umnozhenie(a, b))
	fmt.Println(Delenie(a, b))
}
