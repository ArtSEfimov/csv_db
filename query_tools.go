package main

import "os"

func fileChecker(tableName string) (bool, string) {
	_, fileExistsErr := os.Stat(tableName)
	if os.IsNotExist(fileExistsErr) {
		return false, "file does not exist"
	}
	if fileExistsErr != nil {
		return false, "error checking if file exists"
	}
	return true, ""
}
