/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел,
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

// Функция для вычисления и вывода квадрата числа.
func CalculateSquare(wg *sync.WaitGroup, num int) {
	defer wg.Done()            // по окончании функции уведомляем wg о завершении горутины
	fmt.Printf("%d ", num*num) //выводим результат в stdout
}

func main() {
	nums := [5]int{2, 4, 6, 8, 10} // Массив чисел из тз
	var wg sync.WaitGroup          // создаем WaitGroup для ожидания завершения всех горутин
	for _, val := range nums {
		wg.Add(1) // Увеличиваем счетчик wg
		go CalculateSquare(&wg, val)
	}

	wg.Wait() // main Ожидает завершения всех горутин
}
