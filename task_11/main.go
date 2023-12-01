// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
	"math/rand"
)

// Intersection возвращает пересечение двух множеств.
func Intersection(firstSlice, secondSlice []int) []int {
	intersectionMap := make(map[int]bool)
	var intersections []int

	// Заполняем карту элементами первого множества.
	for _, val := range firstSlice {
		intersectionMap[val] = true
	}

	// Проверяем элементы второго множества на наличие в карте.
	for _, item := range secondSlice {
		if _, found := intersectionMap[item]; found {
			intersections = append(intersections, item)
			// Чтобы избежать повторений в результирующем срезе.
			delete(intersectionMap, item)
		}
	}

	return intersections
}

func FillSliceRandomInts(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(101)
	}
}

func main() {
	// Создаем два среза одинаковой(необязательно) длины.
	firstSlice := make([]int, 10)
	secondSlice := make([]int, 10)

	// Заполняем рандомными интами оба слайса.
	FillSliceRandomInts(firstSlice)
	FillSliceRandomInts(secondSlice)

	// Выводим слайсы.
	fmt.Println(firstSlice)
	fmt.Println(secondSlice)

	// Создаем слайс и заполняем его пересечениями.
	intersectionSlice := Intersection(firstSlice, secondSlice)

	// Выводим пересечения двух слайсов.
	fmt.Printf("Intersections:\n%v", intersectionSlice)
}
