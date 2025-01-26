package main

import (
	"github-activity-cli/cmd"
	"log"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
		return
	}
}
