package builtins

import (
	"fmt"
	"os"
	"strconv"
)

var builtinFuncs map[string]func([]string) int

// initはmainよりも先に呼ばれる
func init() {
	builtinFuncs = map[string]func([]string) int{
		"pwd":  pwd,
		"cd":   cd,
		"exit": exit,
		"kill": kill,
		"jobs": jobs,
		"type": _type,
		"fg":   fg,
		"bg":   bg,
	}
}

func Lookupbuiltins(name string) bool {
	_, exists := builtinFuncs[name]
	return exists
}

func Execbuiltins(name string, words []string, inputFile *os.File, outputFile *os.File) (exitStatus int) {
	// 現在の標準入力と標準出力を保存
	originalStdin := os.Stdin
	originalStdout := os.Stdout

	os.Stdin = inputFile
	os.Stdout = outputFile

	exitStatus = builtinFuncs[name](words)

	// 標準入力と標準出力を元に戻す
	os.Stdin = originalStdin
	os.Stdout = originalStdout

	return
}

func pwd(words []string) int {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: pwd:%v\n", err)
		return 1
	}
	fmt.Println(cwd)
	return 0
}

func cd(words []string) int {
	if len(words) != 1 {
		fmt.Fprintf(os.Stderr, "Error: cd: too many arguments\n")
		return 1
	}
	err := os.Chdir(words[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cd: %v\n", err)
		return 1
	}
	return 0
}

func exit(words []string) int {
	if len(words) == 0 {
		os.Exit(0)
	}
	if len(words) != 1 {
		fmt.Fprintf(os.Stderr, "Error: exit: too many arguments\n")
		return 1
	}
	status, err := strconv.Atoi(words[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: exit: %s: inumeric argument required\n", words[0])
		os.Exit(255)
	}
	os.Exit(status % 256)
	return 0
}
