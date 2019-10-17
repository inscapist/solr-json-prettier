package main

import (
	"bufio"
	"fmt"
	"os"

	"gitlab.com/sagittaros/solr-syntax/pkg/ast"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	tree, err := ast.Build(string(text))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	tree.Print("  ")
}
