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
	MUL
	MINUS
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
	current_token Token
	current_char  rune
}

// compare the current token type with the passed token
// type and if they match then "eat" the current token
// and assign the next token to the self.current_token,
// otherwise raise an exception.
func (l *Lexer) Eat(tokenType TokenType) error {
	if l.current_token.tokenType == tokenType {
		if token, err := l.Get_Next_Token(); err == nil {
			l.current_token = token
			return nil
		}

	}
	return fmt.Errorf("Error parsing input")
}

// advance position and set current char
func (l *Lexer) Advance() {
	l.pos += 1
	if l.pos > len(l.text)-1 {
		l.current_char = rune(0)
	} else {
		l.current_char = rune(l.text[l.pos])
	}
}

// skip whitespace
func (l *Lexer) SkipWhiteSpace() {
	for l.current_char != rune(0) && unicode.IsSpace(l.current_char) {
		l.Advance()
	}
}

// integer (multple integers)
func (l *Lexer) Integer() int {

	var result string
	for l.current_char != rune(0) && unicode.IsNumber(l.current_char) {
		result += string(l.current_char)
		l.Advance()
	}
	// Convert string to int
	val, _ := strconv.Atoi(result)
	return val
}

// Get next token
func (l *Lexer) Get_Next_Token() (Token, error) {
	for l.current_char != rune(0) {
		if unicode.IsSpace(l.current_char) {
			l.SkipWhiteSpace()
			continue
		}

		if unicode.IsNumber(l.current_char) {
			return Token{tokenType: INTEGER, value: l.Integer()}, nil
		}

		if l.current_char == '+' {
			l.Advance()
			return Token{tokenType: PLUS, value: l.current_char}, nil
		}

		if l.current_char == '-' {
			l.Advance()
			return Token{tokenType: MINUS, value: l.current_char}, nil
		}

		if l.current_char == '*' {
			l.Advance()
			return Token{tokenType: MUL, value: l.current_char}, nil
		}

		return Token{}, fmt.Errorf("invalid character: %c", l.current_char)
	}

	return Token{tokenType: EOF, value: nil}, nil
}

// Expression => check if it follows grammer
// expr -> INT PLUS INT
// expr -> INT MINUS INT
// expr -> INT MUL INT
func (l *Lexer) Expr() (int, error) {
	token, err := l.Get_Next_Token()
	if err != nil {
		return 0, err
	}

	l.current_token = token
	left := l.current_token

	if err := l.Eat(INTEGER); err != nil {
		fmt.Println("Error eating INTEGER:", err)
		return 0, err
	}

	// expects an operator
	op := l.current_token

	if op.tokenType == PLUS {
		if err := l.Eat(PLUS); err != nil {
			fmt.Println("Error eating PLUS:", err)
			return 0, err
		}
	} else if op.tokenType == MINUS {
		if err := l.Eat(MINUS); err != nil {
			fmt.Println("Error eating MINUS:", err)
			return 0, err
		}
	} else if op.tokenType == MUL {
		if err := l.Eat(MUL); err != nil {
			fmt.Println("Error eating MUL:", err)
			return 0, err
		}
	} else {
		return 0, fmt.Errorf("expected operator, got: %+v", op)
	}

	right := l.current_token
	if err := l.Eat(INTEGER); err != nil {
		fmt.Println("Error eating right INTEGER:", err)
		return 0, err
	}

	var result int
	if op.tokenType == PLUS {
		result = left.value.(int) + right.value.(int)
	} else if op.tokenType == MINUS {
		result = left.value.(int) - right.value.(int)
	} else {
		result = left.value.(int) * right.value.(int)
	}

	return result, nil
}
