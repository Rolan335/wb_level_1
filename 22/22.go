package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	//Используем библиотеку big для чисел больше int64
	a := big.NewFloat(math.Pow(2, 22))
	b := big.NewFloat(math.Pow(2,24))
	fmt.Println("a + b = ", big.NewFloat(0).Add(a,b))
	fmt.Println("a * b = ", big.NewFloat(0).Mul(a,b))
	fmt.Println("a / b = ", big.NewFloat(0).Quo(a,b))
	fmt.Println("a - b = ", big.NewFloat(0).Sub(a,b))
}