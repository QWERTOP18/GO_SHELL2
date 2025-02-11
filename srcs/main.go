package main

import (
	"fmt"
	"os"
	"os/user"
	"shell/executor"
	"shell/repl"
	"strings"
)

func main() {
	if len(os.Args) == 3 && os.Args[1] == "-c" {
		args := strings.Split(os.Args[2], " ")
		os.Exit(executor.ExecSimpleCommandSync(args, os.Stdin, os.Stdout))
	}
	if len(os.Args) == 2 {
		// read from file
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s!\n", user.Username)
	repl.Start()
}

// func main() {
// 	if len(os.Args) < 2 {
// 		println("Usage: go run main.go [command]")
// 		os.Exit(1)
// 	}
// 	args := strings.Split(os.Args[1], " ")
// 	os.Exit(executor.ExecSimpleCommandSync(args, os.Stdin, os.Stdout))
// }
