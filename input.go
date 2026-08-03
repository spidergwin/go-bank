package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString(prompt string) (string, error) {
	fmt.Print(prompt + ": ")

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	text = strings.TrimSpace(text)
	return text, nil
}

func ReadConfirmation(prompt string) (bool, error) {
	fmt.Print(prompt + " (Y/n): ")

	confirmation, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	confirmation = strings.TrimSpace(strings.ToLower(confirmation))

	switch confirmation {
	case "y", "yes", "":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid choice")
	}
}

func ReadInt(prompt string) (int, error) {
	fmt.Print(prompt + ": ")

	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	amount, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return 0, err
	}

	return amount, nil
}
