package builtins

import (
	"fmt"
	"os"
	"shell/job"
	"strconv"
)

func jobs(args []string) int {
	// ジョブが1つもない場合
	if len(job.Jobs) == 0 {
		fmt.Println("No jobs are currently running.")
		return 0
	}

	// ジョブが存在する場合にリスト表示
	for i, j := range job.Jobs {
		status := "isRunning"
		if !j.IsRunning {
			status = "isStopped"
		}
		fmt.Printf("[%d] %s %s\n", i, status, j.CmdName)
	}

	return 0
}

func fg(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "fg: usage: fg <job-id>")
		return 1
	}

	jobID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fg: invalid job id: %s\n", args[0])
		return 1
	}

	if err := job.Foreground(jobID); err != nil {
		fmt.Fprintf(os.Stderr, "fg: %v\n", err)
		return 1
	}

	return 0
}

func bg(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "bg: usage: bg <job-id>")
		return 1
	}

	jobID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "bg: invalid job id: %s\n", args[0])
		return 1
	}

	if err := job.Background(jobID); err != nil {
		fmt.Fprintf(os.Stderr, "bg: %v\n", err)
		return 1
	}

	return 0
}
