package main

import "fmt"

func main() {
	arr := []int{1, 4, 7, 10, 15, 16, 20, 42}
	search := 15
	fmt.Println(binarySearch(arr, search))
}

func binarySearch(arr []int, search int) int {
	//центр массива
	mid := len(arr) / 2
	var result int
	switch {
	//Если дошли до нулевой длины массива, значит элемента нет.
	case len(arr) == 0:
		result = -1
		//Если искомый элемент меньше центра, рекурсивно запускаем функцию, ищем по первой половине.
	case arr[mid] > search:
		result = binarySearch(arr[:mid], search)
		//Если больше центра, запускаем функцию по второй половине
	case arr[mid] < search:
		result = binarySearch(arr[mid:], search)
		//Если что-то находим (не -1) результат - центральный элемент
		if result >= 0 {
			result += mid
		}
		//Если прошлись полностью - элемент найдем
	default:
		result = mid
	}
	return result
}
