package executor

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"shell/builtins"
	"syscall"
)

// 呼び出し側でsignalはsetした方がいい
// todo 引数は取らずグローバル変数でfgのpgidにシグナルを送るようにする。
func setupSignal(cmd *exec.Cmd, ch chan struct{}) {
	// シグナルを受け取るチャネルを作成
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT)

	go func() {
		defer signal.Stop(sigchan)
		defer close(sigchan)

		for {
			select {
			case <-ch:
				return // 終了シグナルを受け取ったら goroutine を停止
			case sig := <-sigchan:
				if cmd.Process != nil {
					cmd.Process.Signal(sig)
				}
			}
		}
	}()
}

/*
親プロセスはコマンド実行中は無視、
exec.Commandを使用するとforkとexecvの間の処理が書けない

goroutine leakに気をつける
*/
func ExecSimpleCommandSync(words []string, inputFile *os.File, outputFile *os.File) int {
	if builtins.Lookupbuiltins(words[0]) {
		return builtins.Execbuiltins(words[0], words[1:], inputFile, outputFile)
	}
	cmd := exec.Command(words[0], words[1:]...)
	cmd.Stdin = inputFile
	cmd.Stdout = outputFile
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	ch := make(chan struct{})

	// シグナル転送用の goroutine を起動
	setupSignal(cmd, ch)

	err := cmd.Run()

	// goroutine を終了させる
	close(ch)
	if err != nil {
		if exitErr, isExitError := err.(*exec.ExitError); isExitError {
			// プロセスの終了コードを取得
			if waitStatus, isWaitStatus := exitErr.Sys().(syscall.WaitStatus); isWaitStatus {
				return waitStatus.ExitStatus()
			}
		}
		fmt.Fprintf(os.Stderr, "Error: %v: %v\n", words[0], "command not found")
		return 127
	}
	return 0
}

func ExecSimpleCommandAsync(words []string, inputFile *os.File, outputFile *os.File) int {

	cmd := exec.Command(words[0], words[1:]...)
	if inputFile != os.Stdin {
		cmd.Stdin = inputFile
	}
	cmd.Stdout = outputFile
	cmd.Stderr = os.Stderr

	//process groupを新たに作成
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	// goroutine を終了させる
	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v: %v\n", words[0], "command not found")
		return 127
	}
	pid := cmd.Process.Pid

	// syscall.Getpgid で GID を取得
	gpid, err := syscall.Getpgid(pid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get PGID for PID %d\n", pid)
		gpid = -1 // 取得に失敗した場合
	}

	fmt.Println("pid: ", pid, " gpid: ", gpid)

	go func() {
		_ = cmd.Wait() // Wait() のエラーは無視
	}()
	return 0
}

/*
signalを受け取るためにはSimple Commandも非同期的に処理しないといけない
builtinsの処理も
*/

// ⬇️readlineがshellの中でshellを起動したときに機能してくれない??
// func ExecSimpleCommandSync(words []string, inputFile *os.File, outputFile *os.File) int {

// 	if builtins.Lookupbuiltins(words[0]) {
// 		return builtins.Execbuiltins(words[0], words[1:], inputFile, outputFile)
// 	}

// 	cmd := exec.Command(words[0], words[1:]...)
// 	cmd.Stdin = inputFile
// 	cmd.Stdout = outputFile
// 	cmd.Stderr = os.Stderr
// 	cmd.SysProcAttr = &syscall.SysProcAttr{
// 		Setpgid: true,
// 	}
// 	//Linux プログラミングインターフェースp616
// 	// signal.Ignore(syscall.SIGINT, syscall.SIGQUIT)
// 	// defer signal.Reset(syscall.SIGINT, syscall.SIGQUIT)

// 	ch := make(chan struct{})

// 	// シグナル転送用の goroutine を起動
// 	//setupSignal(cmd, ch)

// 	// コマンド実行
// 	err := cmd.Run()

//		// goroutine を終了させる
//		close(ch)
//		if err != nil {
//			if exitErr, isExitError := err.(*exec.ExitError); isExitError {
//				// プロセスの終了コードを取得
//				if waitStatus, isWaitStatus := exitErr.Sys().(syscall.WaitStatus); isWaitStatus {
//					return waitStatus.ExitStatus()
//				}
//			}
//			fmt.Fprintf(os.Stderr, "Error: %v: %v\n", words[0], "command not found")
//			return 127
//		}
//		return 0
//	}
