package main

import (
	"fmt"
	"io"
	"os"
)

type Tree struct {
	Left  *Tree
	Value string
	Right *Tree
}
//TODO: mudar o tipo que o Stack espera p/ Tree
type stack struct {
	data []string
}

func findMinMax(tree *Tree, min *int, max *int, hd int) {
	if tree == nil {
		return
	}

	if hd < *min {
		*min = hd
	}
	if hd > *max {
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

func (tree *Tree) print() {
	printer(os.Stdout, tree, 0, 'R')
}

func (tree *Tree) dismember() (*Tree, string, *Tree) {
	return tree.Left, tree.Value, tree.Right
}

//Implementação de Struct de Pilha-> utiliza slices para facilitar

func (s stack) push(info string) stack {
	s.data = append(s.data, info)
	return s
}

func (s stack) pop() (stack, string) {
	if len(s.data) == 0 {
		return s, ""
	}
	var info = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return s, info
}

func (s stack) size() int {
	return len(s.data)
}

func (s *stack) push_tree(tree *Tree) (*stack) {
    l_node, value, r_node := tree.dismember()
    s.push(&Tree(nil, value, nil))
    s.push(l_node)
    s.push(r_node)
    return s
}

func main() {
	tree := Tree{nil, "add", nil}
	tree.insertLeft("3")
	tree.insertRight("2")
	tree.print()
	//tree_l, value, tree_r := tree.dismember()
	s := stack{}
	r := s.push_tree(&tree)
	_, a := r.pop()
	_, b := r.pop()
	_, c := r.pop()
	fmt.Print(a.print())
	fmt.Print(b.print())
	fmt.Print(c.print())
}
