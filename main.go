package main

import "github.com/wjdp/htmltest/cmd"

var (
	version   string
	buildDate string
)

func main() {
	cmd.Version = version
	cmd.BuildDate = buildDate
	cmd.Execute()
}
