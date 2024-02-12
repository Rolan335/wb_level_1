package main

import (
	"fmt"
)

func setBit(n int64, i uint, bitValue int) int64{
	//Маска где i бит = 1
	mask := int64(1 << i)
	fmt.Printf("До преобразования: %b \n", n)
	fmt.Printf("Маска: %b \n", mask)
	if bitValue == 1 {
		//1 - дизъюнкция ставит 1
		n |= mask
	} else {
		// 0 - сброс бита &^ 0
		n &^= mask
	}
	return n
}

func main() {
	num1 := setBit(10, 2, 1)
	num2 := setBit(1000, 3, 0)
	fmt.Printf("После преобразования num1: %b \n", num1)
	fmt.Printf("После преобразования num2: %b \n", num2)
}
