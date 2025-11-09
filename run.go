//$(cd "$(dirname "$0")"; pwd)/bootstrap.sh go run "$0" "$@"; exit
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func version() error {
	branch := getEnv("CIRCLE_BRANCH", gitBranch())
	buildNumber := getEnv("CIRCLE_BUILD_NUM", time.Now().UTC().Format("20060102150405"))
	revisions := gitRevCount(branch)
	version := fmt.Sprintf("0.%s.%s", revisions, buildNumber)

	fmt.Printf("version: %s\n", version)
	os.Setenv("VERSION", version)
	return nil
}

func tag() error {
	if err := version(); err != nil {
		return err
	}

	version := os.Getenv("VERSION")

	cmds := [][]string{
		{"git", "config", "--global", "user.name", "Server"},
		{"git", "config", "--global", "user.email", "server@bodar.com"},
		{"git", "tag", "-a", version, "-m", fmt.Sprintf("Release %s", version)},
		{"git", "push", "origin", version},
	}

	for _, args := range cmds {
		if err := runCmd(args...); err != nil {
			return err
		}
	}
	return nil
}

func clean() error {
	return runCmd("rm", "-rf", "artifacts")
}

func dev() error {
	return runCmd("hugo", "server")
}

func build() error {
	return runCmd("hugo")
}

func ci() error {
	return build()
}

// Helper functions

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func gitBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "master"
	}
	return strings.TrimSpace(string(out))
}

func gitRevCount(branch string) string {
	cmd := exec.Command("git", "rev-list", "--count", branch)
	out, err := cmd.Output()
	if err != nil {
		return "0"
	}
	return strings.TrimSpace(string(out))
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
		"version": version,
		"tag":     tag,
		"clean":   clean,
		"dev":     dev,
		"build":   build,
		"ci":      ci,
	}

	fn, exists := commands[command]
	if !exists {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		fmt.Fprintf(os.Stderr, "Available commands: version, tag, clean, dev, build, ci\n")
		os.Exit(1)
	}

	if err := fn(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
