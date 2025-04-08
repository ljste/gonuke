package process

import (
    "fmt"
    "os"
    "golang.org/x/sys/unix"
    "runtime"
)

func KillTargets(targets []Target, verbose bool) error {
    for _, t := range targets {
        if verbose {
            fmt.Printf("Killing PID %d (%s)\n", t.PID, t.Name)
        }
        if runtime.GOOS == "windows" {
            // TODO: Windows termination
        } else {
            err := unix.Kill(t.PID, unix.SIGKILL)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Failed to kill PID %d: %v\n", t.PID, err)
            }
        }
    }
    return nil
}
