# scron

Simple cron linux for any OS

## Why

Especially for running on windows, but you can use it to run on linux if you want to run a command every second.

## Crontab expression

Example of job definition:

<pre>
.------------------ seconds (0 - 59)
| .---------------- minute (0 - 59)
| |  .------------- hour (0 - 23)
| |  |  .---------- day of month (1 - 31)
| |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
| |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
| |  |  |  |  |
* *  *  *  *  *  command to be executed
</pre>

For more detail see https://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format

## Download

Download the latest version in https://github.com/dsaouda/scron/releases or run `go get github.com/dsaouda`

## Run

Create a file named *crontab* in the same directory as the binary then run `scron`

## Example crontab file for windows

`* * * * * * echo %time% >> time.txt`

This cron print OS time every second
