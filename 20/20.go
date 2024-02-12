package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "dog snow run snow snow"
	fmt.Println(reverseOrder(str))
}

func reverseOrder(str string) string {
	//С помощью пакета string преобразуем строку в массив
	strArr := strings.Split(str, " ")
	//i - считает с начала. j - с конца. С каждым шагом меняем первое и последнее слово
	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	//преобразуем массив обратно в строку
	return strings.Join(strArr, " ")
}
