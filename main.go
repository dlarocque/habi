package main

import (
	"errors"
	"log"
	"time"

	"github.com/dlarocque/habi/cmd"
	"github.com/dlarocque/habi/internal/data"
)

const (
	viewCommandName  = "view"
	trackCommandName = "track"
	logCommandName   = "log"
)

func main() {
	cmd.Execute()
}

func parseArguments(args []string) error {
	var habitName string
	action := args[1]
	jsonData, err := data.GetJsonData(data.JsonDataPath)
	if err != nil {
		return err
	}

	if action == trackCommandName {
		habitName = args[2]
		jsonData.trackHabit(habitName) // Name collision
	} else if action == logCommandName {
		habitName = args[2]
		jsonData.logHabit(habitName) // Name collision
	} else if action == viewCommandName {
		if len(args) == 3 {
			habitName = args[2]
			viewHabit(habitName)
		} else {
			viewAllHabits()
		}
	}

	if err := jsonData.MarshalAndWrite(data.JsonDataPath); err != nil {
		return err
	}
	return nil
}

func validateArguments(args []string) error {
	if len(args) == 1 {
		return errors.New("Invalid arguments")
	}
	return nil
}

func (d data.Data) trackHabit(habitName string) {
	log.Printf("Tracking habit: %s", habitName)

	// Don't do anything if the habit already exists
	if _, ok := d.Habits[habitName]; ok {
		log.Printf("Habit %s already exists", habitName)
		return
	}

	var habit []time.Time
	d.Habits[habitName] = habit
}

func (d Data) logHabit(habitName string) {
	log.Printf("Logging %s", habitName)

	pattern, ok := d.Habits[habitName]
	if !ok {
		log.Printf("Failed to log %s as it is not being tracked\n", habitName)
		log.Printf("To begin tracking %s, use habi track %s", habitName, habitName)
		return
	}

	todaysDate := time.Now()
	if len(pattern) > 0 {
		mostRecentDate := pattern[len(pattern)-1]
		mostRecentYear, mostRecentMonth, mostRecentDay := mostRecentDate.Date()

		currentYear, currentMonth, currentDay := todaysDate.Date()
		if (currentYear == mostRecentYear) &&
			(currentMonth == mostRecentMonth) &&
			(currentDay == mostRecentDay) {
			log.Printf("Already logged %s today!", habitName)
			return
		}
	}

	updatedPattern := append(pattern, todaysDate)
	d.Habits[habitName] = updatedPattern
}

func viewHabit(habitName string) {
	log.Printf("Viewing habit: %s", habitName)

}

func viewAllHabits() {
	log.Printf("Viewing all habits")
}
