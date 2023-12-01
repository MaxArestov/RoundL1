/*
Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

// FillMapWithGrouping заполняет map согласно группировке.
func FillMapWithGrouping(nums []float64, groupTemperatures map[int][]float64) map[int][]float64 {
	var key int
	for _, val := range nums {
		if val > 0 {
			key = int(math.Floor(val/10) * 10)
			groupTemperatures[key] = append(groupTemperatures[key], val)
		} else { // Ключ получаем методом Ceil, чтобы -2х.х не попадали в категорию -30:{...
			key = int(math.Ceil(val/10) * 10)
			groupTemperatures[key] = append(groupTemperatures[key], val)
		}
	}
	return groupTemperatures
}

// PrintMap распечатывает мапу согласно ТЗ.
func PrintMap(printingMap map[int][]float64) {
	var i int
	for key, value := range printingMap {
		fmt.Printf("%d:{", key)

		for i, num := range value {
			fmt.Printf("%.1f", num)
			if i < len(value)-1 {
				fmt.Printf(", ")
			}
		}

		if i < len(printingMap)-1 {
			fmt.Printf("}, ")
		} else {
			fmt.Printf("}.")
		}
		i++
	}
}

func main() {
	groupingTemps := make(map[int][]float64) // Создаем пустую карту, куда сложим температуры по группам.

	// Создаем слайс температур и заполняем его рандомными значениями как >0, так и <0.
	var temperatures []float64
	for i := 0; i < 111; i++ {
		k := rand.Int63n(2)
		if k == 0 {
			temperatures = append(temperatures, float64(rand.Int63n(46))+rand.Float64())
		} else {
			temperatures = append(temperatures, (float64(rand.Int63n(46))+rand.Float64())*-1)
		}
	}

	FillMapWithGrouping(temperatures, groupingTemps) // Заполняем мапу.

	PrintMap(groupingTemps) // Печатаем мапу.
}
