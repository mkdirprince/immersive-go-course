package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("calc> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			break
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if input == "exit" {
			break
		}

		lexer := Lexer{text: input, pos: 0, current_token: Token{}, current_char: rune(input[0])}
		result, err := lexer.Expr()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		fmt.Println(result)
	}

}
