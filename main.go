package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	lastTag := runCommand("git tag | tail -n 1")
	lastCommitMessage := runCommand("git log -1 --pretty=%B")
	fmt.Println("Last tag:", lastTag)
	fmt.Println("Commit message:", lastCommitMessage)

	if lastTag == "" {
		panic("No tags found, you need to create a tag first")
	}

	lastTag = strings.ReplaceAll(lastTag, "v", "")
	lastTag = strings.ReplaceAll(lastTag, "\n", "")
	versions := strings.Split(lastTag, ".")

	if len(versions) != 3 {
		panic("Invalid tag format, you need to use SemVer")
	}

	major, err := strconv.Atoi(versions[0])
	if err != nil {
		panic("Invalid major version, you need to use SemVer " + err.Error())
	}

	minor, err := strconv.Atoi(versions[1])
	if err != nil {
		panic("Invalid minor version, you need to use SemVer " + err.Error())
	}

	patch, err := strconv.Atoi(versions[2])
	if err != nil {
		panic("Invalid patch version, you need to use SemVer " + err.Error())
	}

	if strings.Contains(strings.ToLower(lastCommitMessage), "[major]") ||
		strings.Contains(strings.ToLower(lastCommitMessage), "#major") {
		major++
		minor = 0
		patch = 0
	} else if strings.Contains(strings.ToLower(lastCommitMessage), "[minor]") ||
		strings.Contains(strings.ToLower(lastCommitMessage), "#minor") ||
		strings.Contains(strings.ToLower(lastCommitMessage), "feat") {
		minor++
		patch = 0
	} else {
		patch++
	}

	newTag := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	fmt.Println("New tag:", newTag)
	runCommand(fmt.Sprintf("git tag %s", newTag))
}

// runCommand executes a shell command and returns the output
func runCommand(command string) string {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return string(output)
}
