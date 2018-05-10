package main

import(
	"reflect"
)

type GenericStack struct {
	data [] interface{}
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

func (s GenericStack) print(){
	
}