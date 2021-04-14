package gowsl

import (
	"context"
	"os/exec"
)

var (
	binary    = "wsl"
	available = false
)

func init() {
	path, err := exec.LookPath(binary)
	if err == nil {
		binary = path
		available = true
	}
}

// Available WSL是否可用
func Available() bool {
	return available
}

// Command 返回*exec.Cmd以执行指定的程序
func Command(wsl string, name string, args ...string) *exec.Cmd {
	if wsl == "" {
		args = append([]string{name}, args...)
	} else {
		args = append([]string{"-d", wsl, name}, args...)
	}
	return exec.Command(binary, args...)
}

// CommandContext 与Command类似，但包含一个上下文
func CommandContext(ctx context.Context, wsl string, name string, args ...string) *exec.Cmd {
	if wsl == "" {
		args = append([]string{name}, args...)
	} else {
		args = append([]string{"-d", wsl, name}, args...)
	}
	return exec.CommandContext(ctx, binary, args...)
}
