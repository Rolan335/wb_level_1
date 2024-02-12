package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Мапа для хранения результата
	groups := make(map[int][]float64)

	for _, temp := range temperatures{
		//Округляем чтобы преобразовать значение к группе с шагом 10
		group := int(math.Round(temp/10)) * 10
		//Добавляем температуру к группе
		groups[group] = append(groups[group], temp)
	}
	fmt.Println(groups)
}
