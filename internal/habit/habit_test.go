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

	trackHabit(templateJson, habitName)
	if _, ok := templateJson.Habits[habitName]; !ok {
		t.Errorf("trackHabit does not add a new habit to the data")
	}

	validJson, err := data.ReadAndUnmarshal(jsonValidHabitDataPath)
	if err != nil {
		t.Error(err)
	}

	prevHabitPattern := validJson.Habits[habitName]
	trackHabit(validJson, habitName)
	pattern, ok := validJson.Habits[habitName]
	if (ok && !reflect.DeepEqual(pattern, prevHabitPattern)) || (!ok) {
		t.Fail()
		t.Errorf("trackHabit overwrites existing habit data when a habit is re-added")
	}

}

func TestLogHabit(t *testing.T) {
	habitName := "stretching"

	templateJson, err := data.ReadAndUnmarshal(jsonTemplateDataPath)
	if err != nil {
		t.Errorf("Failed to get valid json data")
	}

	logHabit(templateJson, habitName)
	if _, ok := templateJson.Habits[habitName]; ok {
		t.Fail()
		t.Errorf("logHabit still tracks a habit if it does not exist")
	}

	validJson, err := data.ReadAndUnmarshal(jsonValidDataPath)
	if err != nil {
		t.Errorf("Failed to get valid json data")
	}

	prevNumLogs := len(validJson.Habits[habitName])
	logHabit(validJson, habitName)
	numLogs := len(validJson.Habits[habitName])
	if numLogs == 0 || numLogs-prevNumLogs != 1 {
		t.Errorf("logHabit does not a log a habit if it has not yet been done today")
	}
	year, month, day := time.Now().Date()
	pattern := validJson.Habits[habitName]
	pYear, pMonth, pDay := pattern[len(pattern)-1].Date()
	if (year != pYear) || (month != pMonth) || (day != pDay) {
		t.Errorf("logHabit does not log a habit if it has not yet been done today")
	}

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
		t.Errorf("logHabit allows a habit to log a habit twice in the same day, it should not")
	}
}
