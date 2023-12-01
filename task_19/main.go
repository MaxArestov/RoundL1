/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

package main

import "fmt"

// ReverseString возвращает развернутый вариант входной строки.
func ReverseString(str string) string {
	runes := []rune(str) // Преобразовываем строку в слайс рун

	n := len(runes)

	// Меняем местами символы слайса с краев к центру.
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// Возвращаем приведенный обратно к строке массив байтов.
	return string(runes)
}

func main() {
	str := "главрыба" //Создаем строку.
	reverseStr := ReverseString(str)

	fmt.Printf("Строка - %s\nРазвернутая строка - %s", str, reverseStr) // Выводим развернутый вариант.
}
