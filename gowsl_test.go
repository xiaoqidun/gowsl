package gowsl

import "testing"

func TestRegistered(t *testing.T) {
	t.Log(Registered(`Debian`))
	t.Log(Registered(`Ubuntu`))
}
