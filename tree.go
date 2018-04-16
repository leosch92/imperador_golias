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