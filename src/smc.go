package main

import (
	"fmt"
	"strconv"
)

type SMC struct {
	E map[string]string
	S Stack
	M map[string]string
	C Stack
}

//Função que converte booleanos em String
func BtoA(boolValue bool) string {
	resultStr := ""
	if boolValue {
		resultStr = "true"
	} else {
		resultStr = "false"
	}
	return resultStr
}

func resolveDesigualdade(typeOf string, smc SMC, forest ...*Tree) Tree {
	value1, err1 := strconv.Atoi(forest[0].toString())
	value0, err0 := strconv.Atoi(forest[1].toString())
	if err0 != nil {
		value0, _ = strconv.Atoi(smc.M[forest[1].toString()])
	}
	if err1 != nil {
		value1, _ = strconv.Atoi(smc.M[forest[0].toString()])
	}

	var boolValue bool
	switch typeOf {
	case "eq":
		boolValue = (value0 == value1)
		break
	case "ge":
		boolValue = (value0 >= value1)
		break
	case "le":
		boolValue = (value0 <= value1)
		break
	case "gt":
		boolValue = (value0 > value1)
		break
	case "lt":
		boolValue = (value0 < value1)
		break
	}
	return Tree{Value: BtoA(boolValue), Sons: nil}
}

var dismember map[string]func(SMC, []*Tree) SMC
var evaluate map[string]func(SMC) SMC

func memFindNext(memory map[string]string) int {
	max := 0
	for k, _ := range memory {
		kv, _ := strconv.Atoi(k)
		if max >= kv {
			max = kv
		}
	}
	return max + 1
}

func toMemory(ident *Tree, val *Tree, smc SMC) int {
	/*"att": func(smc SMC) SMC {
		ident := new(Tree)
		value := new(Tree)
		smc.S, value = smc.S.pop()
		smc.S, ident = smc.S.pop()
		smc.M[ident.toString()] = value.toString()
		return smc
	}*/
	l := memFindNext(smc.M)
	smc.E[ident.toString()] = strconv.Itoa(l)
	smc.M[strconv.Itoa(l)] = val.toString()
	return l
}

func findValue(ident *Tree, smc SMC) string {
	l := smc.E[ident.toString()]
	val := smc.M[l]
	return val
}

func criaMapa() map[string]func(SMC) SMC {
	var evaluate = map[string]func(SMC) SMC{
		"add": func(smc SMC) SMC {
			var num = 2
			var t = new(Tree)
			var sum = 0
			for i := 0; i < num; i++ {
				smc.S, t = smc.S.pop()
				value, err := strconv.Atoi(t.toString())
				if err != nil {
					/*value, _ = strconv.Atoi(smc.M[t.toString()])*/
					value, _ = strconv.Atoi(findValue(t, smc))
				}
				sum += value
			}
			smc.S = smc.S.push(Tree{Value: strconv.Itoa(sum), Sons: nil})
			return smc
		},
		"sub": func(smc SMC) SMC {
			var t1 *(Tree)
			var t0 *(Tree)
			smc.S, t1 = smc.S.pop()
			smc.S, t0 = smc.S.pop()
			value0, err0 := strconv.Atoi(t0.toString())
			if err0 != nil {
				/*value0, _ = strconv.Atoi(smc.M[t0.toString()])*/
				value0, _ = strconv.Atoi(findValue(t0, smc))
			}
			value1, err1 := strconv.Atoi(t1.toString())
			if err1 != nil {
				/*value1, _ = strconv.Atoi(smc.M[t1.toString()])*/
				value1, _ = strconv.Atoi(findValue(t1, smc))
			}
			smc.S = smc.S.push(Tree{Value: strconv.Itoa(value0 - value1), Sons: nil})
			return smc
		},
		"mul": func(smc SMC) SMC {
			var num = 2
			var t = new(Tree)
			var product = 1
			for i := 0; i < num; i++ {
				smc.S, t = smc.S.pop()
				value, err := strconv.Atoi(t.toString())
				if err != nil {
					/*value, _ = strconv.Atoi(smc.M[t.toString()])*/
					value, _ = strconv.Atoi(findValue(t, smc))
				}
				product *= value
			}
			smc.S = smc.S.push(Tree{Value: strconv.Itoa(product), Sons: nil})
			return smc
		},
		"div": func(smc SMC) SMC {
			var t1 *(Tree)
			var t0 *(Tree)
			smc.S, t1 = smc.S.pop()
			smc.S, t0 = smc.S.pop()
			value0, err0 := strconv.Atoi(t0.toString())
			if err0 != nil {
				/*value0, _ = strconv.Atoi(smc.M[t0.toString()])*/
				value0, _ = strconv.Atoi(findValue(t0, smc))
			}
			value1, err1 := strconv.Atoi(t1.toString())
			if err1 != nil {
				/*value1, _ = strconv.Atoi(smc.M[t1.toString()])*/
				value1, _ = strconv.Atoi(findValue(t1, smc))
			}
			smc.S = smc.S.push(Tree{Value: strconv.Itoa(value0 / value1), Sons: nil})
			return smc
		},
		"and": func(smc SMC) SMC {
			var num = 2
			var t = new(Tree)
			var result = true
			for i := 0; i < num; i++ {
				smc.S, t = smc.S.pop()
				var str = t.toString()
				value, found := smc.M[str]
				if found {
					str = value
				}
				boolValue := (str == "true")
				result = result && boolValue
			}
			smc.S = smc.S.push(Tree{Value: BtoA(result), Sons: nil})
			return smc
		},
		"or": func(smc SMC) SMC {
			var num = 2
			var t = new(Tree)
			var result = false
			for i := 0; i < num; i++ {
				smc.S, t = smc.S.pop()
				var str = t.toString()
				value, found := smc.M[str]
				if found {
					str = value
				}
				boolValue := (str == "true")
				result = result || boolValue
			}
			smc.S = smc.S.push(Tree{Value: BtoA(result), Sons: nil})
			return smc
		},
		"neg": func(smc SMC) SMC {
			value := new(Tree)
			smc.S, value = smc.S.pop()
			str := value.toString()
			boolVal, found := smc.M[str]
			if found {
				str = boolVal
			}
			boolValue := !(str == "true")
			smc.S = smc.S.push(Tree{Value: BtoA(boolValue), Sons: nil})
			return smc
		},
		"eq": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc.S, t1 = smc.S.pop()
			smc.S, t2 = smc.S.pop()
			smc.S = smc.S.push(resolveDesigualdade("eq", smc, t1, t2))
			return smc
		},
		"gt": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc.S, t1 = smc.S.pop()
			smc.S, t2 = smc.S.pop()
			smc.S = smc.S.push(resolveDesigualdade("gt", smc, t1, t2))
			return smc
		},
		"ge": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc.S, t1 = smc.S.pop()
			smc.S, t2 = smc.S.pop()
			smc.S = smc.S.push(resolveDesigualdade("gq", smc, t1, t2))
			return smc
		},
		"lt": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc.S, t1 = smc.S.pop()
			smc.S, t2 = smc.S.pop()
			smc.S = smc.S.push(resolveDesigualdade("lt", smc, t1, t2))
			return smc
		},
		"le": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc.S, t1 = smc.S.pop()
			smc.S, t2 = smc.S.pop()
			smc.S = smc.S.push(resolveDesigualdade("le", smc, t1, t2))
			return smc
		},
		"while": func(smc SMC) SMC {
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
		},
		"if": func(smc SMC) SMC {
			var result = new(Tree)
			smc.S, result = smc.S.pop()
			var blocoIf = new(Tree)
			var blocoElse = new(Tree)
			smc.S, blocoIf = smc.S.pop()
			smc.S, blocoElse = smc.S.pop()
			if result.toString() == "true" {
				smc.C = smc.C.push(*blocoIf)
			} else {
				smc.C = smc.C.push(*blocoElse)
			}
			return smc
		},
		"att": func(smc SMC) SMC {
			ident := new(Tree)
			value := new(Tree)
			smc.S, value = smc.S.pop()
			smc.S, ident = smc.S.pop()
			/*smc.M[ident.toString()] = value.toString()*/
			toMemory(ident, value, smc)
			return smc
		},
		"seq": func(smc SMC) SMC {
			return smc
		},
		"noop": func(smc SMC) SMC {
			return smc
		},
	}
	return evaluate
}

func iniciaSMC() SMC {
	var smc = *new(SMC)
	smc.E = make(map[string]string)
	smc.M = make(map[string]string)
	return smc
}

//Função En equivale à (S,M,nC)=>(nS,M,C), isto é o valor é transferido da pilha C para a pilha S
func (smc SMC) En() SMC {
	var dado = (new(Tree))
	smc.C, dado = smc.C.pop()
	smc.S = smc.S.push(*dado)
	return smc
}

//Função Ev equivale à (S,M,vC)=>(M(v)S,M,C), isto é o valor na posição v em M é colocada na pilha S
//func (smc SMC) Ev() SMC{
//	var dado = *(new(Tree))
//	smc.C, dado = smc.C.pop()
//	smc.S = smc.S.push(smc.M[dado])
//	return smc
//}

func (smc SMC) Ei() SMC {
	var operacao = (new(Tree))
	smc.C, operacao = smc.C.pop()
	smc = (evaluate[operacao.toString()](smc))
	return smc
}

func printaOperandos(forest []*Tree) {
	for _, v := range forest {
		fmt.Println(v.Value)
	}
}

func (smc SMC) push_tree(tree *Tree) SMC {
	value, sons := tree.dismember()
	funcDismember, isCtrl := dismember[value]
	if isCtrl {
		return funcDismember(smc, sons)
	}
	smc.C = smc.C.push(Tree{Value: value, Sons: nil})
	for i := (len(sons) - 1); i >= 0; i-- {
		smc.C = smc.C.push(*sons[i])
	}
	return smc
}

func criaMapaDismember() map[string]func(SMC, []*Tree) SMC {
	var dismember = map[string]func(SMC, []*Tree) SMC{
		"while": func(smc SMC, forest []*Tree) SMC {
			smc.C = smc.C.push(Tree{Value: "while", Sons: nil})
			smc.C = smc.C.push(*forest[0])
			smc.S = smc.S.push(*forest[1])
			smc.S = smc.S.push(*forest[0])
			return smc
		},
		"if": func(smc SMC, forest []*Tree) SMC {
			smc.C = smc.C.push(Tree{Value: "if", Sons: nil})
			smc.C = smc.C.push(*forest[0])
			smc.S = smc.S.push(*forest[2])
			smc.S = smc.S.push(*forest[1])
			return smc
		},
	}
	return dismember
}

func (tree Tree) dismember() (string, []*Tree) {
	return tree.Value, tree.Sons
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf(" %s:%s", k, v)
	}
}

func (smc SMC) printSmc() {
	fmt.Print("<")
	smc.S.print()
	fmt.Print(",")
	printMap(smc.M)
	fmt.Print(", ")
	smc.C.print()
	fmt.Print(">")
	fmt.Println()
}

func resolverSMC(smc SMC, t Tree, verbose bool) SMC {
	evaluate = criaMapa()
	dismember = criaMapaDismember()
	smc.C = smc.C.push(t)
	if verbose {
		smc.printSmc()
	}
	//fmt.Println(smc)
	for smc.C.size() > 0 {
		_, op := smc.C.pop()
		if op.checkIfNode() {
			_, isOperation := evaluate[op.toString()]
			if isOperation {
				//Se entrou aqui é porque topo da pilha é uma operação válida
				smc = smc.Ei()
			} else {
				//Se entrou aqui é operando no topo da pilha
				smc = smc.En()
			}

		} else {
			//Se entrou aqui topo da pilha é uma árvore com mais de um nó
			smc.C, _ = smc.C.pop()
			smc = smc.push_tree(op)
		}
		//fmt.Println(smc)
		if verbose {
			smc.printSmc()
		}
	}

	return smc

}
