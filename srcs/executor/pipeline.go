package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec1pipe(words1, words2 []string) error {
	cmd1 := exec.Command(words1[0], words1[1:]...)
	cmd2 := exec.Command(words2[0], words2[1:]...)

	stdoutPipe, err := cmd1.StdoutPipe()
	if err != nil {
		panic(err)
	}

	// 標準入出力とエラー出力を設定
	cmd1.Stdin = os.Stdin
	cmd1.Stderr = os.Stderr
	cmd2.Stdin = stdoutPipe
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr

	if err := cmd1.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v: %v\n", words1[0], err)
	}
	if err := cmd2.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v: %v\n", words2[0], err)
	}

	cmd1.Wait()
	err = cmd2.Wait()
	fmt.Println("Pipe finished")
	return err
}

func ExecPipeline(lists [][]string) {
	// listsの長さに基づいてexecListsを初期化
	execLists := make([]*exec.Cmd, len(lists))

	// コマンドを実行可能なexec.Cmdに変換してexecListsにセット
	for i, list := range lists {
		execLists[i] = exec.Command(list[0], list[1:]...)
	}

	// パイプラインをつなげる
	for i := 0; i < len(execLists)-1; i++ {
		// execLists[i] の標準出力を execLists[i+1] の標準入力に接続
		stdoutPipe, err := execLists[i].StdoutPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error setting up pipe: %v\n", err)
			return
		}
		execLists[i+1].Stdin = stdoutPipe
	}
	//最初のコマンドと最後のコマンドの接続先
	execLists[0].Stdin = os.Stdin
	execLists[len(execLists)-1].Stdout = os.Stdout

	// すべてのコマンドを実行
	for _, execCommand := range execLists {
		if err := execCommand.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "command %s: %v\n", execCommand.Args[0], err)
			return
		}
	}

	// すべてのコマンドが終了するまで待機
	for _, execCommand := range execLists {
		if err := execCommand.Wait(); err != nil {
			fmt.Fprintf(os.Stderr, "command %s: %v\n", execCommand.Args[0], err)
		}
	}
}
