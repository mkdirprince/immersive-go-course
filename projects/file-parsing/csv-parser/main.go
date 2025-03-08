package main

import (
	"bufio"
	"errors"
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

		lexer := NewLexer(input)
		result, err := lexer.Expr()
		if err != nil {
			if errors.Is(err, PARSERROR) {
				fmt.Println("Parsing error:", err)
			} else {
				fmt.Println("Unknown error:", err)
			}
		}

		fmt.Println(result)
	}

}
