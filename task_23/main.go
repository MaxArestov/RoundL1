// Удалить i-ый элемент из слайса.

package main

import "fmt"

// DeleteElement удаляет элемент по индексу из среза любого типа
func DeleteElement[T any](slice []T, index int) []T {
	if index < 0 || index > len(slice)-1 {
		fmt.Printf("В слайсе нет элемента с индексом %d.", index)
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	// Срез байтов
	byteSlice := []byte{1, 2, 3, 4, 5}
	newByteSlice := DeleteElement(byteSlice, 2)
	fmt.Println("Срез байтов:", newByteSlice)

	// Срез строк
	stringSlice := []string{"a", "b", "c", "d", "e"}
	newStringSlice := DeleteElement(stringSlice, 2)
	fmt.Println("Срез строк:", newStringSlice)

	// Срез рун
	runeSlice := []rune{'a', 'b', 'c', 'd', 'e'}
	newRuneSlice := DeleteElement(runeSlice, 2)
	fmt.Println("Срез рун:", newRuneSlice)

	// Срез булов
	boolSlice := []bool{true, false, true, false, true}
	newBoolSlice := DeleteElement(boolSlice, 2)
	fmt.Println("Срез булов:", newBoolSlice)
}
