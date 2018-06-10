package main

type Location string
type Constante string

type EnviromentValue interface{
	converteParaString() string
}

func (loc Location) converteParaString() string{
	return string(loc)
} 

func (constante Constante) converteParaString() string{
	return string(constante)
}