package main
import "fmt"

type SMC struct{
	S stack
	//Memória funciona como um mapa -> dado uma posição é encontrada uma info (não sei se é melhor solução)
	M map[string]string
	C stack
}

//Função En equivale à (S,M,nC)=>(nS,M,C), isto é o valor é transferido da pilha C para a pilha S
func (smc SMC) En() SMC{
	//Garante que existe dado na Pilha C
	if smc.C.size()>0{
		var dado = ""
		smc.C, dado = smc.C.pop()
		smc.S = smc.S.push(dado)
	}
	return smc
}

//Função Ev equivale à (S,M,vC)=>(M(v)S,M,C), isto é o valor na posição v em M é colocada na pilha S
func (smc SMC) Ev() SMC{
	if smc.C.size()>0{
		var dado = ""
		smc.C, dado = smc.C.pop()
		smc.S = smc.S.push(smc.M[dado])
	}
	return smc
}

func (smc SMC) Ee()SMC{
	if smc.C.size()>2{
		var e = ""
		var eLinha = ""
		var operador = ""
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
	if smc.C.size()>2{
		var operando1 = ""
		var operando2 = ""
		var operacao = ""
		smc.C,operando1 = smc.C.pop()
		smc.C,operando2 = smc.C.pop()
		smc.C,operacao = smc.C.pop()
		smc.S = smc.S.push("eval("+operando1+operacao+operando2+")")
	}
	return smc
}


func main(){
	var smc = *new(SMC)
	smc.M = make(map[string]string)
	fmt.Println("Inicializa SMC")
	fmt.Println(smc)
	fmt.Println("\nImportante: ler pilha da direita para esquerda")

	smc.C = smc.C.push("e'")
	smc.C = smc.C.push("+")
	smc.C = smc.C.push("e")
	smc.C = smc.C.push("dado1")
	smc.C = smc.C.push("dado2")
	smc.S = smc.S.push("dado3")
	smc.M["dado1"] = "dado4"
	fmt.Println("\nSMC preenchida")
	fmt.Println(smc)

	fmt.Println("\nFunção En equivale à (S,M,nC)=>(nS,M,C), isso é o valor é transferido da pilha C para a pilha S")
	smc = smc.En()
	fmt.Println(smc)

	fmt.Println("\nFunção Ev equivale à (S,M,vC)=>(M(v)S,M,C), isso é o valor da posição v de M é colocada na pilha S")
	smc = smc.Ev()
	fmt.Println(smc)

	fmt.Println("\nFunção Ee equivale à (S,M,e op e' C)=>(S,M,e e' op C), isso é a operação é coloca no modo pós-fixado")
	smc = smc.Ee()
	fmt.Println(smc)

	fmt.Println("\nFunção Ei equivale à (S,M,e e' op C)=>(eval(e,op,e')S,M,C), isso é a operação é executada e colacada na pilha S")
	smc = smc.Ei()
	fmt.Println(smc)

}