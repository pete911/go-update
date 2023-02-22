package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const defaultCommitMessage = "update dependencies"

func Update(dir, version string) error {
	if err := RunCommand(dir, Command{"git", []string{"pull"}}); err != nil {
		return fmt.Errorf("  git pull: %w", err)
	}
	if err := updateGoMod(dir, version); err != nil {
		return fmt.Errorf("  update go.mod: %w", err)
	}
	if err := updateGoDependencies(dir); err != nil {
		return fmt.Errorf("  update go dependencies: %w", err)
	}
	return gitCommitAndPush(dir)
}

func gitCommitAndPush(dir string) error {
	changes, err := HasGitChanges(dir)
	if err != nil {
		return fmt.Errorf("  checking changes: %w", err)
	}
	if changes {
		ok, err := PromptYN(fmt.Sprintf("%s project has changes, push to git", filepath.Base(dir)))
		if err != nil {
			return err
		}
		if ok {
			commitMessage, err := PromptDefault("commit message", defaultCommitMessage)
			if err != nil {
				return err
			}
			return RunCommands(dir, []Command{
				{"git", []string{"add", "."}},
				{"git", []string{"commit", "-m", commitMessage}},
				{"git", []string{"push"}},
			})
		}
	}

	Logf("  no changes in %s", dir)
	return nil
}

func updateGoDependencies(dir string) error {
	return RunCommands(dir, []Command{
		{"go", []string{"get", "-u", "-t", "./..."}},
		{"go", []string{"mod", "tidy"}},
		{"go", []string{"mod", "vendor"}},
	})
}

func updateGoMod(dir, version string) error {
	goModVersion, err := getGoModVersion(dir)
	if err != nil {
		return err
	}
	if version != goModVersion {
		ok, err := PromptYN(fmt.Sprintf("go.mod go %s is different from current %s version, update go.mod", goModVersion, version))
		if err != nil {
			return err
		}
		if ok {
			return updateGoModVersion(dir, version)
		}
	}
	return nil
}

func getGoModVersion(dir string) (string, error) {
	goModPath := getGoModPath(dir)
	lines, err := readFile(goModPath)
	if err != nil {
		return "", err
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "go ") {
			return strings.Split(line, " ")[1], nil
		}
	}
	return "", fmt.Errorf("  no version found in %s", goModPath)
}

func updateGoModVersion(dir, version string) error {
	goModPath := getGoModPath(dir)
	lines, err := readFile(goModPath)
	if err != nil {
		return err
	}

	for i, line := range lines {
		if strings.HasPrefix(line, "go ") {
			lines[i] = fmt.Sprintf("go %s", version)
			break
		}
	}
	return os.WriteFile(goModPath, []byte(strings.Join(lines, "\n")), 0644)
}

func readFile(goModPath string) ([]string, error) {
	file, err := os.Open(goModPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func getGoModPath(dir string) string {
	return filepath.Join(dir, "go.mod")
}

func GoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("go version: %w", err)
	}
	version := strings.TrimPrefix(string(out), "go version go")
	versionParts := strings.Split(strings.Split(version, " ")[0], ".")
	if len(versionParts) < 2 {
		return "", fmt.Errorf("unexpected version part %v of %s", versionParts, version)
	}
	return fmt.Sprintf("%s.%s", versionParts[0], versionParts[1]), nil
}
