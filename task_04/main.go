/*
*Реализовать постоянную запись данных в канал (главный поток).
*Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
*Необходима возможность выбора количества воркеров при старте.
*Программа должна завершаться по нажатию Ctrl+C.
*Выбрать и обосновать способ завершения работы всех воркеров.
 */
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Функция worker представляет собой воркера, который выполняет чтение из канала.
func worker(ctx context.Context, id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Гарантируем отправку сигнала о завершении при завершении функции.
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				fmt.Printf("Worker {%d} exiting\n", id) // Выводим, если канал закрыт и завершаем функцию.
				return
			}
			// Выводим полученные из канала данные
			fmt.Printf("Worker{%d} receive %d\n", id, num)
		case <-ctx.Done():
			// Завершаем работу воркера, если контекст был отменен.
			fmt.Printf("worker {%d} complete reading\n", id)
			return
		}
	}
}

func main() {
	var numOfWorkers int

	// Чтение количества воркеров от пользователя.
	fmt.Print("Введите целое число, означающее количество воркеров: ")
	_, err := fmt.Scan(&numOfWorkers)

	// Если пользователь ввел некорректные данные - уведомление и повторная попытка ввода.
	if err != nil {
		for err != nil {
			fmt.Println()
			fmt.Printf("Некорректно введено число. Введите целое число, означающее количество воркеров: ")
			_, err = fmt.Scan(&numOfWorkers)
		}
	}

	// Установка дефолтного значения воркеров при настырном пользователе
	if numOfWorkers == 0 || numOfWorkers < 0 {
		fmt.Println("Incorrect data. Using default count of workers: 1")
		numOfWorkers = 1
	}
	fmt.Println("Количество воркеров", numOfWorkers)

	ch := make(chan int) // Создание канала для передачи данных воркерам

	ctx, cancel := context.WithCancel(context.Background()) // контекст для контроля горутин

	wg := &sync.WaitGroup{} // WaitGroup для ожидания завершения всех воркеров (Как указатель для удобства).

	// Запуск воркеров
	for i := 0; i < numOfWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, wg)
	}

	// Обработка сигнала Ctrl+C
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		cancel()  // Отменяем контекст, что завершит работу воркеров.
		close(ch) // Закрываем канал передачи воркерам.
	}()

	// Отдельная горутина для отправки данных в канал.
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond) // Таймер для отправки данных.
		defer ticker.Stop()                              // Гарантирует завершение тикера при завершении функции.

		for {
			select {
			case <-ctx.Done():
				// При завершении контекста завершаем отправку данных.
				return
			case <-ticker.C:
				ch <- time.Now().Nanosecond() // Отправляем данные в канал.
			}
		}
	}()

	//Ожидаем завершение работы всех воркеров.
	wg.Wait()
}
