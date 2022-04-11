package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Begin tracking a new habit",
	Long: `Begins tracking habit patterns for a new habit, that can be viewed with the view command.

Example Usage:
$ habi track stretching`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("track called")
	},
}

func init() {
	rootCmd.AddCommand(trackCmd)
}
