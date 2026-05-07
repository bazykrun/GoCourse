package main

import "fmt"

/*Примечание если объявить любой объект маленькой буквой то его будет видно лишь в области одного пакета но при импорте
его не будет видно, если объявить с большой то она будет публичной, используя маленькие буквы мы используем инкапсуляцию */
type Integer struct {
	integer int
	integer2 int
	integer8 int8
	integer16 int16
	integer32 int32
	integer64 int64
	/* главное что нужно знать они не связаны между собой хотя и представляют один класс
	т.к для каждого из них выделяется разное кол-во памяти и места в результате чего если их 
	смешать то может произойти хаос но между собой они работают отлично */
}

func main() {
	integers := Integer{
		integer: 1,
		integer2: 2,
		integer8: 8,
		integer16: 16,
		integer32: 32,
		integer64: 64,
	}
	fmt.Println(integers.integer + integers.integer2) /* выводится 3 т.к мы объединили один и тот же тип между собой */
	fmt.Println(integers.integer + integers.integer8) /* выйдет ошибка (mismatched types int and int8) 
	она озночает что мы здесь хотим объединить два разных типа данных int и int8 */
}
