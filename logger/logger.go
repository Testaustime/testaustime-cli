package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var errLogger *log.Logger = log.New(os.Stderr, "", log.LUTC)
var ColorsEnabled bool = true

func Error(err error) {
	logMessage(errLogger, fmt.Sprint(
		coloredType("AARGH, Error!", 31),
		strings.ToLower(err.Error()),
	))
	os.Exit(1)
}

func logMessage[T any](logger *log.Logger, message T) {
	logger.Println(message)
}

func coloredType(str string, color int) string {
	if ColorsEnabled {
		return fmt.Sprintf("\033[%dm%s\033[0m ", color, str)
	}
	return str + " "
}
