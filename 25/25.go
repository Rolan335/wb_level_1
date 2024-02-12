package main

import (
	"fmt"
	"time"
)

func sleep(sec int) {
	//При запуске функции, select ждёт пока выполнится условие
	select {
		//Как только пройдёт количество указанных секунд, time.After отправит данные в канал что приведёт к return и функция закончит выполнение.
	case <-time.After(time.Second * time.Duration(sec)):
		return
	}
}

func main() {
	fmt.Println("Запускаем")
	sleep(2)
	fmt.Println("Прошло две секунды")
}
