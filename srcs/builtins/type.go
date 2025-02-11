package builtins

import (
	"fmt"
	"os"
	"os/exec"
)

func _type(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "type: usage: type <command>")
		return 1
	}
	for _, v := range args {
		if Lookupbuiltins(v) {
			fmt.Printf("%s is a shell builtin\n", v)
		} else {
			path, err := exec.LookPath(v)
			if err == nil {
				fmt.Printf("%s is %s\n", v, path)
			} else {
				fmt.Printf("%s is not found\n", v)
			}
		}
	}
	return 0
}
