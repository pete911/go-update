package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptYN(msg string) (bool, error) {
	input, err := Prompt(fmt.Sprintf("%s (y/N)", msg))
	if err != nil {
		return false, err
	}
	return strings.ToLower(input) == "y", nil
}

func Prompt(msg string) (string, error) {
	fmt.Printf("%s: ", msg)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(input, "\n"), nil
}

func PromptDefault(msg, defaultInput string) (string, error) {
	fmt.Printf("%s (%s): ", msg, defaultInput)
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	if in := strings.TrimSuffix(input, "\n"); in != "" {
		return in, nil
	}
	return defaultInput, nil
}
