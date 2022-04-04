package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

const (
	viewCommandName      = "view"
	trackCommandName     = "track"
	logCommandName       = "log"
	jsonDataPath         = "./data/data.json"
	jsonTemplateDataPath = "./data/template.json"
)

type Data struct {
	Habits map[string][]string
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

func getJsonData(filePath string) (Data, error) {
	log.Printf("getting JSON data")
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
