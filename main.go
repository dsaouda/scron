package main

import (
	"fmt"
	scron "github.com/dsaouda/scron/util"
	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()

	for _, cron := range scron.Crontab("crontab") {
		fmt.Println("Int spec", cron.Spec, "command", cron.Command)
		c.AddFunc(cron.Spec, cronFunc(cron))
	}

	c.Run()
}

func cronFunc(cron scron.Cron) func() {
	return func() {
		out, _ := scron.ShellRun(cron.Command)
		fmt.Println("Run", cron.Command, "at", time.Now(), "output: ", out)
	}
}
