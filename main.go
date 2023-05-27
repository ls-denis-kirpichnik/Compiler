package main

import (
	"compiler/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! such wow\n",
		user.Username)
	fmt.Printf("gimme input\n")
	repl.Start(os.Stdin, os.Stdout)
}
