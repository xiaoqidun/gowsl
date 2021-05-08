package gowsl

import (
	"syscall"
	"unsafe"
)

var dll, _ = syscall.LoadDLL(`Api-ms-win-wsl-api-l1-1-0.dll`)

// Registered 发行版是否注册
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
