/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

package main

import (
	"fmt"
	"math"
)

// Point структура, представляющая точку в 2D пространстве.
type Point struct {
	x float64
	y float64
}

// NewPoint Создание конструктора Point.
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Distance вычисляет расстояние между двумя точками.
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	p1 := NewPoint(4, 8)
	p2 := NewPoint(2, 75)

	distance := p1.Distance(p2)

	fmt.Printf("Расстояние между точками равно %.2f", distance)
}
