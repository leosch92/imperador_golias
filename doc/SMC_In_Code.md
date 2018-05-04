
# SMC_In_Code

This document shows a little bit of how the SMC functions have been implemented in Golang

## Operations

Pop element from control stack and push it into value stack
`<S,M,nC> -> <nS,M,C>`

``` golang
func (smc SMC) En() SMC {
	var dado = (new(Tree))
	smc.C, dado = smc.C.pop()
	smc.S = smc.S.push(*dado)
	return smc
}
```

Reorder elements in control stack 
`<S,M,e op e' C> -> <S, M, e e' op C>`
``` golang
func (smc SMC) push_tree(tree *Tree) SMC {
	value, sons := tree.dismember()
	smc.C = smc.C.push(Tree{Value: value, Sons: nil})
	for i := (len(sons) - 1); i >= 0; i-- {
		smc.C = smc.C.push(*sons[i])
	}
	return smc
}
``` 

Evalute expressions
`<e' e S, M , op C> -> <eval(e,e',op) S, M, C>`
```golang
func (smc SMC) Ei() SMC {
	var operacao = (new(Tree))
	smc.C, operacao = smc.C.pop()
	smc = (evaluate[operacao.toString()](smc))
	return smc
}
```


Reorder elements in While case
`<S,M, while b do c C> -> <b c S, M, b while C>`
```golang
var dismember = map[string]func(SMC, []*Tree) SMC{
		"while": func(smc SMC, forest []*Tree) SMC {
			smc.C = smc.C.push(Tree{Value: "while", Sons: nil})
			smc.C = smc.C.push(*forest[0])
			smc.S = smc.S.push(*forest[1])
			smc.S = smc.S.push(*forest[0])
			return smc
		}
```

Check if entered in while loop
`< true b c S, M , while C> -> < S,M, c while b do c C>`
```golang
func(smc SMC) SMC {
			var result = new(Tree)
			smc.S, result = smc.S.pop()
			var exp = new(Tree)
			var bloco = new(Tree)
			smc.S, exp = smc.S.pop()
			smc.S, bloco = smc.S.pop()
			if result.toString() == "true" {
				smc.C = smc.C.push(Tree{Value: "while", Sons: append(append(initSons(), exp), bloco)})
				smc.C = smc.C.push(*bloco)
			}
			return smc
}
```
