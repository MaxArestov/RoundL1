/*
Разработать программу, которая проверяет, что все символы в строке - уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

package main

import (
	"fmt"
	"strings"
)

// IsUniqueChars проверяет строку на уникальность символов.
func IsUniqueChars(str string) bool {
	uniqueChars := make(map[rune]bool) // Создаем карту.

	for _, val := range strings.ToLower(str) { // Проходимся по рунам строки, приведенным к нижнему регистру.
		if _, exists := uniqueChars[val]; exists {
			return false // Если такой символ уже есть - возвращаем false.
		}
		uniqueChars[val] = true // Добавляем в карту руну.
	}
	return true // Если после полного цикла не вернулся false, возвращаем true.
}

func main() {
	var str string

	fmt.Print("Введите строку: ")
	if _, err := fmt.Scan(&str); err != nil {
		fmt.Printf("Ошибка ввода: %v", err)
	}

	fmt.Println(IsUniqueChars(str))
}
