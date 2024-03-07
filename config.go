package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// use a config file to write data
// ADD MORE CONN AND DISCON
func LoadConf(path string) {
	file, err := os.Open(path)
	if err != nil {
		logrus.Error("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	config := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Неверный формат строки:", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		defenition := strings.TrimSpace(parts[1])

		config[key] = defenition
	}
	/*for key, def := range config {
		os.Setenv(key, def)
	}*/

}
