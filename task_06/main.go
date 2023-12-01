// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ByContext Функция будет завершаться через контекст.
func ByContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда функция отработает.

	for {
		select {
		case <-ctx.Done(): // Ожидаем уведомления о завершении работы от переданного контекста.
			fmt.Println("Gogoutine closed by context")
			return
		default:
			fmt.Println("Goroutine working and w8 for ctx.Done!")
			time.Sleep(2 * time.Second)
		}
	}
}

// ByChan Функция будет завершаться через канал.
func ByChan(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда функция отработает.

	for {
		select {
		case <-ch: // Ожидаем чего-либо (В данном случае close(ch)) из канала.
			fmt.Println("Goroutine closed by channel")
			return
		default:
			fmt.Println("Goroutine working and w8 for smth from ch")
			time.Sleep(2 * time.Second)
		}
	}
}

// ByTimer ожидает завершения таймера.
func ByTimer(wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик, когда функция отработает.

	timer := time.NewTimer(5 * time.Second) //Создаем таймер.
	for {
		select {
		case <-timer.C: //Ожидаем подачи сигнала из канала таймера. Снова канал.
			fmt.Println("Goroutine closed by timer")
			return
		default:
			fmt.Println("Goroutine working and w8 for timer")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	wg := new(sync.WaitGroup) // Создается WaitGroup.

	ctx, cancel := context.WithCancel(context.Background()) // Создается контекст.

	ch := make(chan int) //Создается канал.

	wg.Add(3) // Увеличивается счетчик WaitGroup.

	// Запускается все 3 функции в отдельных горутинах.
	go ByContext(ctx, wg)
	go ByChan(ch, wg)
	go ByTimer(wg)

	// ch <- 5 Если раскомментить, то горутина с каналом завершится без close(ch)

	time.Sleep(5 * time.Second) // 5 секунд на работу горутин.

	cancel()  // Отменяется контекст.
	close(ch) // Закрывается канал

	wg.Wait() // Ожидание завершения работы всех горутин.
	fmt.Println("All goroutines finish work")
}
