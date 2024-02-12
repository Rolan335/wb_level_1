package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//1. С помощью сигнализирующего канала

	stopChan := make(chan bool)
	//Запускаем горутину, ждём 3 секунды и в стопканал отправляем сигнал о завершении работы
	go stop1(stopChan)
	time.Sleep(time.Second * 3)

	stopChan <- true
	
	fmt.Println("stopped 1")


	//2. С помощью sync.waitgroup

	var w sync.WaitGroup
	//Добавляем счётчик, запускаем горутину и ждём w.Done()
	w.Add(1)
	go stop2(&w)
	w.Wait()
	fmt.Println("stopped 2")

}

func stop2(w *sync.WaitGroup) {
	//Что-то делаем и сигнализируем что завершили выполнение.
	fmt.Println("do smthn 2")
	time.Sleep(time.Second)
	w.Done()
}

func stop1(stop <-chan bool) {
	for {
		select {
		//Благодаря оператору select горутина будет "Работать" пока из канала stop не будут получены данные
		default:
			fmt.Println("working 1")
			time.Sleep(time.Second)
		case <-stop:
			return
		}

	}
}
