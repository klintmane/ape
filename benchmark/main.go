package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ape-lang/ape/src/compiler/compiler"
	"github.com/ape-lang/ape/src/compiler/vm"
	"github.com/ape-lang/ape/src/data"
	"github.com/ape-lang/ape/src/interpreter/eval"
	"github.com/ape-lang/ape/src/lexer"
	"github.com/ape-lang/ape/src/parser"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")
var input = `
	let fibonacci = fn(x) {
		if (x == 0) {
			0
		} else {
			if (x == 1) {
				return 1;
			} else {
				fibonacci(x - 1) + fibonacci(x - 2);
			}
		}
	};
	fibonacci(30);
`

func main() {
	flag.Parse()
	var duration time.Duration
	var result data.Data

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if *engine == "vm" {
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Printf("compiler error: %s", err)
			return
		}

		machine := vm.New(comp.Bytecode())
		start := time.Now()
		err = machine.Run()
		if err != nil {
			fmt.Printf("vm error: %s", err)
			return
		}

		duration = time.Since(start)
		result = machine.Result()
	} else {
		env := data.NewEnvironment()
		start := time.Now()
		result = eval.Eval(program, env)
		duration = time.Since(start)
	}

	fmt.Printf("engine=%s, result=%s, duration=%s\n", *engine, result.Inspect(), duration)
}
