//$(cd "$(dirname "$0")"; pwd)/bootstrap.sh "$0" "$@"; exit
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clean() error {
	return runCmd("rm", "-rf", "public")
}

func start() error {
	clean()
	return runCmd("hugo", "server")
}

func build() error {
	clean()
	return runCmd("hugo")
}

func runCmd(args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func main() {
	command := "build"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	commands := map[string]func() error{
		"clean":   clean,
		"start":   start,
		"build":   build,
	}

	fn, exists := commands[command]
	if !exists {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		fmt.Fprintf(os.Stderr, "Available commands: clean, start, build (default) \n")
		os.Exit(1)
	}

	if err := fn(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
