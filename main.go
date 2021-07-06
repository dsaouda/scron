package main

import (
	"fmt"
	scron "github.com/dsaouda/scron/util"
	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()

	cronlines, err := scron.Crontab("crontab")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if len(cronlines) == 0 {
		fmt.Println("Error: No lines match cron pattern")
		return
	}

	for _, cronEntity := range cronlines {
		fmt.Println("add spec=", cronEntity.Spec, "command=", cronEntity.Command)
		c.AddFunc(cronEntity.Spec, cronFunc(cronEntity))
	}

	c.Run()
}

func cronFunc(cron scron.Cron) func() {
	return func() {
		out, _ := scron.ShellRun(cron.Command)
		fmt.Println("Run", cron.Command, "at", time.Now(), "output: ", out)
	}
}
