package util

import (
	"bufio"
	"fmt"
	"github.com/robfig/cron"
	"os"
	"strings"
)

// Cron represents a cron structure
type Cron struct {
	Spec    string
	Command string
}

// Crontab read the crontab file and generate the structure to be worked on
func Crontab(cronFile string) ([]Cron, error) {
	file, err := os.Open(cronFile)

	s := make([]Cron, 0)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	line := 0
	for scanner.Scan() {
		line++
		text := strings.TrimSpace(scanner.Text())

		// skip empty line
		if len(text) == 0 {
			fmt.Println("line", line, ": skip empty line")
			continue
		}

		// skip comments
		substringFirstChar := string([]rune(text)[0])
		if substringFirstChar == "#" {
			fmt.Println("line", line, ": skip comment")
			continue
		}

		expr := strings.Split(text, " ")

		if len(expr) < 6 {
			fmt.Println("line", line, ": skip by pattern")
			continue
		}

		spec := strings.Join(expr[0:6], " ")

		_, err := cron.Parse(spec)
		if err != nil {
			fmt.Println("line", line, ": skip error parser -->", err)
			continue
		}

		command := strings.Join(expr[6:], " ")
		s = append(s, Cron{Spec: spec, Command: command})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return s, nil
}
