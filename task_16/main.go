// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import (
	"fmt"
)

// quickSort реализует алгоритм быстрой сортировки
func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Выбор опорного элемента
	pivot := len(a) / 2

	// Перемещение опорного элемента в конец
	a[pivot], a[right] = a[right], a[pivot]

	// Перемещение элементов меньших, чем опорный, влево
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	// Перемещение опорного элемента на его конечное место
	a[left], a[right] = a[right], a[left]

	// Рекурсивная сортировка левой и правой части
	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

func main() {
	array := []int{3, 1, 4, 1, 5, 64841, 164, 145, -46}
	fmt.Println("Исходный массив:", array)

	sortedArray := quickSort(array)
	fmt.Println("Отсортированный массив:", sortedArray)
}
