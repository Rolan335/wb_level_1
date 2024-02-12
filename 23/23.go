package main

import "fmt"

func main() {
	slice := []int{2, 4, 23, 5, 235, 325, 15}
	i := 4
	newSlice := deleteElem(slice, i)

	fmt.Println(slice)
	fmt.Println(newSlice)
}

func deleteElem(s []int, i int) []int {
	//убираем элемент с индексом i путём соединения (append) значений до него и после него.
	return append(s[:i], s[i+1:]...)
}