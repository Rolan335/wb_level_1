package main

import "fmt"

//Два интерфейса для html и xml
type XMLReader interface {
	ReadXML() string
}

type HTMLReader interface {
	ReadHTML() string
}

//Адаптер для XMLReader
type XMLReaderAdapter struct {
	HTMLReader HTMLReader
}

//Метод адаптера
func (x *XMLReaderAdapter) ReadXML() string {
	html := x.HTMLReader.ReadHTML()
	xml := html + " converting html to xml... "
	return xml
}

//Реализация нашего htmlreader
type NewHTMLReader struct {
}

func (h *NewHTMLReader) ReadHTML() string {
	return "html from ReadHTML()"
}

func main() {
	//Создаём экземпляр адаптера и вызываем метод адаптера ReadXML()
	adapter := &XMLReaderAdapter{HTMLReader: &NewHTMLReader{}}
	xml := adapter.ReadXML()
	fmt.Println(xml)
}
