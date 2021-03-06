package main

import (
	"fmt"
	"os"

	"github.com/kieron-pivotal/dojo-bowling/bowling"
)

func main() {
	usage := fmt.Sprintf("%s <throw string>", os.Args[0])
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	fmt.Println(bowling.NewGame(os.Args[1]).Score())
}
