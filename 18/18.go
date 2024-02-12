package main

import (
	"fmt"
	"sync"
)

//Структура счётчика с встроенным mutex
type Counter struct {
	count int
	lock  sync.Mutex
}

//Метод инкрементирования. При вызове, блокирует метод для других горутин.
func (c *Counter) Incr(){
	c.lock.Lock()
	c.count++
	c.lock.Unlock()
}

func main() {
	//Используем wg чтобы все горутины завершили выполнение
	var wg sync.WaitGroup

	counter := Counter{}

	//запускаем горутины и ждём выполнения. Из горутины вызываем метод счётчика
	for i := 0; i < 10; i++{
		wg.Add(1)
		go func(){
			counter.Incr()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}