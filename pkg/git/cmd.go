package git

import (
	"fmt"
	"os"
	"os/exec"
)

// makeConfigCmd creates a list of commands to configure git.
// commands:
// git config --global user.name ${{ github.actor }}
// git config --global user.email ${{ github.actor }}@users.noreply.github.com
// nolint: gosec
func makeConfigCmd(user string) []*exec.Cmd {
	list := []*exec.Cmd{
		exec.Command("git", "config", "--global", "user.name", user),
		exec.Command("git", "config", "--global", "user.email", user+"@users.noreply.github.com"),
	}
	return list
}

// makeSaveCmd creates a list of commands to save a file to a branch.
// commands:
// git add ${{ inputs.data_file }}
// git commit -am "save ${{ inputs.data_file }}"
// git push origin main
// nolint: gosec
func makeSaveCmd(branch, path string) []*exec.Cmd {
	list := []*exec.Cmd{
		exec.Command("git", "add", path),
		exec.Command("git", "commit", "-am", fmt.Sprintf("save %s", path)),
		exec.Command("git", "push", "origin", branch),
	}
	return list
}

// runCmd runs a command and returns an error if any.
// may need to write the output to a file at some point.
func runCmd(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
