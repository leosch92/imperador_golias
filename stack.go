package main

//Implementação de Struct de Pilha-> utiliza slices para facilitar
type stack struct{
	 data []string
}

func (s stack) push(info string) stack{
	s.data = append(s.data, info)
	return s
}

func(s stack) pop() (stack,string){
	if len(s.data)==0{
		return s, ""
	}
	var info= s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return s, info
}

func (s stack) size() int{
	return len(s.data)
}