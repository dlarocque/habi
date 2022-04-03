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
	var data1, data2 map[string]interface{}
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

	templateData, err := json.MarshalIndent(templateJson, "", "    ")
	if err != nil {
		t.Error(err)
	}

	initData, err := json.MarshalIndent(initJson, "", "    ")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(templateData, initData) {
		t.Errorf("Initial json data does not match expected template json data")
	}

	// InitJsonData writes the new data file template
	t.Log("Checking if InitJsonData writes the new data file template")

	if _, err := os.Stat(jsonDataPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			t.Errorf("Initial json data not written to expected location: %s", jsonDataPath)
		} else {

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

	data, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		t.Error(err)
	}

	templateData, err := json.MarshalIndent(jsonTemplateData, "", "    ")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(data, templateData) {
		t.Errorf("Failed to correctly initialize JSON data when JSON data does not exist")
	}

	t.Log("Checking that if json data does exist, the json data is returned")

	jsonValidData, err := getJsonData(jsonValidDataPath)
	if err != nil {
		t.Errorf("Failed to get valid json data")
	}

	data, err = json.MarshalIndent(jsonValidData, "", "    ")
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
		t.Errorf("Failed to correctly retrieve JSON data when valid data does exist")
	}

}
