// +build !windows

package gowsl

// Registered 发行版是否注册
func Registered(distribution string) bool {
	return false
}
