package util

import (
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

var re = regexp.MustCompile(`\s`)
var re2 = regexp.MustCompile(`(?m)"(.*)"`)

// ShellRun creates a communication interface with CMD or BASH
func ShellRun(command string) (string, error) {
	var out []byte
	var err error

	for _, match := range re2.FindAllString(command, -1) {

		// use &nbsp; to easy trick to split command
		matchNew := strings.ReplaceAll(match, " ", "&nbsp;")
		command = strings.ReplaceAll(command, match, matchNew)
	}

	var cs = re.Split(command, -1)

	commands := make([]string, 0)

	for _, c := range cs {

		// remove &nbsp;
		c = strings.ReplaceAll(c, "\"", "")
		c = strings.ReplaceAll(c, "&nbsp;", " ")
		commands = append(commands, c)
	}

	if runtime.GOOS == "windows" {
		commands = append([]string{"/c"}, commands...)
		out, err = exec.Command(os.Getenv("COMSPEC"), commands...).CombinedOutput()
	} else {
		commands = append([]string{"-c"}, commands...)
		out, err = exec.Command(os.Getenv("SHELL"), commands...).CombinedOutput()
	}

	return strings.TrimSpace(string(out)), err
}
