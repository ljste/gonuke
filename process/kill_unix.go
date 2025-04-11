//go:build !windows

package process

import (
    "golang.org/x/sys/unix"
)

func killProcess(pid int, signal unix.Signal) error {
    return unix.Kill(pid, signal)
}