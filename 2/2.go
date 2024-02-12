package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	//Используем waitGroup чтобы все горутины завершили выполнение
	var waitGroup sync.WaitGroup

	//Итерируемся по массиву
	for _, i := range arr {
		//увеличиваем кол-во горутин в waitgroup и запускаем горутину
		waitGroup.Add(1)
		go func(i int) {
			// выводим квадрат числа и сигнализируем что горутина закончила выполнение
			defer waitGroup.Done()
			fmt.Println(i * i)
		}(i)

	}
	//ждём пока всё выполнится
	waitGroup.Wait()
}
