package lexer

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"unicode"
)

const (
	Delimeter          = "dlmt"
	Statement          = "stmt"
	Operator           = "oprt"
	Digits             = "dgts"
	Declaration        = "dcln"
	MathematicalSymbol = "mtsl"
	Identifier         = "name"

	multiLineCommentStateZero = iota
	multiLineCommentStateOne
	multiLineCommentStateTwo
)

var digitRegex = regexp.MustCompile(`((\d+)|(.|e(-|\+))\d+)`)

var keywords = map[string]string{
	"if":    Statement,
	"for":   Statement,
	"while": Statement,
	"(":     Delimeter,
	")":     Delimeter,
	";":     Delimeter,
	"{":     Delimeter,
	"}":     Delimeter,
	">":     Operator,
	"<":     Operator,
	">=":    Operator,
	"<=":    Operator,
	"==":    Operator,
	"!=":    Operator,
	":=":    Declaration,
	"=":     Declaration,
	"-":     MathematicalSymbol,
	"+":     MathematicalSymbol,
	"*":     MathematicalSymbol,
	"/":     MathematicalSymbol,
}

type Token struct {
	Line int
	Text string
	Type string
}

func (t Token) String() string {
	return fmt.Sprintf("line %03d, text %s, type %s", t.Line, t.Text, t.Type)
}

type Lexer struct {
	Reader *bufio.Reader
	line   int
}

func (l *Lexer) Line() int {
	return l.line
}

func NewLexer() *Lexer {
	return &Lexer{line: 1}
}

func (l *Lexer) makeToken(text []rune, typeOfInput ...string) *Token {
	t := string(text)
	if len(typeOfInput) == 1 {
		return &Token{Line: l.line, Text: string(text), Type: typeOfInput[0]}
	}
	typeOfText, ok := keywords[t]
	if !ok {
		if digitRegex.MatchString(t) {
			typeOfText = Digits
		} else {
			typeOfText = Identifier
		}
	}
	return &Token{Line: l.line, Text: string(text), Type: typeOfText}
}

func (l *Lexer) isDelimeter(ch rune) bool {
	typeOfRune, exists := keywords[string(ch)]
	if !exists {
		return false
	}
	return typeOfRune == Delimeter
}

func (l *Lexer) Load(input io.Reader) error {
	reader := bufio.NewReader(input)
	firstChar, _, err := reader.ReadRune()
	if err != nil {
		return err
	}
	if firstChar != 0xFEFF { // BOM
		err := reader.UnreadRune()
		if err != nil {
			return err
		}
	}
	l.Reader = reader
	return nil
}

func (l *Lexer) Backward(t *Token) {

}

func (l *Lexer) Next() *Token {
	var (
		val          []rune
		comment      bool
		quoted       bool
		multiComment = multiLineCommentStateZero
	)

	for {
		ch, _, err := l.Reader.ReadRune()
		if err != nil {
			if len(val) > 0 {
				return l.makeToken(val)
			}
			if err == io.EOF {
				return nil
			}
			panic(err)
		}

		if string(val) == "/*" {
			multiComment = multiLineCommentStateOne
			val = []rune{}
		}

		if multiComment != multiLineCommentStateZero {
			if ch == '*' {
				multiComment = multiLineCommentStateTwo
			}
			if ch == '/' {
				multiComment = multiLineCommentStateZero
			}
			if ch == '\n' {
				l.line++
			}
			continue
		}

		if quoted {
			if ch == '"' {
				quoted = false
				return l.makeToken(val, Declaration)
			}
			if ch == '\n' {
				l.line++
			}
			val = append(val, ch)
			continue
		}

		if unicode.IsSpace(ch) {
			if ch == '\n' {
				l.line++
				comment = false
			}

			if len(val) > 0 {
				return l.makeToken(val)
			}
			continue
		}

		if ch == '#' {
			comment = true
		}

		if comment {
			continue
		}

		if l.isDelimeter(ch) {
			if len(val) > 0 {
				l.Reader.UnreadRune()
				return l.makeToken(val)
			}
			return l.makeToken([]rune{ch})
		}

		if len(val) == 0 {
			if ch == '"' {
				quoted = true
				continue
			}
		}

		val = append(val, ch)
	}
}
