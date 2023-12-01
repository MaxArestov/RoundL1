// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

// BinarySearch реализует алгоритм бинарного поиска
// Возвращает индекс искомого элемента или -1, если элемент не найден
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := left + (right-left)/2 // Вычисление среднего индекса

		// Проверка, равен ли средний элемент целевому
		if arr[middle] == target {
			return middle
		}

		// Решение, в какую часть массива двигаться дальше
		if arr[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1 // Цикл завершился, элемент не найден - возвращаем -1
}

func main() {
	array := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21} // Создаем срез интов
	target := 19                                          // Задаем таргет

	result := BinarySearch(array, target)
	if result != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, result)
	} else {
		fmt.Printf("Элемент %d не найден.\n", target)
	}
}
