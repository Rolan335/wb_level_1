package main

import (
	"fmt"
	"sync"
)

func main(){
	three()
}

func three() {
	arr := [5]int{2, 4, 6, 8, 10}
	sum := 0

	//Используем waitGroup чтобы все горутины завершили выполнение
	var waitGroup sync.WaitGroup

	//Итерируемся по массиву
	for _, i := range arr {
		//увеличиваем счётчик горутин в waitgroup и запускаем горутину
		waitGroup.Add(1)
		go func(i int) {
			// суммируем квадрат числа и сигнализируем что горутина закончила выполнение
			defer waitGroup.Done()
			sum = (i * i) + sum
		}(i)

	}
	//ждём пока всё выполнится
	waitGroup.Wait()
	//Выводим квадрат числа после завершения всех горутин
	fmt.Println(sum)
}