package repl

import (
	"bufio"
	"fmt"
	"inter-median/internal/lexer"
	"inter-median/internal/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		for toke := l.NextToken(); toke.Type != token.EOF; toke = l.NextToken() {
			fmt.Printf("%+v\n", toke)
		}

	}

}
