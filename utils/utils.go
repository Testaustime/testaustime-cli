package utils

import (
	"fmt"
	"os"

	"github.com/romeq/testaustime-cli/logger"
)

func Check(err error) {
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func StringOr(str1, str2 string) string {
	if str1 != "" {
		return str1
    } 
    return str2
}

func ColoredPrint(color int, a ...any) {
	if logger.ColorsEnabled {
		fmt.Printf("\033[%dm", color)
		fmt.Print(a...)
		fmt.Printf("\033[0m")
		return
	}
	fmt.Print(a...)
}

func NthElement(list []string, n int) string {
	if n >= 0 && len(list) > n {
		return list[n]
	}
	return ""
}

func EnvOrString(envKey, alt string) string {
    return StringOr(os.Getenv(envKey), alt)
}

