package util

import (
	"bufio"
	"os"
	"strings"
)

type Cron struct {
	Spec string
	Command string
}

func Crontab(cronFile string) []Cron {
	file, err := os.Open(cronFile)

	s := make([]Cron, 0)

	if err != nil {
		return []Cron{}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		expr := strings.Split(scanner.Text(), " ")

		spec := strings.Join(expr[0:6], " ")
		command := strings.Join(expr[6:], " ")

		s = append(s, Cron{Spec: spec, Command: command})
	}

	if err := scanner.Err(); err != nil {
		return []Cron{}
	}

	return s
}
