package main

import(
	"fmt"
	"reflect"
)

func main(){
	//Создаём переменную интерфейс
	var a interface{}
	
	//Реализуем и вызываем функцию
	a = 12

	printType(a)

	a = "hey"

	printType(a)

	a = true

	printType(a)

	a = make(chan int)

	printType(a)

}

func printType(a interface{}){
	//reflect.TypeOf() динамически выводит тип переменной 
	fmt.Println(reflect.TypeOf(a))
}