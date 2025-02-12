package job

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestJob(t *testing.T) {
	// Jobインスタンスを作成
	job := NewJob()
	job.Stdin = os.Stdin

	cmd1 := exec.Command("echo", "Hello")
	cmd2 := exec.Command("cat")
	job.cmds = append(job.cmds, cmd1, cmd2)

	job.Start()

	job.Wait()
}

func TestJobStart(t *testing.T) {
	// Jobインスタンスを作成
	job := NewJob()

	// テスト用コマンドを設定
	cmd1 := exec.Command("echo", "Hello")
	job.cmds = append(job.cmds, cmd1)

	// コマンドの標準出力を捕捉
	var out bytes.Buffer
	job.Stdout = &out

	// コマンドを開始
	gpid, err := job.Start()
	if err != nil {
		t.Fatalf("Job.Start() failed: %v", err)
	}
	job.Wait()
	// プロセスグループIDが設定されているか確認
	if gpid < 0 {
		t.Errorf("Expected positive gpid, got %d", gpid)
	}

	// コマンドが実行されたか確認
	if out.String() != "Hello\n" {
		t.Errorf("Expected output 'Hello', got %s", out.String())
	}
}

func TestJobStart_Error(t *testing.T) {
	// Jobインスタンスを作成
	job := NewJob()

	// 不正なコマンドを設定 (存在しないコマンド)
	cmd1 := exec.Command("nonexistentcommand")
	job.cmds = append(job.cmds, cmd1)

	// コマンドの標準出力を捕捉
	var out []byte
	job.Stdout = os.Stdout

	// コマンドを開始
	_, err := job.Start()
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}
	job.Wait()
	// 標準出力が空であることを確認
	if len(out) > 0 {
		t.Errorf("Expected no output, but got %s", out)
	}
}

func TestJobStart_MultipleCommands(t *testing.T) {
	// Jobインスタンスを作成
	job := NewJob()

	// 複数コマンドを設定 (例: echo "Hello" | awk '{print $1}')
	cmd1 := exec.Command("echo", "Hello")
	cmd2 := exec.Command("awk", "{print $1}")

	// コマンドをJobに追加
	job.cmds = append(job.cmds, cmd1, cmd2)

	// 出力をキャプチャするためのバッファ
	var buf bytes.Buffer
	job.Stdout = &buf

	// コマンドを開始
	_, err := job.Start()
	if err != nil {
		t.Fatalf("Job.Start() failed: %v", err)
	}
	job.Wait()

	// 結果の確認
	expected := "Hello\n"
	actual := buf.String()
	if actual != expected {
		t.Errorf("Expected output '%s', but got '%s'", expected, actual)
	}
}
