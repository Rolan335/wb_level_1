package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(createSet(arr))
}

func createSet(arr []string) map[string]bool {
	//создаём мапу для множества
	set := make(map[string]bool)
	// Итерируемся через массив и присваиваем значения в set. 
	// Такая конструкция сделает невозможным дублирование элементов что нам у нужно.
	for _, str := range arr {
		set[str] = true
	}

	return set
}
