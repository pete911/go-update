package main

import (
	"fmt"
	"os/exec"
)

type Command struct {
	Name string
	Args []string
}

func RunCommands(dir string, commands []Command) error {
	for _, command := range commands {
		if err := RunCommand(dir, command); err != nil {
			return fmt.Errorf("%s %v: %w", command.Name, command.Args, err)
		}
	}
	return nil
}

func RunCommand(dir string, command Command) error {
	cmd := exec.Command(command.Name, command.Args...)
	cmd.Dir = dir
	Logf("running %s %v in %s", command.Name, command.Args, dir)
	out, err := cmd.Output()
	if len(out) != 0 {
		Logf(string(out))
	}
	return err
}

func HasGitChanges(dir string) (bool, error) {
	cmd := exec.Command("git", "status", "-s")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}
	return len(out) != 0, nil
}
