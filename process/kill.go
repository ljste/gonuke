package process

import (
    "fmt"
    "golang.org/x/sys/unix"
)

type killError struct {
    pid int
    err error
}

func (e *killError) Error() string {
    return fmt.Sprintf("PID %d: %v", e.pid, e.err)
}

func KillTargets(targets []Target, signal unix.Signal, verbose bool) error {
    var errs []error

    for _, t := range targets {
        if verbose {
            fmt.Printf("Sending signal %v to PID %d (%s)\n", signal, t.PID, t.Name)
        }

        if err := killProcess(t.PID, signal); err != nil {
            errs = append(errs, &killError{pid: t.PID, err: err})
        }
    }

    if len(errs) > 0 {
        return fmt.Errorf("failed to kill some processes: %v", errs)
    }
    return nil
}