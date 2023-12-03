// Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func mySleep(t time.Duration) {
	<-time.After(t) // After создает канал и отправляет в него сигнал после d единиц времени.
}

func main() {
	fmt.Println("Начало работы функции")
	mySleep(5 * time.Second)
	fmt.Println("Завершение работы функции")
}
