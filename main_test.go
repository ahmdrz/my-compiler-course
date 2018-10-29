package main

import (
	"os"
	"testing"

	lx "github.com/ahmdrz/my-compiler-course/lexer"
)

func TestMain(t *testing.T) {
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
