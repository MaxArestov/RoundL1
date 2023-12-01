//Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

// ConcurrentMap - структура для работы с картой в многопоточноcти
type ConcurrentMap struct {
	mu    sync.Mutex
	items map[string]int
}

// NewConcurrentMap создает новый экземпляр ConcurrentMap
func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		items: make(map[string]int),
	}
}

// Set устанавливает значение для ключа с блокировкой доступа
func (m *ConcurrentMap) Set(key string, val int) {
	m.mu.Lock() // Блокировка доступа к мапе
	defer m.mu.Unlock()
	m.items[key] = val
}

// Get возвращает значение по ключу с блокировкой доступа
func (m *ConcurrentMap) Get(key string) (int, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.items[key]
	return val, ok
}

func main() {
	cm := NewConcurrentMap() // Создание мапы

	var wg sync.WaitGroup // Создание WaitGroup для гарантии ожидания
	// Осуществление конкурентной записи данных
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			cm.Set(key, n)
		}(i)
	}
	wg.Wait()

	// Вывод из мапы
	for i := 0; i < 30; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := cm.Get(key); ok {
			fmt.Printf("{key: %s} [val: %d]\n", key, val)
		} else {
			fmt.Printf("\n no value %d with key %s in %v", key, val, cm)
		}
	}
}
