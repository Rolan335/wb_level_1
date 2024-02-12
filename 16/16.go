package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 5, 1, 6, 8, 32, 14, 62, 22, 67, 10}

	sortedArr := quickSort(arr)
	fmt.Println(sortedArr)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// выбираем начальный опорный элемент
	pivot := len(arr) / 2

	//Начальная и конечная точка массива
	left, right := 0, len(arr)-1

	//меняем местами опорный элемент с конечным элементом
	arr[pivot], arr[right] = arr[right], arr[pivot]

	//Проходимся по массиву. Если элемент меньше опорного, ставим его в начало.
	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			//сдвигаем начальную точку массива на +1
			left++
		}
	}
	// Меняем местами крайний и начальный элемент
	arr[left], arr[right] = arr[right], arr[left]

	//Рекурсивно вызываем функцию для первой и второй половины массива
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}
