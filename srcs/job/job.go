package job

import (
	"io"
	"os/exec"
	"time"
	//"github.com/codegangsta/inject"
)

type Job struct {
	cmds    []*exec.Cmd       // 実行中のコマンド
	dir     string            // コマンドが実行されるディレクトリ
	started bool              // ジョブが開始されたかどうか
	Env     map[string]string // 環境変数
	Stdin   io.Reader         // 標準入力
	Stdout  io.Writer         // 標準出力
	Stderr  io.Writer         // 標準エラー
	ShowCMD bool              // デバッグ用: コマンドの詳細を表示
	timeout time.Duration     // タイムアウトの設定

	CmdName   string // コマンドの名前
	IsRunning bool   // ジョブが実行中かどうか
}

var Jobs []Job

func Foreground(jobID int) error {
	// if jobID < 0 || jobID >= len(Jobs) {
	//     return ErrJobNotFound
	// }

	// job := &Jobs[jobID]
	// job.IsRunning = true
	// job.Start()

	return nil
}

func Background(jobID int) error {
	// if jobID < 0 || jobID >= len(Jobs) {
	//     return ErrJobNotFound
	// }

	// job := &Jobs[jobID]
	// job.IsRunning = true
	// go job.Start()

	return nil
}

// type Cmd struct {
// 	Path string
// 	Args []string
// 	Env []string
// 	Dir string
// 	Stdin io.Reader
// 	Stdout io.Writer
// 	Stderr io.Writer
// 	ExtraFiles []*os.File
// 	SysProcAttr *syscall.SysProcAttr
// 	Process *os.Process
// 	ProcessState *os.ProcessState
// 	Cancel func() error
// 	WaitDelay time.Duration
// }
