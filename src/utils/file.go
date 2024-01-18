package utils

import "os"

// fileExists checks if a file exists at the specified path
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
