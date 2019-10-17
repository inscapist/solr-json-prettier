package ast

import (
	"fmt"
	"strings"
)

type Node struct {
	Label            string
	Children         []*Node
	ChildrenBrackets [2]rune
}

func (n *Node) Print(baseIndent string) {
	// encoder := json.NewEncoder(os.Stdout)
	// encoder.Encode(n)
	// fmt.Printf("\n\n")
	n.print(baseIndent, 0)
}

func (n *Node) print(baseIndent string, depth int) {
	indent := strings.Repeat(baseIndent, depth)
	printIndent(n.Label, indent)
	if len(n.Children) > 0 {
		printIndent(string(n.ChildrenBrackets[0]), indent)
		for _, child := range n.Children {
			child.print(baseIndent, depth+1)
		}
		printIndent(string(n.ChildrenBrackets[1]), indent)
	}
}

func printIndent(text, indent string) {
	if len(text) > 0 {
		fmt.Print(indent)
		fmt.Println(strings.TrimSpace(text))
	}
}
