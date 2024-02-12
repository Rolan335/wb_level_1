package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// Функция случайного числа чтобы передать значение в канал
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	//Считываем с cmd число секунд, устанавливаем таймер
	n, _ := strconv.Atoi(os.Args[1])
	t := time.NewTimer(time.Second * time.Duration(n))
	//Создаём канал и waitgroup
	ch := make(chan int)
	var wg sync.WaitGroup

	//Указатель чтобы завершить цикл for
loop:
	for {
		select {
		//Канал t.C Объявляет что таймер истёк, завершаем цикл
		case <-t.C:
			break loop
		default:
			//Иначе запускаем горутины последовательно
			wg.Add(2)
			go out(ch, &wg)
			go in(ch, &wg)
			wg.Wait()
		}
	}
}

// Функция, отправляющая данные в канал
func out(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand := random(0, 100)
	out <- rand
	fmt.Println("sended ", rand)
}

// Функция, получающая и выводящая данные из канала
func in(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("got ", <-in)
}
