package main

type TokenType int

// csv = ama,yam
//     = ama, 2
//     = ama, "ama"
//     = "ama", "ama"
//     =

// enum
const (
	COMMA TokenType = iota
	LQT
	RQT
	TEXT
)

// token
type Token struct {
	token_type TokenType
	value      interface{}
}

// lexer
type Lexer struct {
	text          string
	pos           int
	current_char  rune
	current_token Token
}
