package repl

import (
	"ape/src/interpreter/data"
	"ape/src/interpreter/eval"
	"ape/src/lexer"
	"ape/src/parser"
	"bufio"
	"fmt"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := data.NewEnvironment()

	for {
		prompt()

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

		evaluated := eval.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Input could not be parsed!\n")
	io.WriteString(out, " Errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func prompt() {
	fmt.Printf(">> ")
}
