//go:build unix

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
	syscall.SIGUSR1,
	syscall.SIGUSR2,
	syscall.SIGALRM,
	syscall.SIGTERM,
}

func osGetExitStatus(state *os.ProcessState) int {
	return state.Sys().(syscall.WaitStatus).ExitStatus()
}
