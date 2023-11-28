/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var duration int // Продолжительность выполнения программы в секундах
	fmt.Print("Введите целое число, означающее количество секунд работы программы ")
	_, err := fmt.Scan(&duration)

	// Если пользователь ввел некорректные данные - уведомление и повторная попытка ввода.
	if err != nil {
		for err != nil {
			fmt.Println()
			fmt.Printf("Некорректно введено число. Введите целое число, означающее время работы программы: ")
			_, err = fmt.Scan(&duration)
		}
	}

	// Установка дефолтного значения секунд работы при настырном пользователе
	if duration == 0 || duration < 0 {
		fmt.Println("Incorrect data. Using default count of seconds: 10")
		duration = 10
	}
	fmt.Printf("Программа проработает %d секунд.\n", duration)

	ch := make(chan int) // Создание канала для отправки данных.

	done := make(chan bool) // Создание канала для остановки работы горутин.

	var wg sync.WaitGroup // Создание WaitGroup(не ссылкой, так как wg будет передаваться корректно).

	wg.Add(2)
	// Горутина для отправки данных в канал
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond) // Интервал отправки
		}
	}()

	// Горутина для чтения данных из канала
	go func() {
		defer wg.Done()
		for {
			select {
			case v := <-ch:
				fmt.Println("Получено из канала: ", v)
			case <-done:
				return
			}
		}
	}()

	// Ожидание и завершение работы после N секунд
	time.Sleep(time.Duration(duration) * time.Second)
	close(done)
	fmt.Println("Время истекло, программа завершила работу")
}
