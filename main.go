package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	viewCommandName  = "view"
	trackCommandName = "track"
	logCommandName   = "log"
)

func main() {
	args := os.Args

	if err := validateArguments(args); err != nil {
		fmt.Println(err)
		return
	}
	if err := parseArguments(args); err != nil {
		fmt.Println(err)
		return
	}
}

func parseArguments(args []string) error {
	var habitName string
	action := args[1]

	if action == trackCommandName {
		habitName = args[2]
		trackHabit(habitName)
	} else if action == logCommandName {
		habitName = args[2]
		logHabit(habitName)
	} else if action == viewCommandName {
		if len(args) == 3 {
			habitName = args[2]
			viewHabit(habitName)
		} else {
			viewAllHabits()
		}
	}

	return nil
}

func validateArguments(args []string) error {
	if len(args) == 1 {
		return errors.New("Invalid arguments")
	}
	return nil
}

func trackHabit(habitName string) {
	log.Printf("Tracking habit: %s", habitName)
}

func logHabit(habitName string) {
	log.Printf("Logging habit: %s", habitName)

}

func viewHabit(habitName string) {
	log.Printf("Viewing habit: %s", habitName)

}

func viewAllHabits() {
	log.Printf("Viewing all habits")

}
