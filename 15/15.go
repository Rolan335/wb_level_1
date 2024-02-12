package main

import "fmt"

var justString string

//Функция createHugeString создаёт строку длинной 1024 байта,
//но в переменной justString хранятся первые 100 байт строки.
//Остальные 924 байта становятся бесполезными и занимают память

func createHugeString(size int) string {

	return "huge string JKLGJLSKDJGLKASJDGHOPDSHAGOPHJEPGI9H4T2938HT9823HT9EJGJOPWQJEGIWEQOPGJWEOIPJGOWIEJTOPQR9-23JREWIOJFOPGINDSOVPNAOPNEOPNSOINVWOPIEJWGPOJGOIPWQJETIOHWEGOPHJSDOPHGOPSNDGOPNWOIVNEOPINVAOPINVOEGOIHWEOPGAIHOWEPIHGOPWIAHEGOPISHDOGIHSADIOGJSDNVOPSINAVOPISDGIHOSIHGJOPWEJTOPIHQTOIHWEOFNSDOPFNDSOAFNOSDNFAIO"
}

func someFunc() {
	//Уберём лишнюю переменную v, вызовем функцию, и сразу возьмём только первые 100 символов.
	//Тем самым, ненужная переменная не будет занимать место в памяти
	justString = createHugeString(1 << 10)[:100]

	fmt.Println(justString)
}

func main() {
	someFunc()
}
