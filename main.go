package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/weiyouwozuiku/Pangu/consts"
	"github.com/weiyouwozuiku/Pangu/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\nHello %s! This is the Pangu programming language!\n",
		consts.Pangu, user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
