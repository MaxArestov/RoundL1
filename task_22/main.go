/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Использование big.Int для поддержки очень больших чисел
	a := big.NewInt(1)
	b := big.NewInt(1)

	a.Lsh(a, 91) // 2^91
	b.Lsh(b, 92) // 2^92
	fmt.Printf("a=%d\nb=%d\n", a, b)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сумма: %s\n", sum)

	// Вычитание
	diff := new(big.Int).Sub(b, a)
	fmt.Printf("Разность: %s\n", diff)

	// Умножение
	product := new(big.Int).Mul(a, b)
	fmt.Printf("Произведение: %s\n", product)

	// Деление
	quotient := new(big.Int).Div(b, a)
	fmt.Printf("Частное: %s\n", quotient)
}
