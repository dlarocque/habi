package habit

import (
	"log"
	"time"

	"github.com/dlarocque/habi/internal/data"
)

func trackHabit(jsonData data.Data, habitName string) error {
	log.Printf("Tracking habit: %s", habitName)

	// Don't do anything if the habit already exists
	if _, ok := jsonData.Habits[habitName]; ok {
		log.Printf("Habit %s already exists", habitName)
		return nil
	}

	var habit []time.Time
	jsonData.Habits[habitName] = habit
	return nil
}

func logHabit(jsonData data.Data, habitName string) error {
	log.Printf("Logging %s", habitName)

	pattern, ok := jsonData.Habits[habitName]
	if !ok {
		log.Printf("Failed to log %s as it is not being tracked\n", habitName)
		log.Printf("To begin tracking %s, use habi track %s", habitName, habitName)
		return nil
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
			return nil
		}
	}

	updatedPattern := append(pattern, todaysDate)
	jsonData.Habits[habitName] = updatedPattern
	return nil
}

func viewHabit(habitName string) {
	log.Printf("Viewing habit: %s", habitName)

}

func viewAllHabits() {
	log.Printf("Viewing all habits")
}
