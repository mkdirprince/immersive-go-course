package main

// grammer for csv
// csv ::= row(NEWLINE row)*
// row ::= field(COMMA field)*
// field ::= text | "text"
// text ::= [any char except QOUTES, NEWLINE or COMMA]*

// Functiosn and Methods
// parse-csv() Function
// row() - method
// field() - method
// text() - method

// alternative (text | "text")
// if text (no-qoutes)-else if "text" (qouted)

// optional groudping
// (NEWLINE ROW)* && (COMMA FIELD)*
// while loop

// needs
// lexer with methods
// NewLexer => initial start for the parse
// skipwhitespace => skip white spaces
// advance => move the pos of the lexer
// get_next_token => to get the next token
// csv (parse-csv) => parse rows line by line
// eat => checks if the token type is correct and advance
// row => parses fields seperated by comma
// field => parses and return texts (eg. ama or "ama")
// text(terminal) => characters execpt qoutes, comma and newline
