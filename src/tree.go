package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type Tree struct {
	Value string
	Sons  []*Tree
}

func initSons() []*Tree {
	return *new([]*Tree)
}

func (t *Tree) checkIfNode() bool {
	for _, son := range t.Sons {
		if son != nil {
			return false
		}
	}
	return true
}

func (t *Tree) toString() string {
	return t.Value
}

/*func findMinMax(tree *Tree, min *int, max *int, hd int) {
	if tree == nil {
		return
	}

	if hd < *min {
		*min = hd
	} else if hd > *max {
		*max = hd
	}

	findMinMax(tree.Left, min, max, hd-1)
	findMinMax(tree.Right, min, max, hd+1)
}*/

/*func (t *Tree) insert(Value string, id int) *Tree {
	t.Sons[id] = &Tree{Value, initSons()}
	return t.Sons[id]
}*/

func printer(w io.Writer, tree *Tree, ns int, sonId string) {
	if tree == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%s:%v\n", sonId, tree.Value)

	for j := 0; j < len(tree.Sons); j++ {
		printer(w, tree.Sons[j], ns+2, strconv.Itoa(j))
	}

}

func printTree(tree *Tree) {
	printer(os.Stdout, tree, 0, "R")
}

func printBplc(t *Tree) {

	if t == nil {
		return
	}
	leaf := true
	for _, son := range t.Sons {
		if !(son == nil) {
			leaf = false
		}
	}
	
	if leaf {
		fmt.Print(t.Value)
		return
	} else {
		fmt.Print(t.Value)
		fmt.Print("(")
		printBplc(t.Sons[0])
		for _, son := range t.Sons[1:] {
			fmt.Print(", ")
			printBplc(son)
		}
		fmt.Print(")")
	}

}
