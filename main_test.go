package main

import (
	"os"
	"testing"

	lx "github.com/ahmdrz/my-compiler-course/lexer"
)

func TestFile(t *testing.T) {
	f, err := os.Open("example.mo")
	if err != nil {
		t.Fatal(err)
	}

	lexer := lx.NewLexer()
	err = lexer.Load(f)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Reading example file successfully finished!")
}

func TestFileTokenizer(t *testing.T) {
	f, err := os.Open("example.mo")
	if err != nil {
		t.Fatal(err)
	}

	lexer := lx.NewLexer()
	err = lexer.Load(f)
	if err != nil {
		t.Fatal(err)
	}

	for token := range lexer.Tokenizer() {
		t.Logf("Line %03d Token %-18s Type %s", token.Line, token.Text, token.Type)
	}
}
