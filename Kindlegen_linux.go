// +build linux

package main

import (
	"os/exec"
)

func KindlegenCmd(args ...string) *exec.Cmd {
	cmd := exec.Command("./tools/kindlegenLinux", args...)
	return cmd
}
