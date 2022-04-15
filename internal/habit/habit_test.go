package habit

import (
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/dlarocque/habi/internal/data"
)

var (
	// TODO: Make these not as highly coupled with the file paths
	dataPath               = "../../data/"
	testDataPath           = "testdata"
	jsonTemplateDataPath   = filepath.Join(dataPath, "template.json")
	jsonValidDataPath      = filepath.Join(testDataPath, "valid.json")
	jsonValidHabitDataPath = filepath.Join(testDataPath, "validhabit.json")
)

func TestTrackHabit(t *testing.T) {
	habitName := "stretching"
	templateJson, err := data.ReadAndUnmarshal(jsonTemplateDataPath)
	if err != nil {
		t.Error(err)
	}

	t.Log("Checking to see if new habit exists after adding it")
	trackHabit(templateJson, habitName)
	if _, ok := templateJson.Habits[habitName]; !ok {
		t.Fail()
	}

	validJson, err := data.ReadAndUnmarshal(jsonValidHabitDataPath)
	if err != nil {
		t.Error(err)
	}

	t.Log("Checking to see if an existing habit is overwritten when re-added")

	prevHabitPattern := validJson.Habits[habitName]
	trackHabit(validJson, habitName)
	pattern, ok := validJson.Habits[habitName]
	if (ok && !reflect.DeepEqual(pattern, prevHabitPattern)) || (!ok) {
		t.Fail()
	}

}

func TestLogHabit(t *testing.T) {
	habitName := "stretching"

	t.Log("Checking to see if a habit is logged if the habit is not being tracked")
	templateJson, err := data.ReadAndUnmarshal(jsonTemplateDataPath)
	if err != nil {
		t.Errorf("Failed to get valid json data")
	}

	logHabit(templateJson, habitName)
	if _, ok := templateJson.Habits[habitName]; ok {
		t.Fail()
	}

	t.Log("Checking to see if a habit is updated if the habit has not been updated today")
	validJson, err := data.ReadAndUnmarshal(jsonValidDataPath)
	if err != nil {
		t.Errorf("Failed to get valid json data")
	}

	prevNumLogs := len(validJson.Habits[habitName])
	logHabit(validJson, habitName)
	numLogs := len(validJson.Habits[habitName])
	if numLogs == 0 || numLogs-prevNumLogs != 1 {
		t.Fail()
	}
	year, month, day := time.Now().Date()
	pattern := validJson.Habits[habitName]
	pYear, pMonth, pDay := pattern[len(pattern)-1].Date()
	if (year != pYear) || (month != pMonth) || (day != pDay) {
		t.Fail()
	}

	t.Log("Checking to see if a habit is updated if the habit has already been updated today")
	validJson, err = data.ReadAndUnmarshal(jsonValidDataPath)
	if err != nil {
		t.Errorf("Failed to get valid json data")
	}

	prevNumLogs = len(validJson.Habits[habitName])
	logHabit(validJson, habitName)
	logHabit(validJson, habitName)
	numLogs = len(validJson.Habits[habitName])
	if numLogs-prevNumLogs > 1 {
		t.Fail()
	}
}
