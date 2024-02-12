package main

import "fmt"

func main() {
	a := 10
	b := 78

	//что-то вроде xor обмена средствами языка
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

	c, d := 33, 55

	//Алгоритмически
	fmt.Println(c, d)
	c = c + d
	d = c - d
	c = c - d
	fmt.Println(c, d)

}
