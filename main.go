package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/phoorichet/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s This is Monkey programming\n", user)
	fmt.Printf("Type any command\n")
	repl.Start(os.Stdin, os.Stdout)
}
