/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

package main

import (
	"fmt"
	"strings"
)

// reverseWords переворачивает слова в строке
func reverseWords(s string) string {
	words := strings.Fields(s) // Разбиваем строку на слова
	n := len(words)

	// Перестановка слов местами
	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}

	return strings.Join(words, " ") // Собираем слова обратно в строку
}

func main() {
	str := "snow dog sun"
	reversed := reverseWords(str)
	fmt.Println("Оригинальная строка:", str)
	fmt.Println("Перевернутая строка:", reversed)
}
