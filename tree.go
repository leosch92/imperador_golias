package main

import(
	"fmt"
	"io"
	"os"
)

type Tree struct {
	Left  *Tree
	Value string
	Right *Tree
}

func findMinMax(tree *Tree, min *int, max *int, hd int) {
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
}

func (t *Tree) insertLeft(Value string) *Tree {
	t.Left = &Tree{nil, Value, nil}
	return t.Left
}

func (t *Tree) insertRight(Value string) *Tree {
	t.Right = &Tree{nil, Value, nil}
	return t.Right
}

func printer(w io.Writer, tree *Tree, ns int, ch rune) {
	if tree == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, tree.Value)
	printer(w, tree.Left, ns+2, 'E')
	printer(w, tree.Right, ns+2, 'D')
}

func printTree(tree *Tree) {
	printer(os.Stdout, tree, 0, 'R')
}