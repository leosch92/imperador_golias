package main

import (
	"fmt"
	"os"
	"reflect"
)

var toBPLC = map[string]string{
	"+":     "add",
	"-":     "sub",
	"*":     "mul",
	"/":     "div",
	"~":     "neg",
	">=":    "ge",
	">":     "gt",
	"<=":    "le",
	"<":     "lt",
	"==":    "eq",
	"/\\":   "and",
	"\\/":   "or",
	"var":   "ref",
	"const": "cns",
}

type Declaration struct {
	Typ   string
	Name  string
	Value *Tree
}

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

func constructArithExpr(first, rest interface{}) *Tree {
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
		// do retorno do construct de um Factor
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

func constructWhile(boolExp, body interface{}) *Tree {
	t := Tree{"while", initSons()}
	t.Sons = append(t.Sons, boolExp.(*Tree))
	t.Sons = append(t.Sons, body.(*Tree))
	return &t
}

func constructKeywordBoolExp(keyword interface{}) *Tree {
	return &Tree{keyword.(string), initSons()}
}

func constructUnaryBoolExp(unBoolOp, boolExp interface{}) *Tree {
	t := Tree{toBPLC[unBoolOp.(string)], initSons()}
	t.Sons = append(t.Sons, boolExp.(*Tree))
	return &t
}

func constructBinaryLogicExp(binBoolOp, id1, id2 interface{}) *Tree {
	t := Tree{toBPLC[binBoolOp.(string)], initSons()}
	t.Sons = append(t.Sons, &Tree{id1.(string), initSons()})
	t.Sons = append(t.Sons, &Tree{id2.(string), initSons()})
	return &t
}

func constructBinaryArithExp(binBoolOp, id1, id2 interface{}) *Tree {
	t := Tree{toBPLC[binBoolOp.(string)], initSons()}
	t.Sons = append(t.Sons, &Tree{id1.(string), initSons()})
	t.Sons = append(t.Sons, id2.(*Tree))
	return &t
}

func constructAssignment(id, expr interface{}) *Tree {
	t := Tree{"ass", initSons()}
	t.Sons = append(t.Sons, &Tree{id.(string), initSons()})
	t.Sons = append(t.Sons, expr.(*Tree))
	return &t
}

func constructPrint(exp interface{}) *Tree {
	t := Tree{"print", initSons()}
	t.Sons = append(t.Sons, exp.(*Tree))
	return &t
}

func constructSequence(first, rest interface{}) *Tree {
	t := Tree{"seq", initSons()}
	t.Sons = append(t.Sons, first.(*Tree))

	restSl := toIfaceSlice(rest)

	for _, v := range restSl {
		restCmd := toIfaceSlice(v)
		// o elemento 0 contém espaçamento e o 1 contém de fato uma árvore de comando
		t.Sons = append(t.Sons, restCmd[3].(*Tree))
	}

	return &t
}

func constructBlock(declSeq, cmd interface{}) *Tree {
	var declSeqConv []*Declaration
	if declSeq != nil {
		declSeqConv = toIfaceSlice(declSeq)[0].([]*Declaration)
	}

	cmdConv := cmd.(*Tree)
	var declTrees []*Tree

	for _, decl := range declSeqConv {
		declTrees = append(declTrees, &Tree{toBPLC[decl.Typ], []*Tree{
			&Tree{decl.Name, initSons()}, // nó de identificador
			decl.Value}})                 // nó da expressão
	}

	t, lastNode := buildTreeWithDeclarationBlocks(declTrees)
	if t != nil {
		lastNode.Sons = append(lastNode.Sons, cmdConv)
		return t
	} else {
		return cmdConv
	}
}

func constructIf(boolExp, ifBody, elseStatement interface{}) *Tree {
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

func constructInitialization(sInit, rest interface{}) map[string]*Tree {
	inits := make(map[string]*Tree, 0)
	identifier := toIfaceSlice(sInit)[0].(string)
	expression := toIfaceSlice(sInit)[1].(*Tree)
	inits[identifier] = expression
	restSlice := toIfaceSlice(rest)
	for _, initWithSpacesInterface := range restSlice {
		initWithSpaces := toIfaceSlice(initWithSpacesInterface)
		init := initWithSpaces[3]
		identifier = toIfaceSlice(init)[0].(string)
		expression = toIfaceSlice(init)[1].(*Tree)
		inits[identifier] = expression
	}

	return inits
}

func constructSingleInit(id, expr interface{}) []interface{} {
	singleInit := make([]interface{}, 0)
	singleInit = append(singleInit, id)
	singleInit = append(singleInit, expr)
	return singleInit
}

func findLastBlockSon(tp *Tree) *Tree {
	aux := tp
	for aux.Sons[len(aux.Sons)-1].Value == "blk" {
		aux = aux.Sons[len(aux.Sons)-1]
	}
	return aux
}

func constructClauseDeclaration(id, rest interface{}) []string {
	declared := make([]string, 0)
	declared = append(declared, id.(string))

	restSlice := toIfaceSlice(rest)
	for _, singleRest := range restSlice {
		varWithSpaces := toIfaceSlice(singleRest)
		declared = append(declared, varWithSpaces[3].(string))
	}

	return declared
}

func constructDeclaration(declOp, initSeq interface{}) []*Declaration {
	declOpConv := declOp.(string)
	initSeqConv := initSeq.(map[string]*Tree)
	var declarations []*Declaration

	for k, v := range initSeqConv {
		declarations = append(declarations, &Declaration{declOpConv, k, v})
	}

	/*for _, decl := range declarations {
		fmt.Println(decl.Typ, decl.Name, decl.Value)
	}*/
	return declarations
}

func constructClauses(variable, constant, init, proc interface{}) *Tree {
	variableConverted, constantConverted, initConverted, procConverted := checkAndConvert(variable, constant, init, proc)

	if len(initConverted) != (len(variableConverted) + len(constantConverted)) {
		fmt.Println("Numero de variáveis declaradas e inicializadas é diferente. Encerrando o programa...")
		os.Exit(-1)
		return nil
	} else {
		variableTrees := createDeclarationTree(variableConverted, initConverted, "ref")
		constantTrees := createDeclarationTree(constantConverted, initConverted, "cns")
		allTrees := append(variableTrees, constantTrees...)
		t, lastNode := buildTreeWithDeclarationBlocks(allTrees)
		lastNode.Sons = append(lastNode.Sons, procConverted)

		return t
	}
}

func checkAndConvert(variable, constant, init, cmd interface{}) ([]string, []string, map[string]*Tree, *Tree) {
	var variableConv, constantConv []string
	var initConv map[string]*Tree
	if variable != nil {
		variableConv = toIfaceSlice(variable)[0].([]string)
	} else {
		variableConv = []string{}
	}
	if constant != nil {
		constantConv = toIfaceSlice(constant)[0].([]string)
	} else {
		constantConv = []string{}
	}
	if init != nil {
		initConv = toIfaceSlice(init)[0].(map[string]*Tree)
	} else {
		initConv = map[string]*Tree{}
	}
	return variableConv, constantConv, initConv, cmd.(*Tree)
}

func createDeclarationTree(variable []string, init map[string]*Tree, declOp string) []*Tree {
	var varTrees []*Tree
	for _, v := range variable {
		expression, ok := init[v]
		if !ok {
			conversao := map[string]string{"cns": "Constante", "ref": "Variável"}
			fmt.Printf("%s %s declarada mas não inicializada. Encerrando o programa...\n", conversao[declOp], v)
			os.Exit(-1)
		} else {
			idTree := &Tree{v, initSons()}
			varTrees = append(varTrees, &Tree{declOp, []*Tree{idTree, expression}})
		}
	}

	return varTrees
}

func buildTreeWithDeclarationBlocks(allTrees []*Tree) (*Tree, *Tree) {
	var t, last, new *Tree
	for i, tree := range allTrees {
		if i == 0 {
			t = &Tree{"blk", []*Tree{tree}}
			last = t
		} else {
			new = &Tree{"blk", []*Tree{tree}}
			last.Sons = append(last.Sons, new)
			last = new
		}
	}
	return t, last
}

func constructDeclarationSequence(first, rest interface{}) []*Declaration {
	firstConv := first.([]*Declaration)
	restConv := rest.([]*Declaration)
	allDeclarations := append(firstConv, restConv...)
	return allDeclarations
}

func constructProcedure(id, formals, blk interface{}) *Tree {
	procTree := &Tree{"prc", initSons()}
	idTree := &Tree{id.(string), initSons()}
	var formalsTree *Tree
	if formals == nil {
		formalsTree = &Tree{"for", initSons()}
	} else {
		formalsTree = formals.(*Tree)
	}
	blkTree := blk.(*Tree)
	procTree.Sons = append(procTree.Sons, idTree)
	procTree.Sons = append(procTree.Sons, formalsTree)
	procTree.Sons = append(procTree.Sons, blkTree)
	return procTree
}

func constructFormals(first, rest interface{}) *Tree {
	formalsTree := &Tree{"for", initSons()}
	firstTree := &Tree{first.(string), initSons()}
	formalsTree.Sons = append(formalsTree.Sons, firstTree)

	for _, formalWithComma := range toIfaceSlice(rest) {
		formalName := toIfaceSlice(formalWithComma)[2].(string)
		formalTree := &Tree{formalName, initSons()}
		formalsTree.Sons = append(formalsTree.Sons, formalTree)
	}

	return formalsTree
}
