package util

import (
	"bufio"
	"github.com/robfig/cron"
	"os"
	"strings"
)

// Representa uma estrutura de cron
type Cron struct {
	Spec string
	Command string
}

// Trabalhar o arquivo de cron tab e gerar a estrura
func Crontab(cronFile string) []Cron {
	file, err := os.Open(cronFile)

	s := make([]Cron, 0)

	if err != nil {
		return []Cron{}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		// ignorar linha em branco
		if len(text) == 0 {
			continue
		}

		// ignorar quando iniciar com comentário
		substringFirstChar := string([]rune(text)[0])
		if substringFirstChar == "#" {
			continue
		}

		expr := strings.Split(text, " ")

		if len(expr) < 6 {
			continue
		}

		spec := strings.Join(expr[0:6], " ")

		_, error := cron.Parse(spec)
		if error != nil {
			continue
		}

		command := strings.Join(expr[6:], " ")
		s = append(s, Cron{Spec: spec, Command: command})
	}

	if err := scanner.Err(); err != nil {
		return []Cron{}
	}

	return s
}
