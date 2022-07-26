package datahelper

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/romeq/testaustime-cli/logger"
	"github.com/romeq/testaustime-cli/utils"
	"golang.org/x/term"
)

func AskInput(prompt string) string {
	utils.ColoredPrint(35, prompt)
	fmt.Print(": ")
	result, err := bufio.NewReader(os.Stdin).ReadString('\n')
	utils.Check(err)

	return strings.TrimSpace(string(result))
}

func AskPassword(prompt string) string {
	utils.ColoredPrint(35, utils.StringOr(prompt, "Password"))
	fmt.Print(": ")

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	utils.Check(err)
	fmt.Print("\n")

	password := strings.TrimSpace(string(bytePassword))
	if password == "" {
		logger.Error(errors.New("Password is required"))
	} else if len(password) < 8 || len(password) > 128 {
		logger.Error(errors.New(
			"Password has to be between 8 and 128 characters long",
		))
	}

	return password
}
