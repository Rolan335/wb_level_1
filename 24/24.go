package main

import (
	"fmt"
	"math"
)

//Инкапсуляция в golang есть только на уровне пакета. В одном пакете параметры инкапсулировать невозможно

//Структура point
type Point struct {
	x float64
	y float64
}

//"Конструктор" привязанный к структуре point возвращает новую структуру
func (p Point) Create(x float64, y float64) Point {
	return Point{x, y}
}

//Метод структуры принимает значение второй структуры и вычисляет расстояние между 2мя point
func (p Point) findDistanceFromPoint(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2) + math.Pow(p2.y-p.y, 2))
}

func main() {
	point := Point.Create(Point{}, 10, 5)
	point2 := Point.Create(Point{}, 15, 10)
	fmt.Println(point.findDistanceFromPoint(point2))
}
