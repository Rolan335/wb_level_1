package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isUnique("aajqwifpowipork"))
	fmt.Println(isUnique("abcdefghjkl"))
}

func isUnique(s string) bool {
	//Ставим строку в нижний регистр
	str := strings.ToLower(s)
	//Создаём мапу для хранения символов
	sMap := make(map[rune]bool)
	//Проходимся посимвольно
	for _, char := range str {
		//Если символ уже встречался, он будет в мапе со значением true, значит возвращаем false
		if sMap[char]{
			return false
		}
		//каждый символ добавляем в мапу со значением true
		sMap[char] = true
	}
	// если цикл успешно завершился, возвращаем true
	return true
}
