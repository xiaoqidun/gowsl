package gowsl

import (
	"context"
	"os/exec"
	"syscall"
	"unsafe"
)

var (
	binary    = "wsl"
	available = false
	dll, _    = syscall.LoadDLL(`Api-ms-win-wsl-api-l1-1-0.dll`)
)

func init() {
	path, err := exec.LookPath(binary)
	if err == nil {
		binary = path
		available = true
	}
}

// Available 检查WSL是否可用
func Available() bool {
	return available
}

// Registered 检查发行版是否注册
func Registered(distribution string) bool {
	if dll != nil {
		proc, err := dll.FindProc("WslIsDistributionRegistered")
		if err == nil {
			ptr, _ := syscall.UTF16PtrFromString(distribution)
			ret, _, _ := proc.Call(uintptr(unsafe.Pointer(ptr)))
			return ret == 1
		}
	}
	return false
}

// Command 返回*exec.Cmd以执行指定的程序
func Command(distribution string, user string, name string, args ...string) *exec.Cmd {
	args = append([]string{name}, args...)
	if user != "" {
		args = append([]string{"-u", user}, args...)
	}
	if distribution != "" {
		args = append([]string{"-d", distribution}, args...)
	}
	return exec.Command(binary, args...)
}

// CommandContext 与Command类似，但包含一个上下文
func CommandContext(ctx context.Context, distribution string, user string, name string, args ...string) *exec.Cmd {
	args = append([]string{name}, args...)
	if user != "" {
		args = append([]string{"-u", user}, args...)
	}
	if distribution != "" {
		args = append([]string{"-d", distribution}, args...)
	}
	return exec.CommandContext(ctx, binary, args...)
}
