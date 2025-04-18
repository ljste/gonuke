package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ljste/gonuke/process"
)

func main() {
	pattern := flag.String("pattern", "", "Regex pattern to match process names")
	user := flag.String("user", "", "Target processes owned by this user")
	signal := flag.String("signal", "KILL", "Signal to send (TERM, KILL, HUP, INT)")
	dryRun := flag.Bool("dry-run", false, "Show what would be killed without killing")
	force := flag.Bool("force", false, "Skip confirmation prompt")
	verbose := flag.Bool("verbose", false, "Verbose output")
	flag.Parse()

	if *pattern == "" && *user == "" {
		fmt.Println("Please specify --pattern and/or --user to select targets")
		flag.Usage()
		os.Exit(1)
	}

	sig, err := process.ParseSignal(*signal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid signal: %v\n", err)
		os.Exit(1)
	}

	targets, err := process.FindTargets(*pattern, *user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding processes: %v\n", err)
		os.Exit(1)
	}

	if len(targets) == 0 {
		fmt.Println("No processes matched")
		return
	}

	fmt.Printf("Found %d target(s):\n", len(targets))
	for _, proc := range targets {
		fmt.Printf("- PID %d : %s (user: %s)\n", proc.PID, proc.Name, proc.User)
	}

	if *dryRun {
		fmt.Println("Dry run mode. Exiting without killing.")
		return
	}

	if !*force {
		var answer string
		fmt.Printf("Proceed to send signal %s to these processes? (y/N): ", *signal)
		fmt.Scanln(&answer)
		if answer != "y" && answer != "Y" {
			fmt.Println("Aborted.")
			return
		}
	}

	err = process.KillTargets(targets, sig, *verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error killing processes: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Done.\n")
}