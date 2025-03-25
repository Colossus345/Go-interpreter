package main

import (
	"fmt"
	"github.com/Colossus345/Go-interpreter/internal/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)

}
