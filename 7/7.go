package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)

	var mutex sync.Mutex
	//Запускаем 10 горутин
	for i := 1; i <= 10; i++ {
		go func(i int) {
			//Блокируем остальные горутины для записи и записываем данные в map
			mutex.Lock()
			m[i] = i * 2
			//Разблокировка
			mutex.Unlock()
		}(i)
	}
	fmt.Println(m)
}
