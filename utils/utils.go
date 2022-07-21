package utils

import (
	"os"

	"github.com/romeq/testaustime-cli/logger"
)

func Check(err error) {
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func StringOr(str1 string, str2 string) string {
	if str1 != "" {
		return str1
	} else {
		return str2
	}
}
