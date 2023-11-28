/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

// Human Родительская структура
type Human struct {
	name   string
	height int
}

// Функции родительской структуры
func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) SetHeight(height int) {
	h.height = height
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetHeight() int {
	return h.height
}

// Action Структура-наследник
type Action struct {
	Human // Встраивание родительской структуры - аналог наследования
	speed int
}

// SetSpeed Самостоятельная функция наследника
func (a *Action) SetSpeed(newSpeed int) {
	a.speed = newSpeed
}

func main() {
	// Создание экземпляра Action
	action := Action{Human: Human{"John", 185}, speed: 7}

	// Вывод изначальных данных
	fmt.Printf("Initial name is %s and height is %d, speed is %d.\n", action.name, action.height, action.speed)

	// Использование Action родительских и собственной функции
	action.SetName("Wick... John Wick!")
	action.SetHeight(186)
	action.SetSpeed(9)

	// Вывод новых данных
	fmt.Printf("New name is %s New height is %d and speed now is over %dk",
		action.name, action.height, action.speed)
}
