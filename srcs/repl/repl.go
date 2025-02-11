package repl

import (
	"fmt"
	"os"
	"runtime"
	"shell/executor"
	"strings"

	"github.com/chzyer/readline"
)

const PS1 = "🐠$ "

func Start() {
	rl, err := readline.New(PS1)
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt {
				// Ctrl-C が押されたら次のプロンプトを表示する
				continue
			}

			fmt.Println("exit: ", err)
			return
		}
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		// コマンドの実行
		var exitStatus int
		if args[len(args)-1] == "&" {
			exitStatus = executor.ExecSimpleCommandAsync(args[:len(args)-1], os.Stdin, os.Stdout)
		} else {
			exitStatus = executor.ExecSimpleCommandSync(args, os.Stdin, os.Stdout)
		}
		fmt.Println("Exit Status: ", exitStatus, "  Goroutine: ", runtime.NumGoroutine())
	}
}
