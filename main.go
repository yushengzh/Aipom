// main.go

package main

import (
	"Aipom/repl"
	"fmt"
	"os"
	"os/user"
)

/**
 * 测试REPL
 */

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi, %s! Welcome to use AIPOM.", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
