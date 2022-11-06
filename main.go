package main

import (
	"fmt"
	"os/user"

	"github.com/panglang/pangu/consts"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Pangu programming language!\n",
		user.Username)
	fmt.Println(consts.Pangu)
}
