package lexer

import (
	"bufio"
	"io"
	"regexp"
	"unicode"
)

const (
	delimeter   = "dlmt"
	statement   = "stmt"
	operator    = "oprt"
	digits      = "dgts"
	declaration = "dcln"
)

var digitRegex = regexp.MustCompile(`(\d+)(.|e(-|\+))\d+`)

var keywords = map[string]string{
	"if":  statement,
	"for": statement,
	"(":   delimeter,
	")":   delimeter,
	";":   delimeter,
	"{":   delimeter,
	"}":   delimeter,
	">":   operator,
	"<":   operator,
	">=":  operator,
	"<=":  operator,
	"=":   operator,
	"<>":  operator,
	":=":  declaration,
}

type Token struct {
	Line int
	Text string
	Type string
}

type Lexer struct {
	Reader *bufio.Reader
	line   int
}

func NewLexer() *Lexer {
	return &Lexer{line: 1}
}

func (l *Lexer) makeToken(text []rune) Token {
	t := string(text)
	typeOfText, ok := keywords[t]
	if !ok {
		if digitRegex.MatchString(t) {
			typeOfText = digits
		} else {
			typeOfText = "unknown"
		}
	}
	return Token{Line: l.line, Text: string(text), Type: typeOfText}
}

func (l *Lexer) isDelimeter(ch rune) bool {
	typeOfRune, exists := keywords[string(ch)]
	if !exists {
		return false
	}
	return typeOfRune == delimeter
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

func (l *Lexer) Tokenizer() chan Token {
	var val []rune
	var comment bool
	output := make(chan Token)

	go func() {
		for {
			ch, _, err := l.Reader.ReadRune()
			if err != nil {
				if len(val) > 0 {
					output <- l.makeToken(val)
					val = []rune{}
					continue
				}
				if err == io.EOF {
					break
				}
				panic(err)
			}

			if unicode.IsSpace(ch) {
				if ch == '\n' {
					l.line++
					comment = false
				}
				if len(val) > 0 {
					output <- l.makeToken(val)
					val = []rune{}
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
					output <- l.makeToken(val)
					val = []rune{}
				}
				output <- l.makeToken([]rune{ch})
				continue
			}

			val = append(val, ch)
		}
		close(output)
	}()

	return output
}
