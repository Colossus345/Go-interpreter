package repl

import (
	"bufio"
	"fmt"
	"inter-median/internal/evaluator"
	"inter-median/internal/lexer"
	"inter-median/internal/parser"
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
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, program.String())
			io.WriteString(out, "\n")
		}

	}

}
func printErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
