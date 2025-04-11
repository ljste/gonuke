# gonuke

**Cross-platform CLI tool to find and force-kill unwanted processes**

---

## Features
- Find running processes by **name pattern** (regex)
- Filter processes owned by a specific user
- Multiple signal support (TERM, KILL, HUP, INT)
- Dry run mode: *see* what would be nuked before killing
- Confirmations to avoid accidents (override with `--force`)
- Cross-platform support (Unix/Windows)
- Designed for dev, CI cleanup, or stubborn zombie removal

---

## Usage

```bash
# See all matching processes named 'foo' (dry-run, no kill)
gonuke --pattern foo --dry-run

# Force kill all 'myapp' processes, no confirmation
gonuke --pattern myapp --force

# Gracefully terminate matching processes with SIGTERM
gonuke --pattern myapp --signal TERM

# Kill processes owned by specific user
gonuke --user someuser --pattern myapp

# See CLI flags
gonuke --help
```

---

## Install / Build

```bash
git clone https://github.com/yourusername/gonuke.git
cd gonuke
go build -o gonuke ./cmd/gonuke
```

---

## TODO / planned

- Recursive killing of children and process trees
- Timeout & graceful SIGTERM before nuke
- Better Windows user filtering support