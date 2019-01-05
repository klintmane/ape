package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ape-lang/ape/src/compiler/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Ape 0.0.1 (%s)\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
