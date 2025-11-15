//go:build linux

package run

import (
	"os/exec"
	"syscall"
)

func osSetSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGKILL}
}
