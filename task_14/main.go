/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

package main

import (
	"fmt"
	"reflect"
)

// determineType принимает интерфейс и определяет его тип
func determineType(i interface{}) {
	// Использование type switch для определения типа i
	switch v := i.(type) {
	case int:
		fmt.Printf("Переменная имеет тип int: %d\n", v)
	case string:
		fmt.Printf("Переменная имеет тип string: %s\n", v)
	case bool:
		fmt.Printf("Переменная имеет тип bool: %t\n", v)
	case chan int: // Пример с каналом, обрабатывающим int
		fmt.Printf("Переменная имеет тип channel: %s\n", reflect.TypeOf(v))
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	// Тестирование функции determineType с разными типами
	determineType(5) // Тест с типом int

	determineType("Hello") // Тест с типом string

	determineType(true) // Тест с типом bool

	determineType(make(chan int)) // Тест с типом channel

	determineType(new(struct{}))
}
