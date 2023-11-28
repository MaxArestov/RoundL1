/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
)

// CalculateSquares	Создаем функцию, отправляющую в канал результат
func CalculateSquares(wg *sync.WaitGroup, num int, ch chan<- int) {
	defer wg.Done()
	ch <- num * num
	//TODO: delete
	fmt.Println(num * num)
}

func main() {
	nums := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	var sum int

	chanInt := make(chan int, len(nums))

	for _, val := range nums {
		wg.Add(1)
		go CalculateSquares(&wg, val, chanInt)
	}
	wg.Wait()
	close(chanInt)

	for square := range chanInt {
		sum += square
	}

	fmt.Println(sum)
}
