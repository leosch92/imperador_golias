package main

import (
	"reflect"
)

var toBPLC = map[string]string{
	"+":   "add",
	"-":   "sub",
	"*":   "mul",
	"/":   "div",
	"~":   "neg",
	">=":  "ge",
	">":   "gt",
	"<=":  "le",
	"<":   "lt",
	"==":  "eq",
	"/\\": "and",
	"\\/": "or",
}

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

func evalArithExpr(first, rest interface{}) *Tree {
	stringType := reflect.TypeOf("string teste")
	var t Tree

	if reflect.TypeOf(first) == stringType {
		t = Tree{first.(string), initSons()}
	} else {
		t = *first.(*Tree)
	}

	restSl := toIfaceSlice(rest)

	for _, v := range restSl {
		restExpr := toIfaceSlice(v)
		operator := toBPLC[restExpr[1].(string)]
		auxTree := t
		t = Tree{operator, initSons()}

		t.Sons = append(t.Sons, &auxTree)

		// Se recebe string da interface, constroi árvore nova, senão pega árvore já construída
		// do retorno do eval de um Factor
		rightOp := restExpr[3]

		if reflect.TypeOf(rightOp) == stringType {

			// Insere o operando da direita como novo filho
			t.Sons = append(t.Sons, &Tree{rightOp.(string), initSons()})

		} else {
			t.Sons = append(t.Sons, rightOp.(*Tree))
		}

	}
	return &t
}

func evalWhile(boolExp, body interface{}) *Tree {
	t := Tree{"while", initSons()}
	t.Sons = append(t.Sons, boolExp.(*Tree))
	t.Sons = append(t.Sons, body.(*Tree))
	return &t
}

func evalKeywordBoolExp(keyword interface{}) *Tree {
	return &Tree{keyword.(string), initSons()}
}

func evalUnaryBoolExp(unBoolOp, boolExp interface{}) *Tree {
	t := Tree{toBPLC[unBoolOp.(string)], initSons()}
	t.Sons = append(t.Sons, boolExp.(*Tree))
	return &t
}

func evalBinaryLogicExp(binBoolOp, id1, id2 interface{}) *Tree {
	t := Tree{toBPLC[binBoolOp.(string)], initSons()}
	t.Sons = append(t.Sons, &Tree{id1.(string), initSons()})
	t.Sons = append(t.Sons, &Tree{id2.(string), initSons()})
	return &t
}

func evalBinaryArithExp(binBoolOp, id1, id2 interface{}) *Tree {
	t := Tree{toBPLC[binBoolOp.(string)], initSons()}
	t.Sons = append(t.Sons, &Tree{id1.(string), initSons()})
	t.Sons = append(t.Sons, id2.(*Tree))
	return &t
}

func evalAssignment(id, expr interface{}) *Tree {
	t := Tree{"ass", initSons()}
	t.Sons = append(t.Sons, &Tree{id.(string), initSons()})
	t.Sons = append(t.Sons, expr.(*Tree))
	return &t
}

func evalPrint(exp interface{}) *Tree {
	t := Tree{"print", initSons()}
	t.Sons = append(t.Sons, exp.(*Tree))
	return &t
}

func evalSequence(first, rest interface{}) *Tree {
	t := Tree{"seq", initSons()}
	t.Sons = append(t.Sons, first.(*Tree))

	restSl := toIfaceSlice(rest)

	for _, v := range restSl {
		restCmd := toIfaceSlice(v)
		// o elemento 0 contém espaçamento e o 1 contém de fato uma árvore de comando
		t.Sons = append(t.Sons, restCmd[1].(*Tree))
	}

	return &t
}

func evalBlock(decl, cmd interface{}) *Tree {
	if decl != nil {

	}
	tCmd := cmd.(*Tree)
	return tCmd
}

func evalIf(boolExp, ifBody, elseStatement interface{}) *Tree {
	t := Tree{"if", initSons()}
	t.Sons = append(t.Sons, boolExp.(*Tree))
	t.Sons = append(t.Sons, ifBody.(*Tree))

	if elseStatement == nil {
		t.Sons = append(t.Sons, &Tree{"noop", initSons()})
	} else {
		// Pega o slice correspondente à string else(0), espaçamento(1) e bloco(2)
		elseSlice := toIfaceSlice(elseStatement)
		t.Sons = append(t.Sons, elseSlice[2].(*Tree))
	}

	return &t
}

func evalInitialization(sInit, rest interface{}) *Tree {
	tFirst := sInit.(*Tree)

	if rest != nil {
		t := Tree{"init-seq", initSons()}
		t.Sons = append(t.Sons, tFirst)
		restSlice := toIfaceSlice(rest)
		for _, value := range restSlice {
			restSingleInit := toIfaceSlice(value)
			t.Sons = append(t.Sons, restSingleInit[3].(*Tree))
		}
		return &t
	}

	return tFirst
}

func evalDeclOp(typ interface{}) *Tree {
	if typ.(string) == "var" {
		return &Tree{"is_var", initSons()}
	} else {
		return &Tree{"is_const", initSons()}
	}
}

func evalSingleInit(id, expr interface{}) *Tree {
	t := Tree{"init", initSons()}
	t.Sons = append(t.Sons, &Tree{id.(string), initSons()})
	t.Sons = append(t.Sons, expr.(*Tree))
	return &t
}

func evalClauses(variable, constant, init, cmd interface{}) *Tree {
	var t *Tree
	if variable != nil {
		t = toIfaceSlice(variable)[0].(*Tree)
	}
	if constant != nil {

	}

	aux := findLastRightSon(t)
	if init != nil {
		if t != nil {
			aux.Sons = append(aux.Sons, toIfaceSlice(init)[0].(*Tree))
		} else {
			t = toIfaceSlice(init)[0].(*Tree)
		}
	}
	aux.Sons = append(aux.Sons, cmd.(*Tree))
	return t
}

func findLastRightSon(tp *Tree) *Tree {
	aux := tp
	for len(aux.Sons) > 1 {
		aux = aux.Sons[len(aux.Sons)-1]
	}
	return aux
}

func evalVariable(id, rest interface{}) *Tree {
	tFirstBlock := Tree{"block", initSons()}
	tVar := Tree{"var", initSons()}
	tVar.Sons = append(tVar.Sons, &Tree{id.(string), initSons()})
	tFirstBlock.Sons = append(tFirstBlock.Sons, &tVar)

	tPreviousBlock := &tFirstBlock
	if rest != nil {
		restSlice := toIfaceSlice(rest)

		for _, value := range restSlice {
			variableSlice := toIfaceSlice(value)
			tBlock := Tree{"block", initSons()}
			tVar := Tree{"var", initSons()}
			tVar.Sons = append(tVar.Sons, &Tree{variableSlice[3].(string), initSons()})
			tBlock.Sons = append(tBlock.Sons, &tVar)
			tPreviousBlock.Sons = append(tPreviousBlock.Sons, &tBlock)
			tPreviousBlock = &tBlock
		}

		return &tFirstBlock
	}

	return &tFirstBlock
}
