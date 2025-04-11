//go:build windows

package process

import (
    "golang.org/x/sys/unix"
    "golang.org/x/sys/windows"
)

func killProcess(pid int, _ unix.Signal) error {
    handle, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, uint32(pid))
    if err != nil {
        return err
    }
    defer windows.CloseHandle(handle)
    return windows.TerminateProcess(handle, 1)
}