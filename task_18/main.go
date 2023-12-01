/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика с встроенным мьютексом для синхронизации
type Counter struct {
	val int
	mu  sync.Mutex
}

// NewCounter конструктор нового Counter
func NewCounter() *Counter {
	return &Counter{val: 0}
}

// Функция увеличения val в Counter
func (c *Counter) Increment() {
	c.mu.Lock()         // Блокировка доступа к Increment
	defer c.mu.Unlock() //  Гарантия разблокировки после завершения работы функции
	c.val++             // Увеличение val Counter
}

func (c *Counter) GetValue() int {
	c.mu.Lock()         // Блокировка доступа к Increment
	defer c.mu.Unlock() //  Гарантия разблокировки после завершения работы функции
	return c.val        // Возвращение val Counter
}

func main() {
	wg := new(sync.WaitGroup) // Создание указателя на WaitGroup

	counter := NewCounter() // Получение указателя на Counter

	// Запуск горутин для инкрементирования счетчика
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			counter.Increment()
			fmt.Printf("Goroutine %d complete work\n", i) // Доказательство конукрентного выполнения функции.
		}(i)
	}

	wg.Wait() // Ожидание завершения работы всех горутин

	fmt.Printf("value of counetr is %d after increments.", counter.val)
}
