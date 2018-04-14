package main

import (
	"fmt"
)

//Implementação de Struct de Pilha-> utiliza slices para facilitar
type Stack struct{
	 data []Tree
}

func (s Stack) push(info Tree) Stack{
	s.data = append(s.data, info)
	return s
}

func(s Stack) pop() (Stack,Tree){
	if len(s.data)==0{
		return s, *new(Tree)
	}
	var info= s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return s, info
}

func (s Stack) size() int{
	return len(s.data)
}

/*func (s Stack) print() {
	for _, tree := range s.data {
		printBplc(&tree)
	}
}*/

func (s Stack) print() {
	l := len(s.data)
	if l == 0 {
		fmt.Print("0")
		return
	}
	printBplc(&s.data[l-1])
	for i := l - 2; i >= 0; i-- {
		fmt.Print(" ")
		printBplc(&s.data[i])
	}
}
