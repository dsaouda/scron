package util

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ShellRun cria uma interface de comunicação com CMD ou BASH
func ShellRun(line string) (string, error) {
	var out []byte
	var err error

	if runtime.GOOS == "windows" {
		out, err = exec.Command(os.Getenv("COMSPEC"), "/c", line).CombinedOutput()
	} else {
		out, err = exec.Command(os.Getenv("SHELL"), "-c", line).CombinedOutput()
	}

	return strings.TrimSpace(string(out)), err
}
