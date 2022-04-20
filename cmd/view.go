/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// trackCmd represents the track command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Visually displays habit patterns over time.",
	Long: `Shows which days habits were logged since tracking began.
Patterns can be viewed for a single habit or all habits at once.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement
		fmt.Println("view called")
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
