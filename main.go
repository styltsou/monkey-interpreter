package main

import (
	"fmt"
	"os/user"

	"github.com/styltsou/monkey-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type commands.")

	repl.Start()
}
