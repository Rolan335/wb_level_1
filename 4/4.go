package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, wg *sync.WaitGroup, input <-chan int) {

	//Воркер читает данные пока они есть в канале
	for m := range input {
		fmt.Printf("Воркер %d сообщение: %#v\n", id, m)
	}

	//Если данных нет - выходим из цикла и сигнализируем что завершили работу
	fmt.Printf("Воркер %d закончил\n", id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	//Вводим кол-во воркеров с cmd
	fmt.Println("Введите кол-во воркеров")
	var numWorkers int
	fmt.Scanln(&numWorkers)

	//Канал для ввода данных
	input := make(chan int)

	//Добавляем в wg счётчик равный кол-ву воркеров
	wg.Add(numWorkers)
	//Запускаем горутины
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, input)
	}

	//Канал для ожидания ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		select {
			//Если данные из канала получены, закрываем канал ввода и завершаем программу
		case <-c:
			close(input)
			wg.Wait()
			os.Exit(0)
		default:
			//Иначе отправляем данные в канал
			input <- rand.Intn(100)
			time.Sleep(time.Millisecond * 500)
		}
	}
}
