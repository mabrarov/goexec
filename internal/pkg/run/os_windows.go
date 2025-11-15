//go:build windows

package run

import (
	"os"
	"syscall"
)

var osNotifySignals = []os.Signal{
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGQUIT,
	syscall.SIGABRT,
	syscall.SIGALRM,
	syscall.SIGTERM,
}

func osGetExitStatus(state *os.ProcessState) int {
	return state.ExitCode()
}
