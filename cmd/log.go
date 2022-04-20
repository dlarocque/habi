package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// logCmd represents the track command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Logs a habit as being completed",
	Long: `Logs that a habit has been completed, allowing users to look back at habit patterns for a habit over time.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("log called")
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
