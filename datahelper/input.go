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

// AskInput prompts user for input in stdin.
// result is trimmed.
func AskInput(prompt string) string {
	utils.ColoredPrint(35, prompt)
	fmt.Print(": ")

	result, err := bufio.NewReader(os.Stdin).ReadString('\n')
	utils.Check(err)
	result = strings.TrimSpace(string(result))
	if result == "" {
		logger.Error(errors.New("You fool, input can't be an empty string!"))
	}

	return result
}

// AskPassword prompts user for hidden input in stdin, most
// used for passwords. Returns a pointer to spacetrimmed result.
func AskPassword(prompt string) *string {
	utils.ColoredPrint(35, utils.StringOr(prompt, "Password"))
	fmt.Print(": ")

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	utils.Check(err)
	fmt.Print("\n")

	password := strings.TrimSpace(string(bytePassword))
	if password == "" {
		logger.Error(errors.New("You fool, password can't be an empty string!"))
	} else if len(password) < 8 || len(password) > 128 {
		logger.Error(errors.New(
			"Password has to be between 8 and 128 characters long",
		))
	}

	return &password
}
