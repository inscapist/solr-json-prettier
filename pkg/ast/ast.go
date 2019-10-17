package ast

import (
	"errors"
)

var (
	ErrMismatchBrackets = errors.New("Mismatch brackets")
)

func Build(code string) (*Node, error) {
	var (
		nodes    *NodeStack
		brackets *BracketStack
		last     *Node
		buf      []rune
	)
	nodes = &NodeStack{}
	brackets = &BracketStack{}
	for _, c := range code {
		last = nodes.Last()
		switch {
		case bracketContains(OpenBrackets, c):
			brackets.Push(c)
			node := &Node{Label: string(buf)}
			if last != nil {
				last.Children = append(last.Children, node)
			}
			nodes.Push(node)
			node.ChildrenBrackets[0] = c
			buf = nil
		case bracketContains(CloseBrackets, c):
			closeBracket := BracketPair[brackets.Pop()]
			if closeBracket != c {
				return nil, ErrMismatchBrackets
			}
			last.ChildrenBrackets[1] = c
			if len(buf) > 0 {
				last.Children = append(last.Children, &Node{Label: string(buf)})
			}
			nodes.Pop()
			buf = nil
		default:
			buf = append(buf, c)
		}
	}
	return last, nil
}
