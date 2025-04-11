package process

import (
	"fmt"
	"strings"

	"golang.org/x/sys/unix"
)

// ParseSignal converts a signal name (e.g., "TERM", "KILL") to a unix.Signal.
func ParseSignal(sigName string) (unix.Signal, error) {
	switch strings.ToUpper(sigName) {
	case "TERM":
		return unix.SIGTERM, nil
	case "KILL":
		return unix.SIGKILL, nil
	case "HUP":
		return unix.SIGHUP, nil
	case "INT":
		return unix.SIGINT, nil
	default:
		return 0, fmt.Errorf("unsupported signal: %s", sigName)
	}
}
