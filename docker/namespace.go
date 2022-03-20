package namespace

import (
	"os"
	"os/exec"
	"syscall"
)

func cloneNmespace() (cmd exec.Cmd) {
	cmd = exec.Cmd{Path: "/bin/sh"}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
