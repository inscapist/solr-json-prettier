package ast

import (
	"errors"
	"strings"
)

var (
	ErrMismatchBrackets = errors.New("Mismatch brackets")
)

func Build(code string, maxLineLength int) (*Node, error) {
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
			var node *Node
			for _, token := range tokenizedLabel(buf, maxLineLength) {
				node = &Node{Label: token}
				if last != nil {
					last.Children = append(last.Children, node)
				}
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
				for _, token := range tokenizedLabel(buf, maxLineLength) {
					last.Children = append(last.Children, &Node{Label: token})
				}
			}
			nodes.Pop()
			buf = nil
		default:
			buf = append(buf, c)
		}
	}
	return last, nil
}

func tokenizedLabel(buf []rune, maxLineLength int) []string {
	label := string(buf)
	if len(label) > maxLineLength {
		return strings.Fields(label)
	}
	return []string{label}
}
