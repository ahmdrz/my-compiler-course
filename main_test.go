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
		return
	}

	lexer := lx.NewLexer()
	err = lexer.Load(f)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Reading example file successfully finished!")
}
