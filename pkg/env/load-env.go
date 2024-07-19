package env

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnvFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid string format: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		err = os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("couldn't set the environment variable: %s", err)
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}
