package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	addCommands()
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
