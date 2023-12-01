/*
Имеется последовательность строк - (cat, cat, dog, cat, tree)
Создать для нее собственное множество.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Исходный список строк
	oldStrings := []string{"cat", "cat", "dog", "cat", "tree", "dog", "fish", "Fish", "fIsh"}

	newStrings := make([]string, 0, len(oldStrings))

	newUniqueStrings := make([]string, 0, len(oldStrings))
	// Создаем множество неуникальных значений
	mapStrings := make(map[string]bool)
	for _, str := range oldStrings {
		mapStrings[str] = true
	}

	// Создаем множество уникальных значений
	mapUniqueStrings := make(map[string]bool)
	for _, str := range oldStrings {
		uniqueString := strings.ToLower(str)
		mapUniqueStrings[uniqueString] = true
	}

	// Заполняем неуникальное множество новых strings
	for key := range mapStrings {
		newStrings = append(newStrings, key)
	}

	// Заполняем уникальное множество новых strings
	for key := range mapUniqueStrings {
		newUniqueStrings = append(newUniqueStrings, key)
	}

	// Выводим элементы множеств
	fmt.Println(newStrings)
	fmt.Println(newUniqueStrings)
}
