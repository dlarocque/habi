package cmd

import (
	"path/filepath"

	"github.com/dlarocque/habi/internal/data"
	"github.com/dlarocque/habi/internal/habit"
	"github.com/spf13/cobra"
)

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Begin tracking a new habit",
	Long: `Begins tracking habit patterns for a new habit, that can be viewed with the view command.
`,
	Run: func(cmd *cobra.Command, args []string) {
		dataPath := filepath.Join("..", data.DataPath, data.JsonDataFileName)
		data, err := data.GetJsonData(dataPath)
		if err != nil {
			panic(err) // TODO: Should not manually invoke a panic
		}

		habit.TrackHabit(data, args[0])
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
}
