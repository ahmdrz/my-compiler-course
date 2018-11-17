package lexer

import (
	"bytes"
	"testing"
)

func TestNewLexer(t *testing.T) {
	lexer := NewLexer()
	err := lexer.Load(bytes.NewBufferString(`
	### Hello World
	a := 100;
	b := 200;
	c := a + b;
	`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("String reader: OK !")
}

func TestTokenizer(t *testing.T) {
	lexer := NewLexer()
	err := lexer.Load(bytes.NewBufferString(`	
	a := 100;
	`))
	if err != nil {
		t.Fatal(err)
	}
	variableDetector := false

	for {
		token := lexer.Next()
		if token == nil {
			break
		}
		if token.Text == "a" {
			variableDetector = true
			break
		}
	}
	if !variableDetector {
		t.Log("Did not detect variable !")
	} else {
		t.Log("Variable detected !")
	}
}
