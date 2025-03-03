package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// tokenize (lexer) - convert input string into tokens (object of type and value) - get_next_token  (position s key here) - (text, pos)

// expr - veries of series of token matches corresponding to a grammar rule

// lexer

// csv = expr
// expr = val,val
// val = primitive data type (string, number(int, float), boolean), identifier (variable name)
// , = delimiter

/*


type TokenType int

const (
 STRING TokenType = iota
 IDENTIFIER
 INTEGER
 COMMA
 NEWLINE
 EOF
)


type Lexer struct {
 text string
 pos int
}

func (l *Lexer) get_next_token() Token {
 // get next token
 // switch or if else
 // determin token type
 // return Token{type: "STRING", value: "hello"}
}



program = expr
expr = val op val
val = int
op = +
int = 0-9

=> val + val

**/

type TokenType int

// TOKEN types
const (
	INTEGER TokenType = iota
	PLUS
	EOF
)

type ParseErr string

func (e ParseErr) Error() string {
	return string(e)
}

// Error
const PARSERROR = ParseErr("Error parsing input")

type Token struct {
	tokenType TokenType
	value     interface{}
}

type Lexer struct {
	text          string
	pos           int
	current_Token Token
}

// compare the current token type with the passed token
// type and if they match then "eat" the current token
// and assign the next token to the self.current_token,
// otherwise raise an exception.
func (l *Lexer) Eat(tokenType TokenType) error {
	if l.current_Token.tokenType == tokenType {
		if token, err := l.Get_Next_Token(); err == nil {
			l.current_Token = token
			return nil
		}

	}
	return fmt.Errorf("Error parsing input")
}

// skip whitespace
func (l *Lexer) SkipWhiteSpace() {
	if l.pos < len(l.text)-1 && unicode.IsSpace(rune(l.text[l.pos])) {
		l.pos++
	}
}

// Get next token
func (l *Lexer) Get_Next_Token() (Token, error) {

	l.SkipWhiteSpace()

	if l.pos > len(l.text)-1 {
		return Token{tokenType: EOF, value: nil}, nil
	}

	text := l.text

	current_char := rune(text[l.pos])

	if unicode.IsNumber(current_char) {
		start := l.pos
		for l.pos < len(l.text)-1 && unicode.IsNumber(rune(l.text[l.pos])) {
			l.pos++
		}

		parsedVal, err := strconv.Atoi(string(l.text[start:l.pos]))
		if err != nil {
			return Token{}, fmt.Errorf("failed to parse integer: %v", text[start:l.pos])
		}
		return Token{tokenType: INTEGER, value: parsedVal}, nil
	}

	if string(current_char) == "+" {
		l.pos += 1
		return Token{tokenType: PLUS, value: string(current_char)}, nil
	}

	return Token{}, fmt.Errorf("Error passing input: %q", current_char)
}

// Expression => check if it follows grammer
func (l *Lexer) Expr() (int, error) {
	token, err := l.Get_Next_Token()
	if err != nil {
		return 0, err
	}

	l.current_Token = token

	left := l.current_Token
	if err := l.Eat(INTEGER); err != nil {
		return 0, err
	}

	// expects a plus
	if err := l.Eat(PLUS); err != nil {
		return 0, err
	}

	right := l.current_Token
	if err := l.Eat(INTEGER); err != nil {
		return 0, err
	}

	result := left.value.(int) + right.value.(int)

	return result, nil

}
