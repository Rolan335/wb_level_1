package main

import "fmt"

func main() {
	//Создаём 2 неупорядоченных массива и вызываем функцию
	a := []int{1, 4, 2, 7, 3, 6, 8, 7}
	b := []int{3, 5, 6, 2, 1, 5, 6, 8}
	fmt.Println(intersection(a,b))
}

func intersection(a, b []int) []int {
	//newSet - возвращаемое пересечение. 
	newSet := make([]int, 0)
	//hashMap - будет хранить уникальные элементы первого множества с значением true
	hashMap := make(map[int]bool)

	//Проходимся по первому множеству, записываем в hashMap данные уникальных значений 
	for _, num := range a {
		hashMap[num] = true
	}
	//Проходимся по второму множеству. Если в hashMap есть значение второго множества,
	//выполнится if. Тем самым находим пересечение множеств и добавляем данные в новое множество 
	for _, num := range b{
		if hashMap[num]{
			newSet = append(newSet, num)
		}
	}

	return newSet
}