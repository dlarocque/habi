package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"time"
)

const (
	viewCommandName      = "view"
	trackCommandName     = "track"
	logCommandName       = "log"
	jsonDataPath         = "./data/data.json"
	jsonTemplateDataPath = "./data/template.json"
)

type Data struct {
	Habits map[string][]time.Time
}

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
	data, err := getJsonData(jsonDataPath)
	if err != nil {
		return err
	}

	if action == trackCommandName {
		habitName = args[2]
		data.trackHabit(habitName)
	} else if action == logCommandName {
		habitName = args[2]
		data.logHabit(habitName)
	} else if action == viewCommandName {
		if len(args) == 3 {
			habitName = args[2]
			viewHabit(habitName)
		} else {
			viewAllHabits()
		}
	}

	if err := data.marshalAndWrite(jsonDataPath); err != nil {
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

func (d Data) trackHabit(habitName string) {
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

func getJsonData(filePath string) (Data, error) {
	log.Printf("getting JSON data at %s", filePath)
	var jsonData Data
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// data does not exist in expected directory, create it
			return initJsonData()
		} else {
			return Data{}, err
		}

	}

	jsonData, err := readAndUnmarshal(filePath)
	if err != nil {
		return Data{}, err
	}

	return jsonData, nil
}

func initJsonData() (Data, error) {
	log.Printf("Initializing JSON data")
	jsonTemplate, err := ioutil.ReadFile(jsonTemplateDataPath)
	if err != nil {
		return Data{}, err
	}

	if err := ioutil.WriteFile(jsonDataPath, jsonTemplate, 0644); err != nil {
		return Data{}, err
	}

	var jsonData Data
	if err = json.Unmarshal(jsonTemplate, &jsonData); err != nil {
		return Data{}, err
	}

	return jsonData, nil
}

func readAndUnmarshal(filePath string) (Data, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Data{}, err
	}

	var jsonData Data
	if err = json.Unmarshal(data, &jsonData); err != nil {
		return Data{}, err
	}

	return jsonData, nil
}

func (d Data) equalJson(other Data) bool {
	jsonData, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return false
	}

	otherJsonData, err := json.MarshalIndent(other, "", "    ")
	if err != nil {
		return false
	}

	return reflect.DeepEqual(jsonData, otherJsonData)
}

func (d Data) marshalAndWrite(fileName string) error {
	bytes, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, bytes, 0644); err != nil {
		return err
	}

	return nil
}
