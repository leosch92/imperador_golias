package main

type Location string
type Constante string

type EnviromentValue interface{
	converteParaString() string
	WhatIsMyType() string
}

func (loc Location) converteParaString() string{
	return string(loc)
} 

func (constante Constante) converteParaString() string{
	return string(constante)
}

func (tree *Tree) converteParaString() string{
	return tree.toString()
}

func (loc Location) WhatIsMyType() string{
	return "main.Location"
} 

func (constante Constante) WhatIsMyType() string{
	return "main.Constante"
}

func (tree *Tree) WhatIsMyType() string{
	return "main.*Tree"
}