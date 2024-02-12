package main

import "fmt"

func main() {
	fmt.Println(reverse(" ۝hello۩  its ۞ me ک"))
}

func reverse(str string) string {
	//Конвертируем строку в массив рун чтобы работать с символами юникода
	rstr := []rune(str)
	//i - считает с начала. j - с конца. С каждым шагом меняем первый и последний символ
	for i, j := 0, len(rstr)-1; i < j; i, j = i+1, j-1 {
		rstr[i], rstr[j] = rstr[j], rstr[i]
	}
	return string(rstr)
}
