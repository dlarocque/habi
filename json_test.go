package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var (
	jsonValidDataPath = "./testdata/valid.json"
)

func TestJsonEqual(t *testing.T) {
	t.Log("Checking if newlines are compared")
	var data1, data2 Data
	jsonData1 := `{}
	`
	jsonData2 := `{}`

	if err := json.Unmarshal([]byte(jsonData1), &data1); err != nil {
		t.Error(err)
	}

	if err := json.Unmarshal([]byte(jsonData2), &data2); err != nil {
		t.Error(err)
	}

	if !Data(data1).equalJson(Data(data2)) {
		t.Fail()
	}

}

func TestInitJsonData(t *testing.T) {
	t.Log("Checking if InitJsonData returns the json data template")

	templateJson, err := getJsonData(jsonTemplateDataPath)
	if err != nil {
		t.Errorf("Failed to get json template data, %v", err)
	}

	initJson, err := initJsonData()
	if err != nil {
		t.Errorf("Failed to initialize json data, %v", err)
	}

	if !initJson.equalJson(templateJson) {
		t.Fail()
	}

	// InitJsonData writes the new data file template
	t.Log("Checking if InitJsonData writes the new data file template")

	if _, err := os.Stat(jsonDataPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			t.Fail()
		} else {
			// Check if the template was created

		}
	}
}

func TestGetJsonData(t *testing.T) {
	t.Log("Checking that if json data does not exist, the template is generated and returned")

	jsonTemplateData, err := getJsonData(jsonTemplateDataPath)
	if err != nil {
		t.Errorf("Failed to get json template data")
	}

	filePath := ""
	jsonData, err := getJsonData(filePath)
	if err != nil {
		t.Errorf("Failed to execute when JSON data does not exist")
	}

	if !jsonData.equalJson(jsonTemplateData) {
		t.Fail()
	}

	t.Log("Checking that if json data does exist, json data is returned")

	jsonValidData, err := getJsonData(jsonValidDataPath)
	if err != nil {
		t.Errorf("Failed to get valid json data")
	}

	data, err := json.MarshalIndent(jsonValidData, "", "    ")
	if err != nil {
		t.Error(err)
	}

	expectedData, err := ioutil.ReadFile(jsonValidDataPath)
	if err != nil {
		t.Error(err)
	}

	var expectedJson Data
	json.Unmarshal(expectedData, &expectedJson)
	expectedData, _ = json.MarshalIndent(expectedJson, "", "    ")

	if !reflect.DeepEqual(data, expectedData) {
		t.Fail()
	}

}

func TestTrackHabit(t *testing.T) {
	habitName := "stretching"
	templateJson, err := readAndUnmarshal(jsonTemplateDataPath)
	if err != nil {
		t.Error(err)
	}

	t.Log("Checking to see if new habit exists after adding it")
	templateJson.trackHabit(habitName)
	if _, ok := templateJson.Habits[habitName]; !ok {
		t.Fail()
	}

	validJson, err := readAndUnmarshal(jsonValidDataPath)
	if err != nil {
		t.Error(err)
	}

	t.Log("Checking to see if an existing habit is overwritten when re-added")

	validJson.trackHabit(habitName)
	pattern, ok := validJson.Habits[habitName]
	if (ok && !reflect.DeepEqual(pattern, []string{"test"})) || (!ok) {
		t.Fail()
	}

}
