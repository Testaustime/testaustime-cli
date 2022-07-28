package logger

import (
	"fmt"
	"log"
	"os"
)

var errLogger *log.Logger = log.New(os.Stderr, "", log.LUTC)
var ColorsEnabled bool = true

func Error(err error) {
	logMessage(errLogger, fmt.Sprint(
		coloredType("Aw, an error occured :(", 31),
		err.Error(),
	))
	os.Exit(1)
}

func Info(message string) {
	logMessage(errLogger, fmt.Sprint(
		coloredType("info", 32),
		message,
	))
}

func logMessage(logger *log.Logger, message any) {
	logger.Println(message)
}

func coloredType(str string, color int) string {
	if ColorsEnabled {
		return fmt.Sprintf("\033[%dm%s\033[0m ", color, str)
	}
	return str + " "
}
