package job

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
	//"github.com/codegangsta/inject"
)

type Job struct {
	cmds    []*exec.Cmd
	Dir     string
	Env     map[string]string
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	ShowCMD bool
	timeout time.Duration

	CmdName   string
	IsRunning bool
	// 	ExtraFiles []*os.File
	gpid     int
	PipeFail bool
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

var Jobs []Job

func NewJob() *Job {
	j := &Job{
		cmds:    make([]*exec.Cmd, 0, 4),
		Dir:     ".",
		Env:     make(map[string]string),
		ShowCMD: true,
		//Stdin: ,端末に繋がってる時のみ
		Stdin:  strings.NewReader(""),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	return j
}

func (j *Job) writePrompt(prompt string) {
	// プロンプトを表示
	fmt.Println(prompt)
}

func (s *Job) Start() (pgid int, err error) {
	s.IsRunning = true
	var rd *io.PipeReader
	var wr *io.PipeWriter
	var length = len(s.cmds)
	// if s.ShowCMD {
	// 	var cmds = make([]string, 0, 4)
	// 	for _, cmd := range s.cmds {
	// 		cmds = append(cmds, strings.Join(cmd.Args, " "))
	// 	}
	// 	s.writePrompt(strings.Join(cmds, " | "))
	// }
	for index, cmd := range s.cmds {
		if index == 0 {
			cmd.Stdin = s.Stdin
		} else {
			cmd.Stdin = rd
		}
		if index != length-1 {
			rd, wr = io.Pipe() // create pipe
			cmd.Stdout = wr
			cmd.Stderr = os.Stderr
		}
		if index == length-1 {
			cmd.Stdout = s.Stdout
			cmd.Stderr = s.Stderr
		}
		err = cmd.Start()
		if err != nil {
			return
		}
		if index == 0 {
			pgid = cmd.Process.Pid
		}
	}

	return
}

func (s *Job) Wait() error {
	var pipeErr, lastErr error
	for _, cmd := range s.cmds {
		if lastErr = cmd.Wait(); lastErr != nil {
			pipeErr = lastErr
		}
		wr, ok := cmd.Stdout.(*io.PipeWriter)
		if ok {
			wr.Close()
		}
	}
	if s.PipeFail {
		return pipeErr
	}
	return lastErr
}

// func (j *Job) Start() (gpid int, err error) {
// 	var rd *io.PipeReader
// 	var wr *io.PipeWriter
// 	var length = len(j.cmds)
// 	var pgid int

// 	// コマンドを表示するオプションが有効な場合
// 	if j.ShowCMD {
// 		var cmds = make([]string, 0, 4)
// 		for _, cmd := range j.cmds {
// 			cmds = append(cmds, strings.Join(cmd.Args, " "))
// 		}
// 		j.writePrompt(strings.Join(cmds, " | "))
// 	}

// 	// 最初のコマンドで新しいセッションを開始 (setsid() を呼び出し)
// 	// _, _, errno := syscall.Syscall(syscall.SYS_SETSID, 0, 0, 0)
// 	// if errno != 0 {
// 	// 	err = fmt.Errorf("setsid() failed: %v", errno)
// 	// 	return
// 	// }

// 	// 最初のコマンドのプロセスグループIDを設定
// 	for index, cmd := range j.cmds {
// 		if index == 0 {
// 			cmd.Stdin = j.Stdin
// 		} else {
// 			cmd.Stdin = rd
// 		}
// 		if index != length-1 {
// 			rd, wr = io.Pipe()
// 			cmd.Stdout = wr
// 			cmd.Stderr = os.Stderr

// 		}
// 		if index == length-1 {
// 			cmd.Stdout = j.Stdout
// 			cmd.Stderr = j.Stderr
// 		}

// 		// コマンドを開始
// 		err = cmd.Start()
// 		if err != nil {
// 			return
// 		}

// 		// 最初のコマンドのプロセスグループIDを設定
// 		if index == 0 {
// 			pgid = cmd.Process.Pid
// 		} // 最初のコマンドのPIDを取得
// 		// 	err := syscall.Setpgid(cmd.Process.Pid, pgid) // 新しいプロセスグループを設定
// 		// 	if err != nil {
// 		// 		return gpid, fmt.Errorf("failed to setpgid for pid %d: %v", cmd.Process.Pid, err)
// 		// 	}
// 		// 	gpid = pgid // プロセスグループIDを設定
// 		// } else {
// 		// 	// その他のコマンドも最初のプロセスグループに設定
// 		// 	err := syscall.Setpgid(cmd.Process.Pid, gpid) // 最初のコマンドのPGIDを設定
// 		// 	if err != nil {
// 		// 		return gpid, fmt.Errorf("failed to setpgid for pid %d: %v", cmd.Process.Pid, err)
// 		// 	}
// 	}

// 	// 全てのコマンドが開始され、gpid が設定される
// 	j.gpid = pgid
// 	return pgid, nil
// }

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
