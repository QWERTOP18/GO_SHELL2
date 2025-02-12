package expand

import (
	"os"
	"path/filepath"
)

func expandTilde(t string) string {
	if t[0] != '~' {
		return t
	}
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return t
	}
	return filepath.Join(homeDir, t[1:])
}
