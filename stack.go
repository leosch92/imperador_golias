package main

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