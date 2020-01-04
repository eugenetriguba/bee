package main

import (
	"fmt"
	"os"
	"os/user"

	"bee/repl"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Bee programming language!\n", currentUser.Username)
	fmt.Printf("Feel free to type in commands. Enter 'exit' to leave the repl\n")
	repl.Start(os.Stdin, os.Stdout)
}
