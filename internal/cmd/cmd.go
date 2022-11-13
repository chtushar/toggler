package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toggler",
	Short: "Toggler is a lightweight CLI tool for toggling features on and off",
	Long:  "TODO",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the toggler server",
	Long:  "TODO",
	Run:   start,
}

func addCommands() {
	rootCmd.AddCommand(startCmd)
}
