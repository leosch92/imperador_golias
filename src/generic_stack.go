package main

import (
	"fmt"
	"reflect"
)

type GenericStack struct {
	data []interface{}
}

func (s GenericStack) push(info interface{}) GenericStack {
	s.data = append(s.data, info)
	return s
}

func (s GenericStack) pop() (GenericStack, interface{}, string) {
	var info = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return s, info, reflect.TypeOf(info).String()
}

func (s GenericStack) size() int {
	return len(s.data)
}

func printElem(e *interface{}) {
	var typeOf = reflect.TypeOf(*e).String()
	if typeOf == "*main.Tree" {
		var tree = new(Tree)
		tree = (*e).(*Tree)
		printBplc(tree)
	} else if typeOf == "*map[string]string" {
		var env = (*e).(*map[string]string)
		printMap(env)
	} else {
		fmt.Print(reflect.TypeOf(*e).String())
	}
}

func (s GenericStack) print() {
	l := s.size()
	if l == 0 {
		fmt.Print("0")
		return
	}
	printElem(&s.data[l-1])
	for i := l - 2; i >= 0; i-- {
		fmt.Print(" ")
		printElem(&s.data[i])
	}
}
