package utils

import (
	"os"
	"path"
)

// resolveConfigPath will return expanded dir joined to project name and file argument
func ResolveWantedPath(dir, file string) string {
	globalConfigDir := os.ExpandEnv(dir)
	return path.Join(globalConfigDir, "/testaustime-cli/", file)
}
