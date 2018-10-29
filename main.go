package main

import (
	"fmt"
	"os"

	lx "github.com/ahmdrz/my-compiler-course/lexer"
)

func main() {
	lexer := lx.NewLexer()
	err := lexer.Load(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for token := range lexer.Tokenizer() {
		fmt.Printf("In line %03d token %-15s type %s\n", token.Line, token.Text, token.Type)
	}
}
