package process

import (
    "fmt"
    "os"
    "os/user"
    "regexp"
    "runtime"
    "strconv"
    "syscall"

    ps "github.com/mitchellh/go-ps"
)

type Target struct {
    PID  int
    Name string
    User string
}

func getUserID(username string) (int, error) {
    u, err := user.Lookup(username)
    if err != nil {
        return -1, err
    }
    uid, err := strconv.Atoi(u.Uid)
    if err != nil {
        return -1, err
    }
    return uid, nil
}

func getProcessUser(pid int) (string, error) {
    if runtime.GOOS == "windows" {
        return "", fmt.Errorf("user filtering not implemented for Windows")
    }

    procPath := fmt.Sprintf("/proc/%d", pid)
    info, err := os.Stat(procPath)
    if err != nil {
        return "", err
    }

    stat, ok := info.Sys().(*syscall.Stat_t)
    if !ok {
        return "", fmt.Errorf("failed to get process stats")
    }

    u, err := user.LookupId(strconv.Itoa(int(stat.Uid)))
    if err != nil {
        return "", err
    }

    return u.Username, nil
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
        // Skip if name doesn't match pattern
        if re != nil && !re.MatchString(p.Executable()) {
            continue
        }

        // Get process username
        username := ""
        if runtime.GOOS != "windows" {
            username, err = getProcessUser(p.Pid())
            if err != nil {
                // Skip if we can't get user info
                continue
            }

            // Filter by user if specified
            if userFilter != "" && username != userFilter {
                continue
            }
        }

        targets = append(targets, Target{
            PID:  p.Pid(),
            Name: p.Executable(),
            User: username,
        })
    }

    return targets, nil
}