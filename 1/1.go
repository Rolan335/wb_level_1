package main

//Структура Human с набором полей
type Human struct {
	name   string
	age    int
	gender string
}

//Структура Action с полем Human
type Action struct {
	run  bool
	walk bool
	Human
}

//Метод структуры Human
func (h *Human) isSleep() bool {
	return true
}

func main() {
	//Создаём экземпляр Human и экземпляр Action
	human := Human{"alex", 21, "male"}
	action := Action{true, true, human}

	//Теперь мы можем обратиться к методу структуры human через Action
	action.isSleep()
}
