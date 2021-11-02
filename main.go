package main

import (
	"io"
	"log"
	"os/exec"
	"strings"
)

func main() {
	var gitPresent bool
	var gitVersion string

	gitTagCmd := exec.Command("git", "describe --tags")
	stdErr, err := gitTagCmd.StderrPipe()
	if err != nil {
		log.Fatalf("failed connecting to Stderr pipe for git command: %v", err)
	}
	stdOut, err := gitTagCmd.StdoutPipe()
	if err != nil {
		log.Fatalf("failed connecting to Stdout pipe for git command: %v", err)
	}
	err = gitTagCmd.Start()
	if err != nil {
		log.Fatalf("failed starting git command: %v", err)
	}

	gitErr, _ := io.ReadAll(stdErr)
	gitVersionBytes, _ := io.ReadAll(stdOut)

	if err := gitTagCmd.Wait(); err != nil {
		log.Fatalf("failed waiting for git command: %v", err)
	}

	if len(gitErr) > 0 {
		log.Printf("received error from git command: %s\ngit version tag not present\n", string(gitErr))
	} else {
		gitVersion = strings.TrimSpace(string(gitVersionBytes))
		gitPresent = true
		log.Printf("git version tag: %s", gitVersion)
	}

	buildCmd := exec.Command("go", "build")
}
