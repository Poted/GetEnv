package getenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv(file string) error {

	// Open the .env file
	envFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("error opening .env file: %w", err)
	}
	defer envFile.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines and comments (lines starting with #)
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		// Split the line into key and value (key=value)
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {

			key, value := parts[0], parts[1]

			// Remove apostrophes if there are any
			if value[0] == '"' {
				value = value[1:]
			}
			if value[len(value)-1] == '"' {
				value = value[:len(value)-1]
			}

			// Set the environment variable using os.Setenv
			err := os.Setenv(key, value)
			if err != nil {
				return fmt.Errorf("error setting environment variable %s: %w", key, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading .env file: %w", err)
	}

	return nil
}
