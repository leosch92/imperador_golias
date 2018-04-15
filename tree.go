package main

import (
	"fmt"
	"io"
	"os"
)

const sonsLength = 3

type Tree struct {
	Value string
	Sons  [sonsLength]*Tree
}

func initSons() [sonsLength]*Tree {
	var sons [sonsLength]*Tree
	return sons
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

func (t *Tree) insert(Value string, id int) *Tree {
	t.Sons[id] = &Tree{Value, initSons()}
	return t.Sons[id]
}

func printer(w io.Writer, tree *Tree, ns int, ch rune) {
	if tree == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, tree.Value)
	printer(w, tree.Sons[0], ns+2, '0')
	printer(w, tree.Sons[1], ns+2, '1')
	printer(w, tree.Sons[2], ns+2, '2')
}

func printTree(tree *Tree) {
	printer(os.Stdout, tree, 0, 'R')
}

/*func printBplc(t *Tree) {

	if t == nil {
		return
	}

	if (t.Left == nil) && (t.Right == nil) {
		fmt.Print(t.Value)
	} else {
		fmt.Print(t.Value)
		fmt.Print("(")
		printBplc(t.Left)
		fmt.Print(", ")
		printBplc(t.Right)
		fmt.Print(")")
	}

}*/
