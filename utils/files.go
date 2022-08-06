package utils

import (
	"errors"
	"os"
)

func WriteFile(file, content string) {
	err := os.WriteFile(file, []byte(content), 0644)
	Check(err)
}

func ReadFile(file string, allowNonexistent bool) string {
	bytes, err := os.ReadFile(file)
	if allowNonexistent && errors.Is(err, os.ErrNotExist) {
		return ""
	}
	Check(err)
	return string(bytes)
}

func ReadToken(filename string) string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	tokenBytes := make([]byte, 32)
	_, err = file.Read(tokenBytes)
	Check(err)

	return string(tokenBytes)
}
