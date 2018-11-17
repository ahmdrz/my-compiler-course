package parser

import (
	"errors"
	"fmt"

	"github.com/ahmdrz/my-compiler-course/lexer"
)

/*
// block => { dcls, stmts }
block => { stmts }
// dcls => dcl, dcls
// dcl => id := idlist ;
idlist => id, idlist | id
stmts => stmt stmts | landa
stmt => for-stmt | if-stmt | while-stmt | assign-stmt | block
while-stmt => while (EXPR) stmt
if-stmt => if (EXPR) stmt | if (EXPR) stmt else stmt
for-stmt => for (EXPR ; EXPR ; EXPR) stmt
assign-stmt => id = EXPR;
EXPR => EXPR + EXPR | EXPR * EXPR | EXPR / EXPR | EXPR - EXPR | (EXPR) | id | num
*/

const (
	typeMatcher = uint8(iota)
	textMatcher
)

type matcher struct {
	mode  uint8
	value []string
}

func (m matcher) Match(p *Parser) bool {
	if m.mode == typeMatcher {
		return p.checkType(m.value...)
	} else if m.mode == textMatcher {
		return p.checkText(m.value...)
	}
	return false
}

type Parser struct {
	lookahead *lexer.Token
	lexer     *lexer.Lexer
}

func (p *Parser) line() string {
	return fmt.Sprintf("line %03d, ", p.lexer.Line())
}

type Tokens []lexer.Token

func NewParser(lex *lexer.Lexer) (*Parser, error) {
	parser := &Parser{
		lexer: lex,
	}
	parser.next()
	if parser.lookahead == nil {
		return nil, errors.New("there is no token in input")
	}
	return parser, nil
}

func (p *Parser) Parse() error {
	return p.main()
}

func (p *Parser) checkType(inputs ...string) bool {
	for _, t := range inputs {
		if p.lookahead.Type == t {
			return true
		}
	}
	return false
}

func (p *Parser) checkText(inputs ...string) bool {
	for _, t := range inputs {
		if p.lookahead.Text == t {
			return true
		}
	}
	return false
}

func (p *Parser) next() {
	p.lookahead = p.lexer.Next()
}

func (p *Parser) statement() (err error) {
	if p.checkText("if", "while") {
		p.next()
		if !p.checkText("(") {
			return errors.New(p.line() + "statement: ( not found")
		}
		err = p.expression()
		if err != nil {
			return err
		}
		if !p.checkText(")") {
			return errors.New(p.line() + "statement: ) not found")
		}
		p.next()
		return p.block()
	}
	if p.checkText("for") {
		p.next()
		if !p.checkText("(") {
			return errors.New(p.line() + "statement: ( not found")
		}
		p.next()
		err = p.assignment()
		if err != nil {
			return err
		}
		if !p.checkText(";") {
			return errors.New(p.line() + "statement: ; not found")
		}
		p.next()
		err = p.expression()
		if err != nil {
			return err
		}
		if !p.checkText(";") {
			return errors.New(p.line() + "statement: ; not found")
		}
		p.next()
		err = p.assignment()
		if err != nil {
			return err
		}
		if !p.checkText(")") {
			return errors.New(p.line() + "statement: ) not found")
		}
		p.next()
		return p.block()
	}
	return errors.New(p.line() + "statement: unknown statement")
}

func (p *Parser) assignment() (err error) {
	if p.checkType(lexer.Identifier) {
		p.next()

		if !p.checkType(lexer.Declaration) {
			return errors.New(p.line() + "statements: = or := mismatched")
		}

		err = p.expression()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) statements() (err error) {
	if p.checkType(lexer.Statement) {
		err = p.statement()
		if err != nil {
			return err
		}
		return p.statements()
	}
	if p.checkType(lexer.Identifier) {
		p.next()

		if !p.checkType(lexer.Declaration) {
			return errors.New(p.line() + "statements: = or := mismatched")
		}

		err = p.expression()
		if err != nil {
			return err
		}

		if !p.checkText(";") {
			return errors.New(p.line() + "statements: ; forgotten")
		}
		p.next()
		return p.statements()
	}
	return nil
}

func (p *Parser) expression() (err error) {
	p.next()
	if p.checkType(lexer.Digits, lexer.Identifier) {
		return p.expression()
	}
	if p.checkText("(") {
		err = p.expression()
		if err != nil {
			return err
		}
		if p.checkText(")") {
			p.next()
			return nil
		}
	}
	if p.checkType(lexer.MathematicalSymbol, lexer.Operator) {
		return p.expression()
	}
	return nil
}

func (p *Parser) main() (err error) {
	if p.lookahead.Text == "main" {
		p.next()
		return p.block()
	}
	return errors.New(p.line() + "main: main keyword mismatched")
}

func (p *Parser) block() (err error) {
	if p.lookahead.Text == "{" {
		p.next()

		err = p.statements()
		if err != nil {
			return err
		}

		if p.lookahead.Text == "}" {
			p.next()
			return nil
		}
		return errors.New(p.line() + "error, } mismatched in block")
	}
	return errors.New(p.line() + "error, { mismatched block")
}

// func (p *Parser) declaration() (err error) {
// 	p.next()

// 	if !p.checkType(lexer.Declaration) {
// 		return errors.New("declaration: = or := mismatched")
// 	}

// 	if err = p.expression(); err != nil {
// 		return err
// 	}

// 	if !p.checkText(";") {
// 		return errors.New("declaration: ; mismatched")
// 	}

// 	p.next()
// 	return nil
// }

// func (p *Parser) declarations() (err error) {
// 	if p.checkType(lexer.Identifier) {
// 		p.declaration()
// 		p.declarations()
// 	}
// 	return nil
// }
