package process

import (
	"testing"

	"golang.org/x/sys/unix"
)

func TestParseSignal(t *testing.T) {
	tests := []struct {
		name    string
		signal  string
		want    unix.Signal
		wantErr bool
	}{
		{"SIGTERM", "TERM", unix.SIGTERM, false},
		{"SIGKILL", "KILL", unix.SIGKILL, false},
		{"SIGHUP", "HUP", unix.SIGHUP, false},
		{"SIGINT", "INT", unix.SIGINT, false},
		{"Invalid", "INVALID", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSignal(tt.signal)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSignal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseSignal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTargets(t *testing.T) {
	// Basic test to ensure function runs without error
	targets, err := FindTargets("", "")
	if err != nil {
		t.Errorf("FindTargets() error = %v", err)
	}
	if len(targets) == 0 {
		t.Log("No targets found, but function completed successfully")
	}
}
