package main

import (
	"zap/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "nont"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}
