package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/sagittaros/solr-syntax/pkg/ast"
)

func main() {
	var maxLineLen int
	flag.IntVar(&maxLineLen, "max-length", 20, "Maximum length of a line")
	flag.IntVar(&maxLineLen, "m", 20, "Maximum length of a line (shorthand)")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	tree, err := ast.Build(string(text), maxLineLen)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	tree.Print("  ")
}
