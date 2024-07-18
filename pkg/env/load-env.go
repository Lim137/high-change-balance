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
			return fmt.Errorf("неверный формат строки: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		err = os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("ошибка при установке переменной окружения: %s", err)
		}
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	return nil
}
