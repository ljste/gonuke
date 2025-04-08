# gonuke

**Cross-platform CLI tool to find and force-kill unwanted processes**

---

## Features
- Find running processes by **name pattern** (regex)
- Dry run mode: *see* what would be nuked before killing
- Confirmations to avoid accidents (override with `--force`)
- Mass kill matched processes with SIGKILL (`-9`) on Unix
- Designed for dev, CI cleanup, or stubborn zombie removal

---

## Usage

```bash
# See all matching processes named 'foo' (dry-run, no kill)
gonuke --pattern foo --dry-run

# Force kill all 'myapp' processes, no confirmation
gonuke --pattern myapp --force

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

- Filter processes owned by a specific user
- Recursive killing of children and process trees
- Timeout & graceful SIGTERM before nuke
- Windows support

---

## License

MIT

---

# **Use with care!**  
gonuke kills what you tell it — don't run as root unless you mean it.
