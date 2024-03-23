package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// generateRandomName generates a random string of the specified length.
func generateRandomName(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// renameFilesInDirectory renames files in the specified directory to random names.
func renameFilesInDirectory(directoryPath string) error {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			extension := filepath.Ext(file.Name())
			randomName := generateRandomName(32) // Adjust the length as needed
			newName := randomName + extension
			oldPath := filepath.Join(directoryPath, file.Name())
			newPath := filepath.Join(directoryPath, newName)

			err := os.Rename(oldPath, newPath)
			if err != nil {
				return err
			}
			fmt.Printf("Renamed %s to %s\n", oldPath, newPath)
		}
	}
	return nil
}

func bwa() {
	directoryPath := getVTOLDir()[0] // Replace with the actual directory path
	err := renameFilesInDirectory(directoryPath + "\\RadioMusic\\")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
