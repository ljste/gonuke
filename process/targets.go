package process

import (
    "os/user"
    "regexp"
    "runtime"

    ps "github.com/mitchellh/go-ps"
)

type Target struct {
    PID  int
    Name string
}

func FindTargets(pattern string, userFilter string) ([]Target, error) {
    var targets []Target

    procs, err := ps.Processes()
    if err != nil {
        return nil, err
    }

    var re *regexp.Regexp
    if pattern != "" {
        re, err = regexp.Compile(pattern)
        if err != nil {
            return nil, err
        }
    }

    for _, p := range procs {
        // name match
        if re != nil && !re.MatchString(p.Executable()) {
            continue
        }

        // (optionally) user match
        if userFilter != "" {
            if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
                // platform-specific: TODO! skip for now (can implement via /proc or ps output parsing)
            } else {
                continue
            }
        }

        targets = append(targets, Target{
            PID:  p.Pid(),
            Name: p.Executable(),
        })
    }
    return targets, nil
}
