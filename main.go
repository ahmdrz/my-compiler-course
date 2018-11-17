package main

import (
	"io"
	"log"
	"os"

	"github.com/ahmdrz/my-compiler-course/lexer"
	"github.com/ahmdrz/my-compiler-course/parser"
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

	myLexer := lexer.NewLexer()
	err = myLexer.Load(input)
	if err != nil {
		log.Fatalln(err)
	}

	myParser, err := parser.NewParser(myLexer)
	if err != nil {
		log.Fatalln(err)
	}

	err = myParser.Parse()
	if err != nil {
		log.Fatalln(err)
	}
}
