package main

import (
	"fmt"
	"os/exec"
)

// git tag | tail -n 1
func main() {
	lastTagCommand := "git tag | tail -n 1"
	tag := runCommand(lastTagCommand)
	lastCommitMessageCommand := runCommand("git log -1 --pretty=%B")
	commitMessage := runCommand(lastCommitMessageCommand)
	fmt.Print(tag, commitMessage)
}

// this func running command to terminal
func runCommand(command string) string {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return string(output)
}
