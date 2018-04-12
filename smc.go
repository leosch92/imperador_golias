package main

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
