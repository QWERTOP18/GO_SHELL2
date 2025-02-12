package builtins

import (
	"fmt"
	"os"
	"strings"
)

func export(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "export: usage: export VAR=value...")
		return 1
	}

	for _, arg := range args[1:] {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "export: %s: invalid assignment\n", arg)
			return 1
		}

		key := parts[0]
		value := parts[1]

		os.Setenv(key, value)
	}

	return 0
}

func unset(args []string) int {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "unset: usage: unset VAR...")
		return 1
	}

	for _, arg := range args {
		os.Unsetenv(arg)
	}

	return 0
}

// func alias(args []string) int {
// 	if len(args) < 1 {
// 		fmt.Fprintln(os.Stderr, "alias: usage: alias NAME=VALUE...")
// 		return 1
// 	}

// 	for i := 1; i < len(args); i += 2 {

// 	}

// 	return 0
// }

// func unalias(args []string) int {
// 	if len(args) < 1 {
// 		fmt.Fprintln(os.Stderr, "unalias: usage: unalias NAME...")
// 		return 1
// 	}

// 	for _, arg := range args {

// 	}

// 	return 0
// }

// func declare(args []string) int {
// 	if len(args) < 2 {
// 		fmt.Fprintln(os.Stderr, "declare: usage: declare [-aArfiRsxt] [-p] [NAME[=VALUE] ...]...")
// 		return 1
// 	}
// 	return 0
// }
