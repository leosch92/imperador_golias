package main

import (
	"fmt"
	"strconv"
)

type SMC struct {
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

var opCtrl = map[string]bool{
	"while": true,
	"if":    true,
	"att":   true,
	"seq":   true,
}

var pilhaExp Stack
var pilhaBloc Stack

var evaluate map[string]func(SMC) SMC

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
					value, _ = strconv.Atoi(smc.M[t.toString()])
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
				value0, _ = strconv.Atoi(smc.M[t0.toString()])
			}
			value1, err1 := strconv.Atoi(t1.toString())
			if err1 != nil {
				value1, _ = strconv.Atoi(smc.M[t1.toString()])
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
					value, _ = strconv.Atoi(smc.M[t.toString()])
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
				value0, _ = strconv.Atoi(smc.M[t0.toString()])
			}
			value1, err1 := strconv.Atoi(t1.toString())
			if err1 != nil {
				value1, _ = strconv.Atoi(smc.M[t1.toString()])
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
				boolValue := (t.toString() == "true")
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
				boolValue := (t.toString() == "true")
				result = result || boolValue
			}
			smc.S = smc.S.push(Tree{Value: BtoA(result), Sons: nil})
			return smc
		},
		"neg": func(smc SMC) SMC {
			value := new(Tree)
			smc.S, value = smc.S.pop()
			result := !(value.toString() == "true")
			smc.S = smc.S.push(Tree{Value: BtoA(result), Sons: nil})
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
		"lq": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc.S, t1 = smc.S.pop()
			smc.S, t2 = smc.S.pop()
			smc.S = smc.S.push(resolveDesigualdade("lq", smc, t1, t2))
			return smc
		},
		"while": func(smc SMC) SMC {
			var exp = new(Tree)
			var bloco = new(Tree)
			smc.C, exp = smc.C.pop()
			smc.C, bloco = smc.C.pop()
			smcResp := iniciaSMC()
			smcResp.M = smc.M
			smcResp = resolverSMC(smcResp, *exp)
			_, resp := smcResp.S.pop()
			if resp.toString() == "true" {
				pilhaExp = pilhaExp.push(*exp)
				pilhaBloc = pilhaBloc.push(*bloco)
				smc.C = smc.C.push(Tree{Value: "endwhile", Sons: nil})
				smc.C = smc.C.push(*bloco)
			}
			return smc
		},
		"endwhile": func(smc SMC) SMC {
			var exp *Tree
			var bloco *Tree
			pilhaExp, exp = pilhaExp.pop()
			pilhaBloc, bloco = pilhaBloc.pop()
			smc.C = smc.C.push(*bloco)
			smc.C = smc.C.push(*exp)
			smc.C = smc.C.push(Tree{Value: "while", Sons: nil})
			return smc
		},
		"if": func(smc SMC) SMC {
			exp := new(Tree)
			bloco := new(Tree)
			smc.C, exp = smc.C.pop()
			smc.C, bloco = smc.C.pop()
			blkOrEnd := new(Tree)
			smc.C, blkOrEnd = smc.C.pop()
			smcResp := iniciaSMC()
			smcResp.M = smc.M
			smcResp = resolverSMC(smcResp, *exp)
			_, treeResp := smcResp.S.pop()

			if !blkOrEnd.checkIfNode() || blkOrEnd.toString() != "endif" {
				smc.C, _ = smc.C.pop()
			}

			if treeResp.toString() == "true" {
				smc.C = smc.C.push(*bloco)
			} else {
				if !blkOrEnd.checkIfNode() || blkOrEnd.toString() != "endif" {
					smc.C = smc.C.push(*blkOrEnd)
				}
			}

			return smc
		},
		"att": func(smc SMC) SMC {
			ident := new(Tree)
			exp := new(Tree)
			smc.C, ident = smc.C.pop()
			smc.C, exp = smc.C.pop()
			smcResp := iniciaSMC()
			smcResp.M = smc.M
			_, value := (resolverSMC(smcResp, *exp)).S.pop()
			smc.M[ident.toString()] = value.toString()
			return smc
		},
		"seq": func(smc SMC) SMC {
			return smc
		},
	}
	return evaluate
}

func iniciaSMC() SMC {
	var smc = *new(SMC)
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
	_, isCtrl := opCtrl[value]
	if !isCtrl {
		smc.C = smc.C.push(Tree{Value: value, Sons: nil})
	} else if value == "if" {
		smc.C = smc.C.push(Tree{Value: "endif", Sons: nil})
	}
	for i := (len(sons) - 1); i >= 0; i-- {
		smc.C = smc.C.push(*sons[i])
	}
	if isCtrl {
		smc.C = smc.C.push(Tree{Value: value, Sons: nil})
	}
	return smc
}

func (tree Tree) dismember() (string, []*Tree) {
	return tree.Value, tree.Sons
}

/*func (smc SMC) printSmc() {
	fmt.Print("<")
	smc.S.print()
	fmt.Print(", ")
	fmt.Print("0")
	fmt.Print(", ")
	smc.C.print()
	fmt.Print(">")
	fmt.Println()
}*/

func resolverSMC(smc SMC, t Tree) SMC {
	evaluate = criaMapa()
	pilhaExp = (*new(Stack))
	pilhaBloc = (*new(Stack))
	smc.C = smc.C.push(t)
	fmt.Println(smc)
	//smc.printSmc()
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
		fmt.Println(smc)
		//smc.printSmc()
	}

	return smc

}
