package repl

import (
	"ape/src/compiler/compiler"
	"ape/src/compiler/symbols"
	"ape/src/compiler/vm"
	"ape/src/data"
	"ape/src/lexer"
	"ape/src/parser"
	"bufio"
	"fmt"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	constants := []data.Data{}
	globals := make([]data.Data, vm.GLOBALS_SIZE)
	symbols := symbols.New()

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

		comp := compiler.NewWithState(symbols, constants)
		err := comp.Compile(program)

		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n Error: %s\n", err)
			continue
		}

		bytecode := comp.Bytecode()
		constants = bytecode.Constants

		machine := vm.NewWithGlobals(bytecode, globals)
		err = machine.Run()

		if err != nil {
			fmt.Fprintf(out, "Execution failed:\n Error: %s\n", err)
			continue
		}

		result := machine.Result()

		io.WriteString(out, result.Inspect())
		io.WriteString(out, "\n")
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
