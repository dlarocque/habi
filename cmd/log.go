package cmd

import (
	"log"
	"path/filepath"

	"github.com/dlarocque/habi/internal/data"
	"github.com/dlarocque/habi/internal/habit"
	"github.com/spf13/cobra"
)

// logCmd represents the track command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Logs a habit as being completed",
	Long: `Logs that a habit has been completed, allowing users to look back at habit patterns for a habit over time.
`,
	Run: func(cmd *cobra.Command, args []string) {
		dataPath := filepath.Join(absRootPath, data.DataPath, data.JsonDataFileName)
		data, err := data.GetJsonData(dataPath)
		if err != nil {
			panic(err) // TODO: Should not manually invoke a panic
		}

		if len(args) == 0 {
			log.Fatalf("No habit name provided.")
		}

		habit.LogHabit(data, args[0])
		data.MarshalAndWrite(dataPath)
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
