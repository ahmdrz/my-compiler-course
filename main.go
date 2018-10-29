package main

import (
	"fmt"
	"io"
	"log"
	"os"

	lx "github.com/ahmdrz/my-compiler-course/lexer"
)

func main() {
	var err error
	var input io.Reader = os.Stdin

	if len(os.Args) > 0 {
		input, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalln(err)
		}
	}

	lexer := lx.NewLexer()
	err = lexer.Load(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for token := range lexer.Tokenizer() {
		fmt.Printf("In line %03d token %-15s type %s\n", token.Line, token.Text, token.Type)
	}
}
