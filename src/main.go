package main

import (
	"ape/src/interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Ape 0.0.1 (%s)\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
