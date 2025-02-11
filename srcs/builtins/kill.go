package builtins

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

// kill コマンドの実装

const sigList = "HUP INT QUIT ILL TRAP ABRT EMT FPE KILL BUS SEGV SYS PIPE ALRM TERM URG STOP TSTP CONT CHLD TTIN TTOU IO XCPU XFSZ VTALRM PROF WINCH INFO USR1 USR2"

func kill(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "kill: usage: kill [-signal] <pid> ...")
		return 1
	}

	sig := syscall.SIGTERM // デフォルトのシグナル
	argIndex := 0

	// 最初の引数が `-` で始まっている場合、シグナルとして解釈
	if args[0][0] == '-' {
		sigNum, err := strconv.Atoi(args[0][1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "kill: invalid signal: %s\n", args[0])
			fmt.Fprintf(os.Stderr, "kill: type %s\n", sigList)
			return 1
		}
		sig = syscall.Signal(sigNum)
		argIndex = 1
	}

	// プロセスIDを取得して kill
	for _, pidStr := range args[argIndex:] {
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "kill: %s: arguments must be process or job IDs\n", pidStr)
			return 1
		}

		// プロセスを kill する
		if err := syscall.Kill(pid, sig); err != nil {
			fmt.Fprintf(os.Stderr, "kill: failed to kill process %d: %v\n", pid, err)
			return 1
		}
	}

	return 0
}
