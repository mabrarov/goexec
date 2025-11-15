//go:build !linux

package run

import (
	"os/exec"
)

func osSetSysProcAttr(_ *exec.Cmd) {}
