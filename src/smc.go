package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type SMC struct {
	E map[string]string
	S GenericStack
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

func resolveDesigualdade(typeOf string, smc SMC, forest ...*Tree) *Tree {
	value1, err1 := strconv.Atoi(forest[0].toString())
	value0, err0 := strconv.Atoi(forest[1].toString())
	if err0 != nil {
		value0, _ = strconv.Atoi(findValue(forest[1], smc))
	}
	if err1 != nil {
		value1, _ = strconv.Atoi(findValue(forest[0], smc))
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
	return &Tree{Value: BtoA(boolValue), Sons: nil}
}

var dismember map[string]func(SMC, []*Tree) SMC
var evaluate map[string]func(SMC) SMC

func memFindNext(memory map[string]string) int {
	max := 0
	for k, _ := range memory {
		kv, _ := strconv.Atoi(k)
		if kv >= max {
			max = kv
		}
	}
	return max + 1
}

func createInMemory(ident *Tree, val *Tree, smc SMC) SMC {
	l := memFindNext(smc.M)
	smc.E[ident.toString()] = strconv.Itoa(l)
	smc.M[strconv.Itoa(l)] = val.toString()
	return smc
}

func changeValueInMemory(ident *Tree, val *Tree, smc SMC) (SMC, bool) {
	l, exist := smc.E[ident.toString()]
	if !exist {
		return smc, false
	}
	smc.M[l] = val.toString()
	return smc, true
}

func findValue(ident *Tree, smc SMC) string {
	l, _ := smc.E[ident.toString()]
	val := smc.M[l]
	return val
}

func getTreeFromValueStack(smc SMC) (SMC, *Tree) {
	var genericInfo interface{}
	var typeOfGenericInfo string
	smc.S, genericInfo, typeOfGenericInfo = smc.S.pop()

	if typeOfGenericInfo != "*main.Tree" {
		panic("Erro inesperado")
	}
	var tree = new(Tree)
	tree = (genericInfo).(*Tree)
	return smc, tree
}

func getEnviromentFromValueStack(smc SMC) (SMC, map[string]string) {
	var genericInfo interface{}
	var typeOfGenericInfo string
	smc.S, genericInfo, typeOfGenericInfo = smc.S.pop()

	if typeOfGenericInfo != "map[string]string" {
		println(typeOfGenericInfo)
		panic("Erro inesperado")
	}
	var enviroment = genericInfo.(map[string]string)
	return smc, enviroment
}

func criaMapa() map[string]func(SMC) SMC {
	var evaluate = map[string]func(SMC) SMC{
		"var": func(smc SMC) SMC{
			ident := new(Tree)
			smc, ident = getTreeFromValueStack(smc)
			copyOfEnviroment := make(map[string]string)
			for key, value := range smc.E {
				copyOfEnviroment[key] = value
			}
			smc.E[ident.toString()] = strconv.Itoa(memFindNext(smc.M))
			smc.S = smc.S.push(copyOfEnviroment)
			return smc
		},
		"add": func(smc SMC) SMC {
			var num = 2
			var t = new(Tree)
			var sum = big.NewInt(0)
			for i := 0; i < num; i++ {
				smc, t = getTreeFromValueStack(smc)
				//value, err := strconv.Atoi(t.toString())
				value := big.NewInt(0)
				_, err := value.SetString(t.Value, 10)
				if !err {
					value = big.NewInt(0)
					//value, _ = strconv.Atoi(smc.M[t.toString()])
					value.SetString(findValue(t, smc), 10)
				}
				sum.Add(sum, value)
			}
			smc.S = smc.S.push(&Tree{Value: sum.String(), Sons: nil})
			return smc
		},
		"sub": func(smc SMC) SMC {
			var t1 *(Tree)
			var t0 *(Tree)
			var value0 = big.NewInt(0)
			var value1 = big.NewInt(0)

			smc, t1 = getTreeFromValueStack(smc)
			smc, t0 = getTreeFromValueStack(smc)
			value0, err0 := value0.SetString(t0.Value, 10)
			if !err0 {
				value0 = big.NewInt(0)
				//n := findValue(t0, smc)
				value0.SetString(findValue(t0, smc), 10)
			}
			value1, err1 := value1.SetString(t1.Value, 10)
			if !err1 {
				value1 = big.NewInt(0)
				//value1, _ = strconv.Atoi(smc.M[t1.toString()])
				value1, err1 = value1.SetString(findValue(t1, smc), 10)
			}
			res := value0.Sub(value0, value1)
			smc.S = smc.S.push(&Tree{Value: res.String(), Sons: nil})
			return smc
		},
		"mul": func(smc SMC) SMC {
			var num = 2
			var t = new(Tree)
			var product = big.NewInt(1)
			for i := 0; i < num; i++ {
				smc, t = getTreeFromValueStack(smc)
				value, err := big.NewInt(0).SetString(t.Value, 10)
				if !err {
					value = big.NewInt(0)
					//value, _ = strconv.Atoi(smc.M[t.toString()])
					value, err = value.SetString(findValue(t, smc), 10)
				}
				product.Mul(product, value)
			}
			smc.S = smc.S.push(&Tree{Value: product.String(), Sons: nil})
			return smc
		},
		"div": func(smc SMC) SMC {
			var t1 *(Tree)
			var t0 *(Tree)
			smc, t1 = getTreeFromValueStack(smc)
			smc, t0 = getTreeFromValueStack(smc)
			value0, err0 := big.NewInt(0).SetString(t0.Value, 10)
			if !err0 {
				value0 = big.NewInt(0)
				//value0, _ = strconv.Atoi(smc.M[t0.toString()])
				value0, err0 = value0.SetString(findValue(t0, smc), 10)
			}
			value1, err1 := big.NewInt(0).SetString(t1.Value, 10)
			if !err1 {
				value1 = big.NewInt(0)
				//value1, _ = strconv.Atoi(smc.M[t1.toString()])
				value1, err1 = value1.SetString(findValue(t1, smc), 10)
			}
			smc.S = smc.S.push(&Tree{Value: value0.Div(value0, value1).String(), Sons: nil})
			return smc
		},
		"and": func(smc SMC) SMC {
			var num = 2
			var t = new(Tree)
			var result = true
			for i := 0; i < num; i++ {
				smc, t = getTreeFromValueStack(smc)
				var str = t.toString()
				value, found := smc.M[str]
				if found {
					str = value
				}
				boolValue := (str == "true")
				result = result && boolValue
			}
			smc.S = smc.S.push(&Tree{Value: BtoA(result), Sons: nil})
			return smc
		},
		"or": func(smc SMC) SMC {
			var num = 2
			var t = new(Tree)
			var result = false
			for i := 0; i < num; i++ {
				smc, t = getTreeFromValueStack(smc)
				var str = t.toString()
				value, found := smc.M[str]
				if found {
					str = value
				}
				boolValue := (str == "true")
				result = result || boolValue
			}
			smc.S = smc.S.push(&Tree{Value: BtoA(result), Sons: nil})
			return smc
		},
		"neg": func(smc SMC) SMC {
			value := new(Tree)
			smc, value = getTreeFromValueStack(smc)
			str := value.toString()
			boolVal, found := smc.M[str]
			if found {
				str = boolVal
			}
			boolValue := !(str == "true")
			smc.S = smc.S.push(&Tree{Value: BtoA(boolValue), Sons: nil})
			return smc
		},
		"eq": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc, t1 = getTreeFromValueStack(smc)
			smc, t2 = getTreeFromValueStack(smc)
			smc.S = smc.S.push(resolveDesigualdade("eq", smc, t1, t2))
			return smc
		},
		"gt": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc, t1 = getTreeFromValueStack(smc)
			smc, t2 = getTreeFromValueStack(smc)
			smc.S = smc.S.push(resolveDesigualdade("gt", smc, t1, t2))
			return smc
		},
		"ge": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc, t1 = getTreeFromValueStack(smc)
			smc, t2 = getTreeFromValueStack(smc)
			smc.S = smc.S.push(resolveDesigualdade("gq", smc, t1, t2))
			return smc
		},
		"lt": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc, t1 = getTreeFromValueStack(smc)
			smc, t2 = getTreeFromValueStack(smc)
			smc.S = smc.S.push(resolveDesigualdade("lt", smc, t1, t2))
			return smc
		},
		"le": func(smc SMC) SMC {
			t1 := new(Tree)
			t2 := new(Tree)
			smc, t1 = getTreeFromValueStack(smc)
			smc, t2 = getTreeFromValueStack(smc)
			smc.S = smc.S.push(resolveDesigualdade("le", smc, t1, t2))
			return smc
		},
		"while": func(smc SMC) SMC {
			var result = new(Tree)
			var holdInterface interface{}
			smc.S, holdInterface, _ = smc.S.pop()
			result = holdInterface.(*Tree)
			var exp = new(Tree)
			var bloco = new(Tree)
			smc, exp = getTreeFromValueStack(smc)
			smc, bloco = getTreeFromValueStack(smc)
			if result.toString() == "true" {
				smc.C = smc.C.push(Tree{Value: "while", Sons: append(append(initSons(), exp), bloco)})
				smc.C = smc.C.push(*bloco)
			}
			return smc
		},
		"if": func(smc SMC) SMC {
			var result = new(Tree)
			smc, result = getTreeFromValueStack(smc)
			var blocoIf = new(Tree)
			var blocoElse = new(Tree)
			smc, blocoIf = getTreeFromValueStack(smc)
			smc, blocoElse = getTreeFromValueStack(smc)
			if result.toString() == "true" {
				smc.C = smc.C.push(*blocoIf)
			} else {
				smc.C = smc.C.push(*blocoElse)
			}
			return smc
		},
		"ass": func(smc SMC) SMC {
			ident := new(Tree)
			value := new(Tree)
			smc, value = getTreeFromValueStack(smc)
			smc, ident = getTreeFromValueStack(smc)
			var found bool
			smc, found = changeValueInMemory(ident, value, smc)
			if !found {
				panic(fmt.Sprint("Variable %s not declared.", value.Value))
			}
			return smc
		},
		"clauses": func(smc SMC) SMC {
			return smc
		},
		"init-seq": func(smc SMC) SMC {
			return smc
		},
		"init": func(smc SMC) SMC {
			value := new(Tree)
			ident := new(Tree)
			smc, value = getTreeFromValueStack(smc)
			smc, ident = getTreeFromValueStack(smc)
			smc = createInMemory(ident, value, smc)
			return smc
		},
		"seq": func(smc SMC) SMC {
			return smc
		},
		"noop": func(smc SMC) SMC {
			return smc
		},
		"dec": func(smc SMC) SMC {
			ident := new(Tree)
			smc, ident = getTreeFromValueStack(smc)
			return createInMemory(ident, nil, smc)
		},
		"block": func(smc SMC) SMC {
			var enviroment map[string]string
			smc, enviroment = getEnviromentFromValueStack(smc)
			smc.E = enviroment
			return smc
		},
		"print": func(smc SMC) SMC {
			smc, value := getTreeFromValueStack(smc)
			num, err := big.NewInt(0).SetString(value.Value, 10)
			if err {
				fmt.Println(num)
			} else {
				num = big.NewInt(0)
				num, err = big.NewInt(0).SetString(findValue(value, smc), 10)
				if err {
					fmt.Println(num)
				} else {
					panic(fmt.Sprintf("Variable %s not declared", value.Value))
				}
			}
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
	var dado = new(Tree)
	smc.C, dado = smc.C.pop()
	smc.S = smc.S.push(dado)
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
			smc.S = smc.S.push(forest[1])
			smc.S = smc.S.push(forest[0])
			return smc
		},
		"if": func(smc SMC, forest []*Tree) SMC {
			smc.C = smc.C.push(Tree{Value: "if", Sons: nil})
			smc.C = smc.C.push(*forest[0])
			smc.S = smc.S.push(forest[2])
			smc.S = smc.S.push(forest[1])
			return smc
		},
		// "block": func(smc SMC, forest []*Tree) SMC {
		// 	smc.C = smc.C.push(Tree{Value: "block", Sons: nil})
		// 	for i := (len(forest) - 1); i >= 0; i-- {
		// 		smc.C = smc.C.push(*forest[i])
		// 	}
		// 	copyOfEnviroment := make(map[string]string)
		// 	for key, value := range smc.E {
		// 		copyOfEnviroment[key] = value
		// 	}
		// 	smc.S = smc.S.push(copyOfEnviroment)
		// 	return smc
		// },
	}
	return dismember
}

func (tree Tree) dismember() (string, []*Tree) {
	return tree.Value, tree.Sons
}

func printMap(m *map[string]string) {
	for k, v := range *m {
		fmt.Printf(" %s:%s", k, v)
	}
}

func (smc *SMC) printSmc() {
	fmt.Print("<")
	printMap(&smc.E)
	fmt.Print(", ")
	smc.S.print()
	fmt.Print(",")
	printMap(&smc.M)
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
