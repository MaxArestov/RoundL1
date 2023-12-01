/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// PushToChannel создание функции по отправке int'ов в канал 1.
func PushToChannel(wg *sync.WaitGroup, nums []int, firstChan chan<- int) {
	defer wg.Done() // Уменьшаем количество wg.

	for i := 0; i < len(nums); i++ {
		fmt.Printf("Pusher: Pushing %d to channel 1.\n", nums[i])
		firstChan <- nums[i] // Отправка int из слайса в канал 1.
	}
	close(firstChan) // Закрытие канала.
}

// ReadAndPushToChannel чтение из канала 1, возведение в квадрат и отправка в канал 2.
func ReadAndPushToChannel(wg *sync.WaitGroup, firstChan <-chan int, secondChan chan<- int) {
	defer wg.Done() // Уменьшаем количество wg.

	for val := range firstChan { // Проходим циклом по всем значениям, полученным из канала 1.
		fmt.Printf("Reader&Pusher: Reading %d, Pushing %d to channel 2.\n", val, val*val)
		secondChan <- val * val // Отправка значения из канала 1, возведенного во 2 степень, в канал 2.
	}
	close(secondChan) // Закрытие канала.
}

// ChanReader чтение и вывод в stdout из канала 2.
func ChanReader(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done() // Уменьшаем количество wg.

	for val := range ch { // Проходим циклом по данным из 2 канала и выводим в stdout.
		fmt.Printf("Reader: received %d from channel 2.\n", val)
	}
}

func main() {
	// Создание WaitGroup(указателя) и двух каналов.
	wg := new(sync.WaitGroup)
	firstChan := make(chan int)
	secondChan := make(chan int)

	nums := make([]int, 4) // Создание слайса для отправки в 1 канал.

	// Заполнение рандомными интами 64.
	for i := 0; i < len(nums); i++ {
		nums[i] = int(rand.Int63n(501))
	}
	fmt.Println(nums)

	//Добавление 3 wg.
	wg.Add(3)

	// Запуск трех функций в отдельных горутинах.
	go PushToChannel(wg, nums, firstChan)
	go ReadAndPushToChannel(wg, firstChan, secondChan)
	go ChanReader(wg, secondChan)

	// Ожидание окончания работы всех горутин и объявление о завершении работы программы.
	wg.Wait()
	fmt.Println("All goroutines finished work.")
}
