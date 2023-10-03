package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadDotEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Printf("Invalid .env line: %s\n", line)
			continue
		}
		key := parts[0]
		value := parts[1]
		os.Setenv(key, value)
	}

	return scanner.Err()
}
