package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var (
	// TODO: Make these not as highly coupled with the file paths
	testDataPath           = "testdata"
	dataPath               = filepath.Join(rootPath, "data/")
	jsonTemplateDataPath   = filepath.Join(dataPath, "template.json")
	jsonValidDataPath      = filepath.Join(testDataPath, "valid.json")
	jsonValidHabitDataPath = filepath.Join(testDataPath, "validhabit.json")
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

	if !Data(data1).EqualJson(Data(data2)) {
		t.Fail()
	}

}

func TestInitJsonData(t *testing.T) {
	t.Log("Checking if InitJsonData returns the json data template")

	templateJson, err := GetJsonData(jsonTemplateDataPath)
	if err != nil {
		t.Errorf("Failed to get json template data, %v", err)
	}

	initJson, err := InitJsonData()
	if err != nil {
		t.Errorf("Failed to initialize json data, %v", err)
	}

	if !initJson.EqualJson(templateJson) {
		t.Fail()
	}

	// InitJsonData writes the new data file template
	t.Log("Checking if InitJsonData writes the new data file template")

	if _, err := os.Stat(filepath.Join(JsonDataPath)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			t.Fail()
		} else {
			// Check if the template was created

		}
	}
}

func TestGetJsonData(t *testing.T) {
	t.Log("Checking that if json data does not exist, the template is generated and returned")

	jsonTemplateData, err := GetJsonData(jsonTemplateDataPath)
	if err != nil {
		t.Errorf("Failed to get json template data")
	}

	filePath := ""
	jsonData, err := GetJsonData(filePath)
	if err != nil {
		t.Errorf("Failed to execute when JSON data does not exist")
	}

	if !jsonData.EqualJson(jsonTemplateData) {
		t.Fail()
	}

	t.Log("Checking that if json data does exist, json data is returned")

	jsonValidData, err := GetJsonData(jsonValidDataPath)
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
