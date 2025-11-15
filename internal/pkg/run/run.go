package run

import (
	"context"
	"os"
	"os/exec"
	"os/signal"
)

func Cmd(name string, args ...string) (int, error) {
	cmd := exec.CommandContext(context.Background(), name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	osSetSysProcAttr(cmd)

	err := cmd.Start()
	if err != nil {
		return 0, err
	}

	sigCh := make(chan os.Signal, len(osNotifySignals)*2)
	signal.Notify(sigCh, osNotifySignals...)
	defer func() {
		signal.Stop(sigCh)
		close(sigCh)
	}()

	go func() {
		for sig := range sigCh {
			_ = cmd.Process.Signal(sig)
		}
	}()

	err = cmd.Wait()

	if cmd.ProcessState != nil {
		return osGetExitStatus(cmd.ProcessState), nil
	}
	return 0, err
}
