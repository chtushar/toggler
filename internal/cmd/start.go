package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func start(cmd *cobra.Command, _ []string) {
	fmt.Printf("Starting toggler server")
}
