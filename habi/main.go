package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	viewCommandName  = "view"
	trackCommandName = "track"
	logCommandName   = "log"
)

func main() {
	args := os.Args

	err := validateArguments(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	parseArguments(args)
}

func parseArguments(args []string) error {
	action := args[1]
	// habit := args[2]

	if action == viewCommandName {
		// ExecuteView()
		// ExecuteView(habit)
	} else if action == trackCommandName {
		// ExecuteView(habit)
	} else if action == logCommandName {
		// ExecuteLog(habit)
	}

	return nil
}

func validateArguments(args []string) error {
	if len(args) == 1 {
		return errors.New("Invalid arguments")
	}
	return nil
}
