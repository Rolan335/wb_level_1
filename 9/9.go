package main

import "fmt"

func main() {
	//канал для чисел из массива
	ch1 := make(chan int)
	//канал для x * 2
	ch2 := make(chan int)

	arr := []int{1, 2, 3, 4, 5}

	//Запускаем горутины для записи и умножения. Третья функция не завершится пока все числа из канала не будут прочитаны,
	//что позволит завершиться обеим горутинам.
	go writeX(ch1, arr)
	go writeX2(ch1, ch2)
	writeStdout(ch2)
}

func writeX(out chan<- int, arr []int) {
	//проходимся по массиву и отправляем данные в канал. Как все данные будут отправлены закрываем канал
	for i := range arr {
		out <- arr[i]
	}
	close(out)
}

func writeX2(in <-chan int, out chan<- int) {
	// берём данные из канала 1 и отправляем в канал 2. После, так же его закрываем.
	for x := range in {
		out <- x * 2
	}
	close(out)
}

func writeStdout(in <-chan int) {
	// берём данные из канала 2 и выводим. Канал сам закроется в функции writeX2.
	for x2 := range in {
		fmt.Println(x2)
	}
}
