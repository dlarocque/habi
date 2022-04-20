package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

var (
	rootPath             = "../.."
	JsonDataFileName     = "data.json"
	JsonTemplateFileName = "template.json"
	DataPath             = "data"
	jsonDataPath         = filepath.Join(rootPath, DataPath, JsonDataFileName)
	jsonTemplateDataPath = filepath.Join(rootPath, DataPath, "template.json")
)

type Data struct {
	Habits map[string][]time.Time
}

func GetJsonData(filePath string) (Data, error) {
	log.Printf("getting JSON data at %s", filePath)
	var jsonData Data
	if _, err := os.Stat(filePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// data does not exist in expected directory, create it
			return InitJsonData()
		} else {
			return Data{}, err
		}

	}

	jsonData, err := ReadAndUnmarshal(filePath)
	if err != nil {
		return Data{}, err
	}

	return jsonData, nil
}

func InitJsonData() (Data, error) {
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

func ReadAndUnmarshal(filePath string) (Data, error) {
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

func (d Data) EqualJson(other Data) bool {
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

func (d Data) MarshalAndWrite(fileName string) error {
	bytes, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, bytes, 0644); err != nil {
		return err
	}

	return nil
}
