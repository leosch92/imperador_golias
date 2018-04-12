package main
import(
	"fmt"
	"strconv"
)

type SMC struct{
	S Stack
	M map[string]string
	C Stack
}

var evaluate = map[string] func(string, string) Tree{
	"add": func (op1 string, op2 string) Tree{
		e1,_ := strconv.Atoi(op1)
		e2,_ := strconv.Atoi(op2)	
		return Tree{Left:nil,Value:strconv.Itoa(e1+e2),Right:nil}
	},
}

func iniciaSMC() SMC{
	var smc = *new(SMC)
	smc.M = make(map[string]string)
	return smc
}

//Função En equivale à (S,M,nC)=>(nS,M,C), isto é o valor é transferido da pilha C para a pilha S
func (smc SMC) En() SMC{
	var dado = *(new(Tree))
	smc.C,dado = smc.C.pop()
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

//Função Ee equivale à (S,M,e op e' C)=>(S,M,e e' op C)
func (smc SMC) Ee()SMC{
	if smc.C.size()>2{
		var e = *(new(Tree))
		var eLinha = *(new(Tree))
		var operador = *(new(Tree))
		smc.C,e = smc.C.pop()
		smc.C,operador = smc.C.pop()
		smc.C,eLinha = smc.C.pop()
		smc.C = smc.C.push(operador)
		smc.C = smc.C.push(eLinha)
		smc.C = smc.C.push(e)
	}
	return smc
}

func (smc SMC) Ei()SMC{
	var operando1 = *(new(Tree))
	var operando2 = *(new(Tree))
	var operacao = *(new(Tree))
	smc.S,operando2 = smc.S.pop()
	smc.S,operando1 = smc.S.pop()
	smc.C,operacao = smc.C.pop()
	if(operando1.checkIfNode()&&operando2.checkIfNode()&&operacao.checkIfNode()){
		smc.S = smc.S.push(
			evaluate[operacao.toString()](operando1.toString(),operando2.toString()))
	}
	return smc
}

func (smc SMC) push_tree(tree Tree) (SMC) {
	l_node, value, r_node := tree.dismember()
	smc.C = smc.C.push(Tree{Left:nil, Value:value, Right:nil})
	smc.C = smc.C.push(r_node)
	smc.C = smc.C.push(l_node)
	return smc
}	

func (tree Tree) dismember() (Tree, string, Tree) {
	return *tree.Left, tree.Value, *tree.Right
}

func resolverSMC(smc SMC, t Tree)SMC{
	
	smc.C = smc.C.push(t)
	fmt.Println(smc)
		
	for smc.C.size()>0{
		
		_,op := smc.C.pop()
		if op.checkIfNode(){
			_, isOperation := evaluate[op.toString()]
			if isOperation {
				//Se entrou aqui é porque topo da pilha é uma operação válida
				smc = smc.Ei()
			}
			if !isOperation{
				//Não consegui fazer else não sei porque, mas se entrou aqui é operando no topo da pilha
				smc = smc.En()
			}
		}

		if !op.checkIfNode(){
			//Se entrou aqui topo da pilha é uma árvore com mais de um nó
			smc.C,_ = smc.C.pop()
			smc = smc.push_tree(op)
		}
		fmt.Println(smc)
	}	

	return smc
}
