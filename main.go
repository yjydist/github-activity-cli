package main

import (
	"log"

	"github.com/yjydist/github-activity-cli/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
		return
	}
}
