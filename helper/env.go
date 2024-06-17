package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SetEnv(fileName string) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text() // Text returns the current token, here the next line, from the input
		envVariables := strings.Split(line, "=")
		err := os.Setenv(envVariables[0], envVariables[1])
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
