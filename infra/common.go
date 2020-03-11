package infra

import (
	"bytes"
	"context"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// 拡張子を除いたファイル名を返す
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

type Command struct {
	Name  string
	Args  []string
	Stdin string
}

type CommandResult struct {
	Stdout   string
	Stderr   string
	Error    error
	ExitCode int
}

// 2秒のタイムアウトを設定してコマンドを実行する
func execute(command *Command) *CommandResult {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, command.Name, command.Args...)
	cmd.Stdin = strings.NewReader(command.Stdin)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()

	result := &CommandResult{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: cmd.ProcessState.ExitCode(),
	}

	if err := ctx.Err(); err == context.DeadlineExceeded {
		result.Error = err
	}

	if err != nil {
		result.Error = err
	}

	return result
}
