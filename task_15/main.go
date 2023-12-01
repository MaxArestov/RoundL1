/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

Данный фрагмент кода может привести к проблемам с управлением памятью.
JustString получает срез строки v, созданной в someFunc(), но продолжает указывать на HugeString,
мешая Garbage Collector'у.
В Go строки неизменяемы(иммутабельность), и срез строки фактически ссылается на ту же память, что и исходная строка.
Это означает, что, хотя justString содержит только 100 символов, в памяти остается захваченным весь большой блок данных,
созданный функцией createHugeString.

Это может привести к чрезмерному использованию памяти, особенно если createHugeString создает очень большие строки,
и эти строки не освобождаются для сборщика мусора из-за того, что меньший срез все еще на них ссылается.
*/

package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Printf("Строка до копирования: %s\n", v) // Выводим v для демонстрации исходной стринги.
	// Создаем новую строку, чтобы избежать сохранения всей строки v в памяти.
	justString = string(v[:100])
}

/*
	Можно также привести строку к слайсу байтов (sliceBytes :=[]byte(v)[:100]

justBytes := make([]byte, 100)
copy(justBytes, sliceBytes)
justString := string(justBytes)
*/
func createHugeString(size int) string {
	// Создаем большую строку, повторяя символ "a"
	return strings.Repeat("a", size)
}

func main() {
	someFunc()
	fmt.Printf("Строка после копирования 100 элементов: %s\n", justString) // Выводим justString для демонстрации
}
